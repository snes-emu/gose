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
			for tile := uint16(sprite.y/8) << 4; tile < uint16(sprite.hSize)/8+uint16(sprite.y/8)<<4; tile++ {
				// Go through all pixel in the line
				for pixel := uint8(sprite.y % 8); pixel < uint8(sprite.y%8)+8; pixel++ {
					// Get the color index of the pixel
					colorIndex := uint8(0)
					colorIndex += ppu.vram.bytes[sprite.tileAddress+tile<<5] & (1 << pixel) >> pixel
					colorIndex += ppu.vram.bytes[sprite.tileAddress+tile<<5+0x01] & (1 << pixel) >> pixel << 1
					colorIndex += ppu.vram.bytes[sprite.tileAddress+tile<<5+0x10] & (1 << pixel) >> pixel << 2
					colorIndex += ppu.vram.bytes[sprite.tileAddress+tile<<5+0x11] & (1 << pixel) >> pixel << 3
					// If color is not tansparent, write color value in the pixel
					if colorIndex != 0 {
						pixels[sprite.x+8*uint16(tile)+uint16(pixel)].bgr = utils.JoinUint16(ppu.cgram.bytes[2*(128+16*sprite.paletteIndex+uint16(colorIndex))], ppu.cgram.bytes[2*(128+16*sprite.paletteIndex+uint16(colorIndex))+1])
						pixels[sprite.x+8*uint16(tile)+uint16(pixel)].visible = true
					}
				}

			}
		}
	}
	return pixels
}
