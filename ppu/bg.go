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
	colorMath          bool   // Flag to control colors on the BG (False: Display RAW Main Screen as such (without math), True: Apply math on Mainscreen)
}

// 2105h - BGMODE - BG Mode and BG Character Size (W)
func (ppu *PPU) bgmode(data uint8) uint8 {
	ppu.bgScreenMode = data & 7
	ppu.bg[2].Priority = data&8 != 0
	for i := uint8(0); i < 4; i++ {
		ppu.bg[i].TileSize = data&(1<<(4+i)) != 0
	}
	return 0

}

// 2106h - MOSAIC - Mosaic Size and Mosaic Enable (W)
func (ppu *PPU) mosaic(data uint8) uint8 {
	for i := uint8(0); i < 4; i++ {
		ppu.bg[i].Mosaic = data&(1<<i) != 0
	}
	ppu.mosaicSize = data >> 4
	return 0
}

// 2107h - BG1SC - BG1 Screen Base and Screen Size (W)
func (ppu *PPU) bg1sc(data uint8) uint8 {
	ppu.bg[0].screenSize = data & 3
	ppu.bg[0].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
	return 0
}

// 2108h - BG2SC - BG2 Screen Base and Screen Size (W)
func (ppu *PPU) bg2sc(data uint8) uint8 {
	ppu.bg[1].screenSize = data & 3
	ppu.bg[1].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
	return 0
}

// 2109h - BG3SC - BG3 Screen Base and Screen Size (W)
func (ppu *PPU) bg3sc(data uint8) uint8 {
	ppu.bg[2].screenSize = data & 3
	ppu.bg[2].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
	return 0
}

// 210Ah - BG4SC - BG4 Screen Base and Screen Size (W)
func (ppu *PPU) bg4sc(data uint8) uint8 {
	ppu.bg[3].screenSize = data & 3
	ppu.bg[3].tileMapBaseAddress = uint16(data&^uint8(3)) << 8
	return 0
}

// 210Bh/210Ch - BG12NBA/BG34NBA - BG Character Data Area Designation (W)
func (ppu *PPU) bg12nba(data uint8) uint8 {
	ppu.bg[0].tileSetBaseAddress = uint16(data&0x0F) << 12
	ppu.bg[1].tileSetBaseAddress = uint16(data&0xF0) << 8
	return 0
}

func (ppu *PPU) bg34nba(data uint8) uint8 {
	ppu.bg[2].tileSetBaseAddress = uint16(data&0x0F) << 12
	ppu.bg[3].tileSetBaseAddress = uint16(data&0xF0) << 8
	return 0
}

// 210Dh - BG1HOFS - BG1 Horizontal Scroll (X) (W)
func (ppu *PPU) bg1hofs(data uint8) uint8 {
	ppu.bg[0].horizontalScroll = uint16(data)<<8 | uint16((ppu.bgScrollPrev1 &^ 7)) | uint16(ppu.bgScrollPrev2&7)
	ppu.bgScrollPrev1 = data
	ppu.bgScrollPrev2 = data
	return 0
}

// 210Eh - BG1VOFS - BG1 Vertical Scroll (Y) (W)
func (ppu *PPU) bg1vofs(data uint8) uint8 {
	ppu.bg[0].horizontalScroll = uint16(data)<<8 | uint16(ppu.bgScrollPrev1)
	ppu.bgScrollPrev1 = data
	return 0
}

// 210Fh - BG2HOFS - BG2 Horizontal Scroll (X) (W)
func (ppu *PPU) bg2hofs(data uint8) uint8 {
	ppu.bg[1].horizontalScroll = uint16(data)<<8 | uint16((ppu.bgScrollPrev1 &^ 7)) | uint16(ppu.bgScrollPrev2&7)
	ppu.bgScrollPrev1 = data
	ppu.bgScrollPrev2 = data
	return 0
}

// 2110h - BG2VOFS - BG2 Vertical Scroll (Y) (W)
func (ppu *PPU) bg2vofs(data uint8) uint8 {
	ppu.bg[1].horizontalScroll = uint16(data)<<8 | uint16(ppu.bgScrollPrev1)
	ppu.bgScrollPrev1 = data
	return 0
}

// 2111h - BG3HOFS - BG3 Horizontal Scroll (X) (W)
func (ppu *PPU) bg3hofs(data uint8) uint8 {
	ppu.bg[2].horizontalScroll = uint16(data)<<8 | uint16((ppu.bgScrollPrev1 &^ 7)) | uint16(ppu.bgScrollPrev2&7)
	ppu.bgScrollPrev1 = data
	ppu.bgScrollPrev2 = data
	return 0
}

// 2112h - BG3VOFS - BG3 Vertical Scroll (Y) (W)
func (ppu *PPU) bg3vofs(data uint8) uint8 {
	ppu.bg[2].horizontalScroll = uint16(data)<<8 | uint16(ppu.bgScrollPrev1)
	ppu.bgScrollPrev1 = data
	return 0
}

// 2113h - BG4HOFS - BG4 Horizontal Scroll (X) (W)
func (ppu *PPU) bg4hofs(data uint8) uint8 {
	ppu.bg[3].horizontalScroll = uint16(data)<<8 | uint16((ppu.bgScrollPrev1 &^ 7)) | uint16(ppu.bgScrollPrev2&7)
	ppu.bgScrollPrev1 = data
	ppu.bgScrollPrev2 = data
	return 0
}

// 2114h - BG4VOFS - BG4 Vertical Scroll (Y) (W)
func (ppu *PPU) bg4vofs(data uint8) uint8 {
	ppu.bg[3].horizontalScroll = uint16(data)<<8 | uint16(ppu.bgScrollPrev1)
	ppu.bgScrollPrev1 = data
	return 0
}
