package core

import (
	"fmt"
	"time"

	"github.com/snes-emu/gose/bit"
)

var lastFrame = time.Now()
var frameDuration = time.Duration(16666667)

type pixel struct {
	bgr      uint16
	visible  bool
	priority uint8
}

func (ppu *PPU) renderLine() {
	fmt.Printf("Render line: %v\n", ppu.vCounter)
	pixels := ppu.renderSpriteLine()
	for i := 0; i < HMax; i++ {
		lo, hi := bit.SplitUint16(pixels[i].bgr)
		ppu.screen[int(ppu.vCounter)*HMax*2+2*i] = lo
		ppu.screen[int(ppu.vCounter)*HMax*2+2*i+1] = hi
	}
	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter == ppu.VDisplay()+1 {
		fmt.Println("VBlank !")
		ppu.render(ppu.screen)
		delta := time.Since(lastFrame)
		fmt.Println(delta.String(), frameDuration.String())
		time.Sleep(time.Duration(frameDuration.Nanoseconds() - delta.Nanoseconds()))
		lastFrame = time.Now()
		ppu.cpu.enterVblank()
	}

	if ppu.vCounter == 0 {
		ppu.cpu.leavVblank()
	}
}

// Return the color index for a sprite inside the sprite palette
// Each 8x8 tile occupies 16, 32, or 64 bytes (for 4, 16, or 256 colors). BG tiles can be 4/16/256 colors (depending on BG Mode), OBJs are always 16 color.

//   Color Bits (Planes)     Upper Row ........... Lower Row
//   Plane 0 stored in bytes 00h,02h,04h,06h,08h,0Ah,0Ch,0Eh ;\for 4/16/256 colors
//   Plane 1 stored in bytes 01h,03h,05h,07h,09h,0Bh,0Dh,0Fh ;/
//   Plane 2 stored in bytes 10h,12h,14h,16h,18h,1Ah,1Ch,1Eh ;\for 16/256 colors
//   Plane 3 stored in bytes 11h,13h,15h,17h,19h,1Bh,1Dh,1Fh ;/
//   Plane 4 stored in bytes 20h,22h,24h,26h,28h,2Ah,2Ch,2Eh ;\
//   Plane 5 stored in bytes 21h,23h,25h,27h,29h,2Bh,2Dh,2Fh ; for 256 colors
//   Plane 6 stored in bytes 30h,32h,34h,36h,38h,3Ah,3Ch,3Eh ;
//   Plane 7 stored in bytes 31h,33h,35h,37h,39h,3Bh,3Dh,3Fh ;/
//   In each byte, bit7 is left-most, bit0 is right-most.
//   Plane 0 is the LSB of color number.
func (ppu *PPU) getColorIndex(tileAddress, colorDepth, x, y uint16) uint8 {
	var colorIndex uint8

	lineBaseAddress := tileAddress + 2*y
	colorIndex += (ppu.vram.bytes[lineBaseAddress+0x00] >> (7 - x) & 1) << 0
	colorIndex += (ppu.vram.bytes[lineBaseAddress+0x01] >> (7 - x) & 1) << 1
	if colorDepth >= 4 {
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x10] >> (7 - x) & 1) << 2
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x11] >> (7 - x) & 1) << 3
	}
	if colorDepth >= 8 {
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x20] >> (7 - x) & 1) << 4
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x21] >> (7 - x) & 1) << 5
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x30] >> (7 - x) & 1) << 6
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x31] >> (7 - x) & 1) << 7
	}
	return colorIndex
}

func (ppu *PPU) renderSpriteLine() [HMax]pixel {
	// Initialize pixel line
	var pixels [HMax]pixel
	sprites := make([]sprite, 0, 32)
	firstSprite := uint16(0)
	if ppu.oam.priorityBit {
		firstSprite = (ppu.oam.addr >> 2) & 0xFF
	}
	// Choose the first 32 sprites to appear on screen
	for i := firstSprite; i < firstSprite+128; i++ {
		// Check if line intersects sprite
		spriteIndex := i % 128
		sprite := ppu.getSpriteByIndex(spriteIndex)
		if ppu.vCounter >= sprite.y && ppu.vCounter < sprite.y+sprite.vSize {
			if len(sprites) == 32 {
				ppu.status.rangeOver = true
				break
			} else {
				sprites = append(sprites, sprite)
			}
		}
	}
	// Go through all the selected sprites in reverse order (up to 34 tiles)
	tiles := uint16(0)
	for i := len(sprites) - 1; i >= 0; i-- {
		sprite := sprites[i]
		// Go through all tiles containing the line
		// Tiles are stored in the 2D-array 0xNyx

		// Y coordinate of the tile containing the line
		var yTile = (ppu.vCounter - sprite.y) / 8
		// Y coordinate of the line in the tile
		var y = (ppu.vCounter - sprite.y) % 8
		if sprite.vFlip {
			yTile = (sprite.vSize - 1 - (ppu.vCounter - sprite.y)) / 8
			y = (sprite.vSize - 1 - (ppu.vCounter - sprite.y)) % 8
		}
		// Go through pixels in reverse order if hFlip is true
		for baseTileOffset := yTile << 4; baseTileOffset < yTile<<4+sprite.hSize/8; baseTileOffset++ {
			if tiles == 34 {
				ppu.status.timeOver = true
				break
			}
			// Tile coordinate in the form 0xYX
			var tileOffset uint16
			if sprite.hFlip {
				tileOffset = 2*(yTile<<4) + sprite.hSize/8 - 1 - baseTileOffset
			} else {
				tileOffset = baseTileOffset
			}

			//Address of the current tile in VRAM
			tileAddress := sprite.tileAddress + tileOffset<<5

			// Go through all pixel in the line
			for pix := uint16(0); pix < 8; pix++ {
				var x uint16
				if sprite.hFlip {
					x = 7 - pix
				} else {
					x = pix
				}
				// Get the color index of the pixel
				colorIndex := ppu.getColorIndex(tileAddress, 4, x, y)
				// If color is not tansparent, write color value in the pixel
				if colorIndex != 0 {
					colorAddress := 2 * (0x80 + sprite.paletteIndex<<4 + uint16(colorIndex))
					colorBgr := bit.JoinUint16(ppu.cgram.bytes[colorAddress], ppu.cgram.bytes[colorAddress+1])
					if ppu.display.brightness != 0 {
						r := int(colorBgr) & 0x1F * (int(ppu.display.brightness) + 1) / 16
						g := int(colorBgr) & 0x3E0 >> 5 * (int(ppu.display.brightness) + 1) / 16
						b := int(colorBgr) & 0x7C00 >> 10 * (int(ppu.display.brightness) + 1) / 16
						colorBgr = uint16(b<<10 | g<<5 | r)
					} else {
						colorBgr = 0
					}

					pixels[(sprite.x+8*(tileOffset&0xF)+x)%HMax] = pixel{
						bgr:      colorBgr,
						visible:  true,
						priority: sprite.priority,
					}
				}
			}
			tiles++
		}
	}
	return pixels
}

func (ppu *PPU) renderBgLine(BG uint8, colorDepth uint16) [HMax]pixel {
	var pixels [HMax]pixel
	bg := ppu.backgroundData.bg[BG]
	y := ppu.vCounter
	var size uint16
	if bg.tileSize {
		size = 16
	} else {
		size = 8
	}
	x := uint16(0)
	for x < HMax {
		tile := ppu.getTileFromBG(BG, (x+bg.horizontalScroll)/size, (y+bg.verticalScroll)/size)
		tileAddress := bg.tileSetBaseAddress<<13 + tile.characterIndex*8*colorDepth
		var yTile uint16
		if tile.vFlip {
			yTile = size - 1 - (y+bg.verticalScroll)%size
		} else {
			yTile = (y + bg.verticalScroll) % size
		}
		for xTile := (x + bg.horizontalScroll) % size; xTile < size; xTile++ {
			colorIndex := ppu.getColorIndex(tileAddress, colorDepth, xTile, yTile)
			var colorAddress uint16
			if ppu.backgroundData.screenMode == 0 {
				colorAddress = (uint16(BG)<<5 + tile.paletteIndex<<colorDepth + uint16(colorIndex)) << 1
			} else {
				colorAddress = (tile.paletteIndex<<colorDepth + uint16(colorIndex)) << 1
			}
			if colorIndex != 0 {
				if tile.vFlip {
					pixels[x+(size-1-xTile)] = pixel{
						bgr:      bit.JoinUint16(ppu.cgram.bytes[colorAddress], ppu.cgram.bytes[colorAddress+1]),
						visible:  true,
						priority: bit.BoolToUint8(tile.priority),
					}
				}
			}
		}
		x += size - (x+bg.horizontalScroll)%size
	}
	return pixels
}
