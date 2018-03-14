package ppu

import "github.com/snes-emu/gose/utils"

type pixel struct {
	bgr      uint16
	visible  bool
	priority uint8
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
		var baseTileIndex uint16
		// Y coordinate of the line in the tile
		var yPixel uint16
		if sprite.vFlip {
			baseTileIndex = (sprite.vSize - 1 - (ppu.vCounter - sprite.y)) / 8 << 4
			yPixel = (sprite.vSize - 1 - (ppu.vCounter - sprite.y)) % 8
		} else {
			baseTileIndex = (ppu.vCounter - sprite.y) / 8 << 4
			yPixel = (ppu.vCounter - sprite.y) % 8
		}
		// Go through pixels in reverse order if hFlip is true
		for tile := baseTileIndex; tile < baseTileIndex+sprite.hSize/8; tile++ {
			if tiles == 34 {
				ppu.status.timeOver = true
				break
			}
			var effectiveTile uint16
			if sprite.hFlip {
				effectiveTile = 2*baseTileIndex + sprite.hSize/8 - 1 - tile
			} else {
				effectiveTile = tile
			}
			// Go through all pixel in the line
			for pix := uint16(0); pix < 8; pix++ {
				var effectivePix uint16
				if sprite.hFlip {
					effectivePix = 7 - pix
				} else {
					effectivePix = pix
				}
				lineBaseAddress := sprite.tileAddress + effectiveTile<<5 + 2*yPixel // Base address of the color planes of the pixel line in VRAM
				// Get the color index of the pixel
				colorIndex := uint8(0)
				colorIndex += (ppu.vram.bytes[lineBaseAddress+0x00] >> effectivePix & 1) << 0
				colorIndex += (ppu.vram.bytes[lineBaseAddress+0x01] >> effectivePix & 1) << 1
				colorIndex += (ppu.vram.bytes[lineBaseAddress+0x10] >> effectivePix & 1) << 2
				colorIndex += (ppu.vram.bytes[lineBaseAddress+0x11] >> effectivePix & 1) << 3
				// If color is not tansparent, write color value in the pixel
				if colorIndex != 0 {
					colorAddress := 2 * (128 + 16*sprite.paletteIndex + uint16(colorIndex))
					pixels[(sprite.x+8*effectiveTile+effectivePix)%HMax] = pixel{
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
