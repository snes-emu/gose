package core

import (
	"github.com/snes-emu/gose/log"
	"github.com/snes-emu/gose/render"
)

const WIDTH = 250
const HEIGHT = 250

func (ppu *PPU) renderLine() {
	if ppu.screen == nil {
		ppu.screen = render.NewScreen(WIDTH, HEIGHT)
	}

	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter < ppu.screen.Height {
		ppu.screen.SetPixelLine(ppu.vCounter, ppu.spritesToPixelLine(ppu.oam.intersectingSprites(ppu.vCounter)))
	}

	if ppu.vCounter == ppu.VDisplay()+1 {
		ppu.renderer.Render(ppu.screen)
		log.Debug("VBlank")
		ppu.cpu.enterVblank()
	}

	if ppu.vCounter == 0 {
		log.Debug("End of VBlank")
		ppu.cpu.leavVblank()
	}
}

// spritesToPixelLine takes the given sprites and outputs a row of pixels that intersects with the vCounter
// TODO: vertical flip
// TODO: horizontal flip
// TODO: Update ppu status
// TODO: limit to 32 sprites
// TODO: limit to 34 tiles
// TODO: respect sprite priority
func (ppu *PPU) spritesToPixelLine(sprites []sprite) []render.Pixel {
	// Initialize pixel line
	pixels := make([]render.Pixel, WIDTH)

	for _, sprite := range sprites {
		// Y coordinate of the tile containing the line
		var yTile = (ppu.vCounter - sprite.y) / 8

		// Y coordinate of the line in the tile
		var y = (ppu.vCounter - sprite.y) % 8

		base := yTile << 4

		// Loop over all the tiles contained in the sprite
		for tileNb := uint16(0); tileNb < sprite.hSize/8; tileNb++ {

			// Address of the current tile in the VRAM
			tileAddress := sprite.firstTileAddr + (base+tileNb)<<5

			// Go through all the pixels in the tile line
			for x := uint16(0); x < 8; x++ {
				color := ppu.tileSpriteColor(tileAddress, x, y, sprite.palette)

				// Only change the pixel if the color is not transparent
				if !color.Transparent {
					lineIdx := sprite.x + x + (8 * tileNb)
					pixels[lineIdx%WIDTH] = render.Pixel{
						Color:    color,
						Visible:  true,
						Priority: sprite.priority,
					}
				}
			}
		}
	}

	return pixels
}
