package ppu

// 211Ah - M7SEL - Rotation/Scaling Mode Settings (W)
func (ppu *PPU) m7sel(data uint8) uint8 {
	ppu.m7ScreenOver = data & 0xc0 >> 6
	ppu.m7VerticalFlip = data&0x2 != 0
	ppu.m7HorizontalFlip = data&0x1 != 0
	return 0
}

// 211B - M7A - Rotation/Scaling Parameter A (and Maths 16bit operand) (W)
func (ppu *PPU) m7a(data uint8) uint8 {
	ppu.m7aParam = (ppu.m7aParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7aParam
	return 0
}

// 211C - M7B - Rotation/Scaling Parameter B (and Maths 8bit operand) (W)
func (ppu *PPU) m7b(data uint8) uint8 {
	ppu.m7bParam = (ppu.m7bParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7bParam
	return 0
}

// 211D - M7C - Rotation/Scaling Parameter C (W)
func (ppu *PPU) m7c(data uint8) uint8 {
	ppu.m7cParam = (ppu.m7cParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7cParam
	return 0
}

// 211E - M7D - Rotation/Scaling Parameter D (W)
func (ppu *PPU) m7d(data uint8) uint8 {
	ppu.m7dParam = (ppu.m7dParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7dParam
	return 0
}
