package ppu

type display struct {
	brightness  uint8 // Display brightness
	forceBlank  bool  // If true, force screen to blank
	vScanning   bool  // If true, interlace mode
	objVDisplay bool  // \
	bgVDisplay  bool  //  \
	hPseudoMode bool  //    Some dank parameters
	ExtBgMode   bool  //  /
	ExtSynchro  bool  // /
}

// 2100h - INIDISP - Display Control 1 (W)
func (ppu *PPU) inidisp(data uint8) uint8 {
	ppu.display.brightness = data & 0x0F
	ppu.display.forceBlank = data&0x80 != 0
	return 0
}

// 212Ch - TM - Main Screen Designation (W)
func (ppu *PPU) tm(data uint8) uint8 {
	for i := 0; i < 4; i++ {
		ppu.backgroundData.bg[i].mainScreen = data&0x01 != 0
		data = data >> 1
	}
	ppu.oam.mainScreen = data&0x01 != 0
	return 0
}

// 212Dh - TS - Sub Screen Designation (W)
func (ppu *PPU) sm(data uint8) uint8 {
	for i := 0; i < 4; i++ {
		ppu.backgroundData.bg[i].subScreen = data&0x01 != 0
		data = data >> 1
	}
	ppu.oam.subScreen = data&0x01 != 0
	return 0
}

// 2133h - SETINI - Display Control 2 (W)
func (ppu *PPU) setini(data uint8) uint8 {
	ppu.display.vScanning = data&0x01 != 0
	data = data >> 1
	ppu.display.objVDisplay = data&0x01 != 0
	data = data >> 1
	ppu.display.bgVDisplay = data&0x01 != 0
	data = data >> 1
	ppu.display.hPseudoMode = data&0x01 != 0
	data = data >> 3
	ppu.display.ExtBgMode = data&0x01 != 0
	data = data >> 1
	ppu.display.ExtSynchro = data&0x01 != 0
	return 0
}
