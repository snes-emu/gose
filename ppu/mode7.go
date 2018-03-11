package ppu

type m7 struct {
	verticalFlip                   bool   // Vertical flip flag used in mode7
	horizontalFlip                 bool   // Horizontal flip flag used in mode7
	screenOver                     uint8  // Mode 7 screen over variable (possible values are 0,1,2 or 3)
	cache                          uint16 // Mode 7 cache value used in various registers like M7A or M7HOFS
	aParam, bParam, cParam, dParam uint16 // Rotation/scaling parameters used in mode 7
	hofsParam, vofsParam           uint16 // Mode 7 horizontal and vertical scroll offset parameters
	xParam, yParam                 uint16 // Mode 7 Center Coordinate parameters
}

// 211Ah - M7SEL - Rotation/Scaling Mode Settings (W)
func (ppu *PPU) m7sel(data uint8) uint8 {
	ppu.m7.screenOver = data & 0xc0 >> 6
	ppu.m7.verticalFlip = data&0x2 != 0
	ppu.m7.horizontalFlip = data&0x1 != 0
	return 0
}

// 211B - M7A - Rotation/Scaling Parameter A (and Maths 16bit operand) (W)
func (ppu *PPU) m7a(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.aParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}

// 211C - M7B - Rotation/Scaling Parameter B (and Maths 8bit operand) (W)
func (ppu *PPU) m7b(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.bParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}

// 211D - M7C - Rotation/Scaling Parameter C (W)
func (ppu *PPU) m7c(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.cParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}

// 211E - M7D - Rotation/Scaling Parameter D (W)
func (ppu *PPU) m7d(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.dParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}

// 210D - M7HOFS - Mode 7 Horizontal Scroll (X) (W)
func (ppu *PPU) m7hofs(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.hofsParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}

// 210E - M7VOFS - Mode 7 Vertical Scroll (Y) (W)
func (ppu *PPU) m7vofs(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.vofsParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}

// 211F - M7X - Rotation/Scaling Center Coordinate X (W)
func (ppu *PPU) m7x(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.xParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}

// 2120 - M7Y - Rotation/Scaling Center Coordinate Y (W)
func (ppu *PPU) m7y(data uint8) uint8 {
	data16 := uint16(data)
	ppu.m7.yParam = (data16 << 8) | ppu.m7.cache
	ppu.m7.cache = data16
	return 0
}
