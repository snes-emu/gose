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
func (ppu *PPU) cgadsub(data uint8) uint8 {
	if (data & 0x80) != 0 {
		ppu.colorMathOpSign = -1
	} else {
		ppu.colorMathOpSign = 1
	}

	ppu.colorMathDiv2 = (data & 0x40) != 0

	ppu.bg[3].colorMath = (data & 0x8) != 0
	ppu.bg[2].colorMath = (data & 0x4) != 0
	ppu.bg[1].colorMath = (data & 0x2) != 0
	ppu.bg[0].colorMath = (data & 0x1) != 0

	ppu.colorMathBackdrop = (data & 0x20) != 0
	ppu.colorMathObj = (data & 0x10) != 0

	return 0
}

// 2132 - COLDATA - Color Math Sub Screen Backdrop Color (W)
func (ppu *PPU) coldata(data uint8) uint8 {
	intensity := data & 0x1f

	if (data & 0x80) != 0 {
		ppu.colorBlue = intensity
	}
	if (data & 0x40) != 0 {
		ppu.colorGreen = intensity
	}
	if (data & 0x20) != 0 {
		ppu.colorRed = intensity
	}
	return 0
}
