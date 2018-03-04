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

// 210D - M7HOFS - Mode 7 Horizontal Scroll (X) (W)
func (ppu *PPU) m7hofs(data uint8) uint8 {
	ppu.m7hofsParam = (ppu.m7hofsParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7hofsParam
	return 0
}

// 210E - M7VOFS - Mode 7 Vertical Scroll (Y) (W)
func (ppu *PPU) m7vofs(data uint8) uint8 {
	ppu.m7vofsParam = (ppu.m7vofsParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7vofsParam
	return 0
}

// 211F - M7X - Rotation/Scaling Center Coordinate X (W)
func (ppu *PPU) m7x(data uint8) uint8 {
	ppu.m7xParam = (ppu.m7xParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7xParam
	return 0
}

// 2120 - M7Y - Rotation/Scaling Center Coordinate Y (W)
func (ppu *PPU) m7y(data uint8) uint8 {
	ppu.m7yParam = (ppu.m7yParam << 8) | ppu.m7Cache
	ppu.m7Cache = ppu.m7yParam
	return 0
}
