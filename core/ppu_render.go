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

	if !ppu.cpu.ioMemory.vBlankNMIFlag {
		ppu.cpu.doHDMA()
	}

	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter < ppu.screen.Height {
		ppu.paintPixelLine()
	}

	if ppu.vCounter == ppu.VDisplay()+1 {
		screen := ppu.screen
		if ppu.display.forceBlank {
			screen = ppu.blackScreen
		}
		ppu.renderer.Render(screen)
		log.Debug("VBlank")
		ppu.cpu.enterVblank()
	}

	if ppu.vCounter == 0 {
		log.Debug("End of VBlank")
		ppu.cpu.leavVblank()
		ppu.cpu.reloadHDMA()
	}
}

func compositLayers(base []render.Pixel, top []render.Pixel, priority uint8) {
	for i, pix := range top[:len(base)] {
		if pix.Visible && pix.Priority == priority {
			base[i] = pix
		}
	}
}

func (ppu *PPU) paintMainScreen(priority uint8, layer []render.Pixel, bg uint8) {
	var paintMain bool
	if bg < 4 {
		paintMain = ppu.backgroundData.bg[bg].mainScreen
	} else {
		paintMain = ppu.oam.mainScreen
	}
	if paintMain {
		compositLayers(ppu.mainScreen, layer, priority)
	}
}

func (ppu *PPU) paintSubScreen(priority uint8, layer []render.Pixel, bg uint8) {
	var paintSub bool
	if bg < 4 {
		paintSub = ppu.backgroundData.bg[bg].subScreen
	} else {
		paintSub = ppu.oam.subScreen
	}
	if paintSub {
		compositLayers(ppu.subScreen, layer, priority)
	}
}

// Mode0    Mode1    Mode2    Mode3    Mode4    Mode5    Mode6    Mode7
// -        BG3.1a   -        -        -        -        -        -
// OBJ.3    OBJ.3    OBJ.3    OBJ.3    OBJ.3    OBJ.3    OBJ.3    OBJ.3
// BG1.1    BG1.1    BG1.1    BG1.1    BG1.1    BG1.1    BG1.1    -
// BG2.1    BG2.1    -        -        -        -        -        -
// OBJ.2    OBJ.2    OBJ.2    OBJ.2    OBJ.2    OBJ.2    OBJ.2    OBJ.2
// BG1.0    BG1.0    BG2.1    BG2.1    BG2.1    BG2.1    -        BG2.1p
// BG2.0    BG2.0    -        -        -        -        -        -
// OBJ.1    OBJ.1    OBJ.1    OBJ.1    OBJ.1    OBJ.1    OBJ.1    OBJ.1
// BG3.1    BG3.1b   BG1.0    BG1.0    BG1.0    BG1.0    BG1.0    BG1
// BG4.1    -        -        -        -        -        -        -
// OBJ.0    OBJ.0    OBJ.0    OBJ.0    OBJ.0    OBJ.0    OBJ.0    OBJ.0
// BG3.0    BG3.0a   BG2.0    BG2.0    BG2.0    BG2.0    -        BG2.0p
// BG4.0    BG3.0b   -        -        -        -        -        -
// Backdrop Backdrop Backdrop Backdrop Backdrop Backdrop Backdrop Backdrop

func (ppu *PPU) paintPixelLine() {
	ppu.backdropPixelLine()
	sprites := ppu.spritesToPixelLine(ppu.oam.intersectingSprites(ppu.vCounter))
	vbg := ppu.validBackgrounds()
	backgrounds := make([][]render.Pixel, 4)
	for _, bg := range vbg {
		backgrounds[bg] = ppu.backgroundToPixelLine(bg)
	}

	switch ppu.backgroundData.screenMode {
	case 0:
		ppu.paintSubScreen(0, backgrounds[3], 3)
		ppu.paintSubScreen(0, backgrounds[2], 2)
		ppu.paintSubScreen(0, sprites, 4)
		ppu.paintSubScreen(1, backgrounds[3], 3)
		ppu.paintSubScreen(1, backgrounds[2], 2)
		ppu.paintSubScreen(1, sprites, 4)
		ppu.paintSubScreen(0, backgrounds[1], 1)
		ppu.paintSubScreen(0, backgrounds[0], 0)
		ppu.paintSubScreen(2, sprites, 4)
		ppu.paintSubScreen(1, backgrounds[1], 1)
		ppu.paintSubScreen(1, backgrounds[0], 0)
		ppu.paintSubScreen(3, sprites, 4)

	case 1:
		ppu.paintSubScreen(0, backgrounds[2], 2)
		ppu.paintSubScreen(0, sprites, 4)
		if !ppu.backgroundData.bg3Priority {
			ppu.paintSubScreen(1, backgrounds[2], 2)
		}
		ppu.paintSubScreen(1, sprites, 4)
		ppu.paintSubScreen(0, backgrounds[1], 1)
		ppu.paintSubScreen(0, backgrounds[0], 0)
		ppu.paintSubScreen(2, sprites, 4)
		ppu.paintSubScreen(1, backgrounds[1], 1)
		ppu.paintSubScreen(1, backgrounds[0], 0)
		ppu.paintSubScreen(3, sprites, 4)
		if ppu.backgroundData.bg3Priority {
			ppu.paintSubScreen(1, backgrounds[2], 2)
		}

	case 2, 3, 4, 5:
		ppu.paintSubScreen(0, backgrounds[1], 0)
		ppu.paintSubScreen(0, sprites, 4)
		ppu.paintSubScreen(0, backgrounds[0], 0)
		ppu.paintSubScreen(1, sprites, 4)
		ppu.paintSubScreen(1, backgrounds[1], 1)
		ppu.paintSubScreen(2, sprites, 4)
		ppu.paintSubScreen(1, backgrounds[0], 0)
		ppu.paintSubScreen(3, sprites, 4)

	case 6:
		ppu.paintSubScreen(0, sprites, 4)
		ppu.paintSubScreen(0, backgrounds[0], 0)
		ppu.paintSubScreen(1, sprites, 4)
		ppu.paintSubScreen(2, sprites, 4)
		ppu.paintSubScreen(1, backgrounds[0], 0)
		ppu.paintSubScreen(3, sprites, 4)

	case 7:

		//TODO per pixel priority
		ppu.paintSubScreen(0, backgrounds[1], 1)
		ppu.paintSubScreen(0, sprites, 4)
		ppu.paintSubScreen(0, backgrounds[0], 0)
		ppu.paintSubScreen(1, sprites, 4)
		ppu.paintSubScreen(1, backgrounds[1], 1)
		ppu.paintSubScreen(2, sprites, 4)
		ppu.paintSubScreen(3, sprites, 4)
	}

	if ppu.colorMath.backdrop {
		ppu.applyColorMath(ppu.mainScreen)
	}
	if ppu.colorMath.obj {
		ppu.applyColorMath(sprites)
	}
	for _, bg := range vbg {
		if ppu.backgroundData.bg[bg].colorMath {
			ppu.applyColorMath(backgrounds[bg])
		}
	}

	switch ppu.backgroundData.screenMode {
	case 0:
		ppu.paintMainScreen(0, backgrounds[3], 3)
		ppu.paintMainScreen(0, backgrounds[2], 2)
		ppu.paintMainScreen(0, sprites, 4)
		ppu.paintMainScreen(1, backgrounds[3], 3)
		ppu.paintMainScreen(1, backgrounds[2], 2)
		ppu.paintMainScreen(1, sprites, 4)
		ppu.paintMainScreen(0, backgrounds[1], 1)
		ppu.paintMainScreen(0, backgrounds[0], 0)
		ppu.paintMainScreen(2, sprites, 4)
		ppu.paintMainScreen(1, backgrounds[1], 1)
		ppu.paintMainScreen(1, backgrounds[0], 0)
		ppu.paintMainScreen(3, sprites, 4)

	case 1:
		ppu.paintMainScreen(0, backgrounds[2], 2)
		ppu.paintMainScreen(0, sprites, 4)
		if !ppu.backgroundData.bg3Priority {
			ppu.paintMainScreen(1, backgrounds[2], 2)
		}
		ppu.paintMainScreen(1, sprites, 4)
		ppu.paintMainScreen(0, backgrounds[1], 1)
		ppu.paintMainScreen(0, backgrounds[0], 0)
		ppu.paintMainScreen(2, sprites, 4)
		ppu.paintMainScreen(1, backgrounds[1], 1)
		ppu.paintMainScreen(1, backgrounds[0], 0)
		ppu.paintMainScreen(3, sprites, 4)
		if ppu.backgroundData.bg3Priority {
			ppu.paintMainScreen(1, backgrounds[2], 2)
		}

	case 2, 3, 4, 5:
		ppu.paintMainScreen(0, backgrounds[1], 0)
		ppu.paintMainScreen(0, sprites, 4)
		ppu.paintMainScreen(0, backgrounds[0], 0)
		ppu.paintMainScreen(1, sprites, 4)
		ppu.paintMainScreen(1, backgrounds[1], 1)
		ppu.paintMainScreen(2, sprites, 4)
		ppu.paintMainScreen(1, backgrounds[0], 0)
		ppu.paintMainScreen(3, sprites, 4)

	case 6:
		ppu.paintMainScreen(0, sprites, 4)
		ppu.paintMainScreen(0, backgrounds[0], 0)
		ppu.paintMainScreen(1, sprites, 4)
		ppu.paintMainScreen(2, sprites, 4)
		ppu.paintMainScreen(1, backgrounds[0], 0)
		ppu.paintMainScreen(3, sprites, 4)

	case 7:

		//TODO per pixel priority
		ppu.paintMainScreen(0, backgrounds[1], 1)
		ppu.paintMainScreen(0, sprites, 4)
		ppu.paintMainScreen(0, backgrounds[0], 0)
		ppu.paintMainScreen(1, sprites, 4)
		ppu.paintMainScreen(1, backgrounds[1], 1)
		ppu.paintMainScreen(2, sprites, 4)
		ppu.paintMainScreen(3, sprites, 4)
	}

	ppu.screen.SetPixelLine(ppu.vCounter, ppu.mainScreen)
}

// spritesToPixelLine takes the given sprites and outputs a row of pixels that intersects with the vCounter
// TODO: Update ppu status
// TODO: limit to 32 sprites
// TODO: limit to 34 tiles
func (ppu *PPU) spritesToPixelLine(sprites []sprite) []render.Pixel {
	// Initialize pixel line
	pixels := make([]render.Pixel, WIDTH)

	if ppu.oam.priorityBit {
		log.Error("sprite priority rotation not implemented")
	}
	for i := len(sprites) - 1; i >= 0; i-- {
		sprite := sprites[i]
		// Y coordinate of the tile containing the line
		yTile := (ppu.vCounter - sprite.y) / TILE_SIZE
		if sprite.vFlip {
			yTile = sprite.vSize/TILE_SIZE - yTile - 1
		}

		// Y coordinate of the line in the tile
		y := (ppu.vCounter - sprite.y) % TILE_SIZE
		if sprite.vFlip {
			y = TILE_SIZE - y - 1
		}

		// Loop over all the tiles contained in the sprite
		for xTile := uint16(0); xTile < sprite.hSize/TILE_SIZE; xTile++ {

			tile := sprite.tileAt(xTile, yTile)

			// Go through all the pixels in the tile line
			for x, color := range ppu.tileRowColor(tile, y) {
				// Only change the pixel if the color is not transparent
				if !color.Transparent {
					xp := uint16(x) + (TILE_SIZE * xTile)
					if sprite.hFlip {
						xp = sprite.hSize - xp - 1
					}
					lineIdx := sprite.x + xp
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
func (ppu *PPU) backgroundToPixelLine(bgIndex uint8) []render.Pixel {
	// Initialize pixel line
	pixels := make([]render.Pixel, WIDTH)

	bg := ppu.backgroundData.bg[bgIndex]
	hTileSize, vTileSize := bg.tileSize()

	//Y coordinate of the background tile containing the line
	yBgTile := (ppu.vCounter + bg.verticalScroll) / vTileSize

	//Y coordinate of the base tile inside the background tile containing the line
	yBaseTile := (ppu.vCounter + bg.verticalScroll - yBgTile*vTileSize) / TILE_SIZE

	//Y coordinate of the line inside the base tile
	yBase := (ppu.vCounter + bg.verticalScroll - yBgTile*vTileSize) % TILE_SIZE

	//go through the background tiles
	hStart := bg.horizontalScroll
	hEnd := hStart + uint16(WIDTH)
	for xBgTile := hStart / hTileSize; xBgTile < hEnd/hTileSize+1; xBgTile++ {

		//get the background tile at these coordinates
		bgTile := ppu.tileFromBackground(bgIndex, xBgTile, yBgTile)
		yTile := yBaseTile
		y := yBase
		if bgTile.vFlip {
			yTile = vTileSize/TILE_SIZE - yTile - 1
			y = TILE_SIZE - y - 1
		}

		// Loop over all the base tiles contained in the background tile
		for xTile := uint16(0); xTile < bgTile.hSize/TILE_SIZE; xTile++ {

			tile := bgTile.tileAt(xTile, yTile)

			// Loop over all the pixels in the current tile
			for x, color := range ppu.tileRowColor(tile, y) {
				xp := uint16(x) + xTile*TILE_SIZE
				if bgTile.hFlip {
					xp = bgTile.hSize - xp - 1
				}
				lineIdx := xBgTile*hTileSize - bg.horizontalScroll + xp
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

func (ppu *PPU) backdropPixelLine() {
	backdropPixel := ppu.backdropPixel()
	subscreenBackdropPixel := ppu.subScreenBackdropPixel()
	for i := range ppu.mainScreen {
		ppu.mainScreen[i] = backdropPixel
		ppu.subScreen[i] = subscreenBackdropPixel
	}

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
								img.Set(
									int(xBgTile*bgTile.hSize+xTile*TILE_SIZE+uint16(x)),
									int(yBgTile*bgTile.vSize+yTile*TILE_SIZE+uint16(y)),
									color,
								)
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
