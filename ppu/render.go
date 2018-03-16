package ppu

import "github.com/snes-emu/gose/utils"

type pixel struct {
	bgr      uint16
	visible  bool
	priority uint8
}

func (ppu PPU) decodeTilePixel(tileAddress, colorDepth, x, y uint16) uint8 {
	var colorIndex uint8
	lineBaseAddress := tileAddress + 2*y
	colorIndex += (ppu.vram.bytes[lineBaseAddress+0x00] >> x & 1) << 0
	colorIndex += (ppu.vram.bytes[lineBaseAddress+0x01] >> x & 1) << 1
	if colorDepth >= 4 {
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x10] >> x & 1) << 2
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x11] >> x & 1) << 3
	}
	if colorDepth >= 256 {
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x20] >> x & 1) << 4
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x21] >> x & 1) << 5
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x30] >> x & 1) << 6
		colorIndex += (ppu.vram.bytes[lineBaseAddress+0x31] >> x & 1) << 7
	}
	return colorIndex
}

func (ppu PPU) renderSpriteLine() [HMax]pixel {
	// Initialize pixel line
	var pixels [HMax]pixel
	sprites := make([]uint16, 0, 32)
	firstSprite := uint16(0)
	if ppu.oam.priorityBit {
		firstSprite = (ppu.oam.addr >> 2) & 0xFF
	}
	// Choose the first 32 sprites to appear on screen
	for i := firstSprite; i < firstSprite+128; i++ {
		// Check if line intersects sprite
		spriteIndex := i % 128
		if ppu.vCounter >= uint16(ppu.oam.bytes[4*spriteIndex+1]) && ppu.vCounter < uint16(ppu.oam.bytes[4*spriteIndex+1])+uint16(spriteSizeTable[ppu.oam.objectSize|((ppu.oam.bytes[0x200+i/4]&(1<<(2*(spriteIndex%4)+1)))>>(1<<(2*(spriteIndex%4)+1)))<<4][1]) {
			if len(sprites) == 32 {
				ppu.status.rangeOver = true
				break
			} else {
				sprites = append(sprites, i)
			}
		}
	}
	// Go through all the selected sprites in reverse order (up to 34 tiles)
	tiles := uint16(0)
	for i := uint16(len(sprites) - 1); i >= 0; i-- {
		sprite := ppu.decodeSprite(i)
		// Go through all tiles containing the line
		// Tiles are stored in the 2D-array 0xNyx

		// Y coordinate of the tile containing the line
		var yTile uint16
		// Y coordinate of the line in the tile
		var y uint16
		if sprite.vFlip {
			yTile = (sprite.vSize - 1 - (ppu.vCounter - sprite.y)) / 8
			y = (sprite.vSize - 1 - (ppu.vCounter - sprite.y)) % 8
		} else {
			yTile = (ppu.vCounter - sprite.y) / 8
			y = (ppu.vCounter - sprite.y) % 8
		}
		// Go through pixels in reverse order if hFlip is true
		for baseTileCoor := yTile << 4; baseTileCoor < yTile<<4+sprite.hSize/8; baseTileCoor++ {
			if tiles == 34 {
				ppu.status.timeOver = true
				break
			}
			// Tile coordinate in the form 0xYX
			var tileCoor uint16
			if sprite.hFlip {
				tileCoor = 2*(yTile<<4) + sprite.hSize/8 - 1 - baseTileCoor
			} else {
				tileCoor = baseTileCoor
			}
			// Go through all pixel in the line
			for pix := uint16(0); pix < 8; pix++ {
				var x uint16
				if sprite.hFlip {
					x = 7 - pix
				} else {
					x = pix
				}
				//Address of the current tile in VRAM
				tileAddress := sprite.tileAddress + tileCoor<<5
				// Get the color index of the pixel
				colorIndex := ppu.decodeTilePixel(tileAddress, 16, x, y)
				// If color is not tansparent, write color value in the pixel
				if colorIndex != 0 {
					colorAddress := 2 * (128 + 16*sprite.paletteIndex + uint16(colorIndex))
					pixels[(sprite.x+8*(tileCoor&0xF)+x)%HMax] = pixel{
						bgr:      utils.JoinUint16(ppu.cgram.bytes[colorAddress], ppu.cgram.bytes[colorAddress+1]),
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
