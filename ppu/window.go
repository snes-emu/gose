package ppu

type window struct {
	left  uint8
	right uint8
}

// 2126h - WH0 - Window 1 Left Position (X1) (W)
func (ppu *PPU) wh0(data uint8) {
	ppu.window[0].left = data
}

// 2127h - WH1 - Window 1 Right Position (X2) (W)
func (ppu *PPU) wh1(data uint8) {
	ppu.window[0].right = data
}

// 2128h - WH2 - Window 2 Left Position (X1) (W)
func (ppu *PPU) wh2(data uint8) {
	ppu.window[1].left = data
}

// 2129h - WH3 - Window 2 Right Position (X2) (W)
func (ppu *PPU) wh3(data uint8) {
	ppu.window[1].right = data
}

// 2123h - W12SEL - Window BG1/BG2 Mask Settings (W)
func (ppu *PPU) w12sel(data uint8) {
	ppu.backgroundData.bg[0].windowMask1 = data & 0x3
	ppu.backgroundData.bg[0].windowMask2 = (data >> 2) & 0x3
	ppu.backgroundData.bg[1].windowMask1 = (data >> 4) & 0x3
	ppu.backgroundData.bg[1].windowMask2 = (data >> 6) & 0x3
}

// 2124h - W34SEL - Window BG3/BG4 Mask Settings (W)
func (ppu *PPU) w34sel(data uint8) {
	ppu.backgroundData.bg[2].windowMask1 = data & 0x3
	ppu.backgroundData.bg[2].windowMask2 = (data >> 2) & 0x3
	ppu.backgroundData.bg[3].windowMask1 = (data >> 4) & 0x3
	ppu.backgroundData.bg[3].windowMask2 = (data >> 6) & 0x3
}

// 2125h - WOBJSEL - Window OBJ/MATH Mask Settings (W)
func (ppu *PPU) wobjsel(data uint8) {
	ppu.oam.windowMask1 = data & 0x3
	ppu.oam.windowMask2 = (data >> 2) & 0x3
	ppu.colorMath.windowMask1 = (data >> 4) & 0x3
	ppu.colorMath.windowMask2 = (data >> 6) & 0x3
}

// 212Ah/212Bh - WBGLOG/WOBJLOG - Window 1/2 Mask Logic (W)
func (ppu *PPU) wbglog(data uint8) {
	for i := 0; i < 4; i++ {
		ppu.backgroundData.bg[i].windowMaskLogic = data & 0x3
		data = data >> 2
	}
}

func (ppu *PPU) wobjlog(data uint8) {
	ppu.oam.windowMaskLogic = data & 0x3
	data = data >> 2
	ppu.colorMath.windowMaskLogic = data & 0x3
}

// 212Eh - TMW - Window Area Main Screen Disable (W)
func (ppu *PPU) tmw(data uint8) {
	for i := 0; i < 4; i++ {
		ppu.backgroundData.bg[i].mainScreenWindow = data&0x1 != 0
		data = data >> 1
	}
	ppu.oam.mainScreenWindow = data&0x1 != 0
}

// 212Fh - TSW - Window Area Sub Screen Disable (W)
func (ppu *PPU) tsw(data uint8) {
	for i := 0; i < 4; i++ {
		ppu.backgroundData.bg[i].subScreenWindow = data&0x1 != 0
		data = data >> 1
	}
	ppu.oam.subScreenWindow = data&0x1 != 0
}
