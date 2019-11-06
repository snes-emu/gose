package core

import (
	"fmt"
	"image"
	"image/png"
	"os"

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
		ppu.screen.SetPixelLine(ppu.vCounter, ppu.spritesToPixelLine(ppu.oam.intersectingSprites(ppu.vCounter)))
	}

	if ppu.vCounter == ppu.VDisplay()+1 {
		ppu.renderer.Render(ppu.screen)
		log.Debug("VBlank")
		ppu.cpu.enterVblank()

		for i, img := range ppu.Backgrounds() {
			f, err := os.Create(fmt.Sprintf("/tmp/bg%d.png", i))
			if err != nil {
				panic(err)
			}
			err = png.Encode(f, img)
			if err != nil {
				panic(err)
			}
		}
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

			// Address of the current tile in the VRAM
			tileAddress := sprite.tileAddr(xTile, yTile)

			// Go through all the pixels in the tile line
			for x, color := range ppu.tileSpriteRowColor(tileAddress, y, sprite.palette) {
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

func (ppu *PPU) spriteToImage(sprite sprite) image.Image {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: int(sprite.hSize), Y: int(sprite.hSize)},
	})

	// Loop over all the tiles contained in the sprite
	for yTile := uint16(0); yTile < sprite.vSize/TILE_SIZE; yTile++ {
		for xTile := uint16(0); xTile < sprite.hSize/TILE_SIZE; xTile++ {
			// Address of the current tile in the VRAM
			tileAddress := sprite.tileAddr(xTile, yTile)

			// Loop over all the pixels in the current tile
			for y := uint16(0); y < TILE_SIZE; y++ {
				for x, color := range ppu.tileSpriteRowColor(tileAddress, y, sprite.palette) {

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
	sizeInTile := uint16(64)
	hTileSize, vTileSize := bg.tileSize()

	hSize := sizeInTile * hTileSize
	vSize := sizeInTile * vTileSize
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: int(hSize), Y: int(vSize)},
	})
	for yTile := uint16(0); yTile < uint16(sizeInTile); yTile++ {
		for xTile := uint16(0); xTile < uint16(sizeInTile); xTile++ {
			tile := ppu.tileFromBackground(bgIndex, xTile, yTile)
			for y := uint16(0); y < tile.vSize; y++ {
				for x, color := range ppu.tileRowColor(uint16(tile.firstTileAddr), uint16(tile.colorDepth), hTileSize, y, tile.palette) {
					if !color.Transparent {
						img.Set(int(xTile*hTileSize+uint16(x)), int(yTile*vTileSize+uint16(y)), color)
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
