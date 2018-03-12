package ppu

import (
	"github.com/snes-emu/gose/utils"
)

type pixel struct {
	bgr     uint16
	visible bool
}

func (ppu PPU) renderSpriteLine() [HMax]pixel {
	// Initialize pixel line
	var pixels [HMax]pixel
	// Go through all sprites in reverse order
	for i := uint16(127); i >= 0; i-- {
		sprite := ppu.decodeSprite(i)
		// Check if line intersects sprite
		if ppu.vCounter >= sprite.y && ppu.vCounter < sprite.y+uint16(sprite.vSize) {
			// Go through all tiles containing the line
			baseTileIndex := uint16((ppu.vCounter-sprite.y)/8) << 4 // Y coordinate of the tile containing the line
			for tile := baseTileIndex; tile < baseTileIndex+uint16(sprite.hSize)/8; tile++ {
				// Go through all pixel in the line
				for pixel := uint8(0); pixel < 8; pixel++ {
					yPixel := uint16(ppu.vCounter-sprite.y) % 8                // Y coordinate of the line in the tile
					lineBaseAddress := sprite.tileAddress + tile<<5 + 2*yPixel // Base address of the color planes of the pixel line in VRAM
					// Get the color index of the pixel
					colorIndex := uint8(0)
					colorIndex += ppu.vram.bytes[lineBaseAddress] & (1 << pixel) >> pixel
					colorIndex += ppu.vram.bytes[lineBaseAddress+0x01] & (1 << pixel) >> pixel << 1
					colorIndex += ppu.vram.bytes[lineBaseAddress+0x10] & (1 << pixel) >> pixel << 2
					colorIndex += ppu.vram.bytes[lineBaseAddress+0x11] & (1 << pixel) >> pixel << 3
					// If color is not tansparent, write color value in the pixel
					if colorIndex != 0 {
						colorAddress := 2 * (128 + 16*sprite.paletteIndex + uint16(colorIndex))
						pixels[sprite.x+8*tile+uint16(pixel)].bgr = utils.JoinUint16(ppu.cgram.bytes[colorAddress], ppu.cgram.bytes[colorAddress+1])
						pixels[sprite.x+8*tile+uint16(pixel)].visible = true
					}
				}

			}
		}
	}
	return pixels
}
