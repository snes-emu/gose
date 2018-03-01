package ppu

// BG stores data about a background
type BG struct {
	TileSize           bool   // false 8x8 tiles, true 16x16 tiles
	Mosaic             bool   // mosaic mode enabled
	Priority           bool   // Only useful for BG3
	screenSize         uint8  // 0=32x32, 1=64x32, 2=32x64, 3=64x64 tiles
	tileMapBaseAddress uint16 // base address for tile map in VRAM
	tileSetBaseAddress uint16 // base address for tile set in VRAM
	horizontalScroll   uint16 // horizontal scroll in pixel
	verticalScroll     uint16 // vertical scroll in pixel
}

// 2105h - BGMODE - BG Mode and BG Character Size (W)
func (ppu *PPU) bgmode(data uint8) {
	ppu.bgScreenMode = data & 7
	ppu.bg[2].Priority = data&8 != 0
	for i := uint8(0); i < 4; i++ {
		ppu.bg[i].TileSize = data&(1<<4+i) != 0
	}

}
