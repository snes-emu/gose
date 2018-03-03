package ppu

type window struct {
	left  uint8
	right uint8
}

// 2126h - WH0 - Window 1 Left Position (X1) (W)
func (ppu *PPU) wh0(data uint8) uint8 {
	ppu.window[0].left = data
	return 0
}

// 2127h - WH1 - Window 1 Right Position (X2) (W)
func (ppu *PPU) wh1(data uint8) uint8 {
	ppu.window[0].right = data
	return 0
}

// 2128h - WH2 - Window 2 Left Position (X1) (W)
func (ppu *PPU) wh2(data uint8) uint8 {
	ppu.window[1].left = data
	return 0
}

// 2129h - WH3 - Window 2 Right Position (X2) (W)
func (ppu *PPU) wh3(data uint8) uint8 {
	ppu.window[1].right = data
	return 0
}
