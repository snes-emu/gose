package core

import (
	"image"

	"github.com/snes-emu/gose/log"
	"github.com/snes-emu/gose/render"
)

const WIDTH = 250
const HEIGHT = 250

const TILE_SIZE = 8

func (ppu *PPU) renderLine() {
	if ppu.screen == nil {
		ppu.screen = render.NewScreen(WIDTH, HEIGHT)
	}

	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter < ppu.screen.Height {
		ppu.screen.SetPixelLine(ppu.vCounter, ppu.backdropPixelLine())
		ppu.screen.SetPixelLine(ppu.vCounter, ppu.spritesToPixelLine(ppu.oam.intersectingSprites(ppu.vCounter)))
		//TODO handle background modes and display backgrounds accordingly
		//We only display BG1 for now
		ppu.screen.SetPixelLine(ppu.vCounter, ppu.backgroundToPixelLine(0))
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
		yTile := (ppu.vCounter - sprite.y) / TILE_SIZE

		// Y coordinate of the line in the tile
		y := (ppu.vCounter - sprite.y) % TILE_SIZE

		// Loop over all the tiles contained in the sprite
		for xTile := uint16(0); xTile < sprite.hSize/TILE_SIZE; xTile++ {

			tile := sprite.tileAt(xTile, yTile)

			// Go through all the pixels in the tile line
			for x, color := range ppu.tileRowColor(tile, y) {
				// Only change the pixel if the color is not transparent
				if !color.Transparent {
					lineIdx := sprite.x + uint16(x) + (TILE_SIZE * xTile)
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

//backgroundToPixelLine the row of pixel of the background bgIndex that intersects with vCounter
//TODO: vertical flip
//TODO: horizontal flip
func (ppu *PPU) backgroundToPixelLine(bgIndex uint8) []render.Pixel {
	// Initialize pixel line
	pixels := make([]render.Pixel, WIDTH)

	bg := ppu.backgroundData.bg[bgIndex]
	hTileSize, vTileSize := bg.tileSize()

	//Y coordinate of the background tile containing the line
	yBgTile := (ppu.vCounter) / vTileSize

	//Y coordinate of the base tile inside the background tile containing the line
	yTile := (ppu.vCounter - yBgTile*vTileSize) / TILE_SIZE

	//Y coordinate of the line inside the base tile
	y := (ppu.vCounter - yBgTile*vTileSize - yTile*TILE_SIZE)

	//go through the background tiles
	for xBgTile := uint16(0); xBgTile < (uint16(WIDTH))/hTileSize+1; xBgTile++ {

		//get the background tile at these coordinates
		bgTile := ppu.tileFromBackground(bgIndex, xBgTile, yBgTile)

		// Loop over all the base tiles contained in the background tile
		for xTile := uint16(0); xTile < bgTile.hSize/TILE_SIZE; xTile++ {

			tile := bgTile.tileAt(xTile, yTile)

			// Loop over all the pixels in the current tile
			for x, color := range ppu.tileRowColor(tile, y) {
				lineIdx := uint16(x) + xTile*TILE_SIZE + xBgTile*hTileSize
				if !color.Transparent && lineIdx >= 0 && lineIdx < WIDTH {
					pixels[lineIdx] = render.Pixel{
						Color:   color,
						Visible: true,
					}
					if bgTile.priority {
						pixels[lineIdx].Priority = 1
					}

				}
			}
		}
	}

	return pixels
}

func (ppu *PPU) backdropPixelLine() []render.Pixel {
	pixels := make([]render.Pixel, WIDTH)
	backdropPixel := ppu.backdropPixel()
	for i := range pixels {
		pixels[i] = backdropPixel
	}

	return pixels
}

func (ppu *PPU) spriteToImage(sprite sprite) image.Image {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: int(sprite.hSize), Y: int(sprite.hSize)},
	})

	// Loop over all the tiles contained in the sprite
	for yTile := uint16(0); yTile < sprite.vSize/TILE_SIZE; yTile++ {
		for xTile := uint16(0); xTile < sprite.hSize/TILE_SIZE; xTile++ {
			tile := sprite.tileAt(xTile, yTile)

			// Loop over all the pixels in the current tile
			for y := uint16(0); y < TILE_SIZE; y++ {
				for x, color := range ppu.tileRowColor(tile, y) {

					if !color.Transparent {
						img.Set(int(xTile*TILE_SIZE+uint16(x)), int(yTile*TILE_SIZE+y), color)
					}
				}
			}
		}
	}

	return img
}

// Sprites returns all the sprites in image.Image format
func (ppu *PPU) Sprites() []image.Image {
	sprites := ppu.oam.allSprites()
	images := make([]image.Image, len(sprites))
	for i, sprite := range sprites {
		images[i] = ppu.spriteToImage(sprite)
	}
	return images
}

func (ppu *PPU) bgToImage(bgIndex uint8) image.Image {
	bg := ppu.backgroundData.bg[bgIndex]

	//create an image to fit the background
	sizeInTile := uint16(64)
	hTileSize, vTileSize := bg.tileSize()
	hSize := sizeInTile * hTileSize
	vSize := sizeInTile * vTileSize
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: int(hSize), Y: int(vSize)},
	})

	//fill the image

	//go through the background tiles
	for yBgTile := uint16(0); yBgTile < uint16(sizeInTile); yBgTile++ {
		for xBgTile := uint16(0); xBgTile < uint16(sizeInTile); xBgTile++ {

			//get the background tile at these coordinates
			bgTile := ppu.tileFromBackground(bgIndex, xBgTile, yBgTile)

			// Loop over all the base tiles contained in the background tile
			for yTile := uint16(0); yTile < bgTile.vSize/TILE_SIZE; yTile++ {
				for xTile := uint16(0); xTile < bgTile.hSize/TILE_SIZE; xTile++ {

					tile := bgTile.tileAt(xTile, yTile)

					// Loop over all the pixels in the current tile
					for y := uint16(0); y < bgTile.vSize; y++ {
						for x, color := range ppu.tileRowColor(tile, y) {
							if !color.Transparent {
								img.Set(int(xBgTile*bgTile.hSize+xTile*TILE_SIZE+uint16(x)), int(yBgTile*bgTile.vSize+yTile*TILE_SIZE+uint16(y)), color)
							}
						}
					}

				}
			}

		}
	}

	return img
}

func (ppu *PPU) Backgrounds() []image.Image {
	images := make([]image.Image, 0, 4)
	for _, bg := range ppu.validBackgrounds() {
		images = append(images, ppu.bgToImage(bg))
	}

	return images
}
