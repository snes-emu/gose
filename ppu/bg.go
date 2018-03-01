package ppu

type BG struct {
	tileSize           bool
	mosaic             bool
	size               uint8
	tileMapBaseAddress uint16
	tileSetBaseAddress uint16
	horizontalScroll   uint16
	verticalScroll     uint16
}
