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
		ppu.bg[i].TileSize = data&(1<<(4+i)) != 0
	}

}

// 2106h - MOSAIC - Mosaic Size and Mosaic Enable (W)
func (ppu *PPU) mosaic(data uint8) {
	for i := uint8(0); i < 4; i++ {
		ppu.bg[i].Mosaic = data&(1<<i) != 0
	}
	ppu.mosaicSize = data >> 4
}

// 2107h - BG1SC - BG1 Screen Base and Screen Size (W)
func (ppu *PPU) bg1sc(data uint8) {
	ppu.bg[0].screenSize = data & 3
	ppu.bg[0].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
}

// 2108h - BG2SC - BG2 Screen Base and Screen Size (W)
func (ppu *PPU) bg2sc(data uint8) {
	ppu.bg[1].screenSize = data & 3
	ppu.bg[1].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
}

// 2109h - BG3SC - BG3 Screen Base and Screen Size (W)
func (ppu *PPU) bg3sc(data uint8) {
	ppu.bg[1].screenSize = data & 3
	ppu.bg[1].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
}

// 210Ah - BG4SC - BG4 Screen Base and Screen Size (W)
func (ppu *PPU) bg4sc(data uint8) {
	ppu.bg[1].screenSize = data & 3
	ppu.bg[1].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
}
