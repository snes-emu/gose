package ppu

// 2130 - CGWSEL - Color Math Control Register A (W)
func (ppu *PPU) cgwsel(data uint8) uint8 {
	ppu.mainScreenBlack = (data & 0xc0) >> 6
	ppu.colorMathEnable = (data & 0x30) >> 4
	ppu.enableSubscreen = (data & 0x2) != 0
	ppu.directColor = (data & 0x1) != 0
	return 0
}

// 2131 - CGADSUB - Color Math Control Register B (W)

// 2132 - COLDATA - Color Math Sub Screen Backdrop Color (W)
