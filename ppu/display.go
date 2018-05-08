package ppu

type display struct {
	brightness  uint8 // Display brightness
	forceBlank  bool  // If true, force screen to blank
	vScanning   bool  // If true, interlace mode
	objVDisplay bool  // If true, obj interlace mode, sprites will appear half-sized
	bgVDisplay  bool  // (0=224 Lines, 1=239 Lines) (for NTSC/PAL)
	hPseudoMode bool  // Horizontal pseudo mode
	ExtBgMode   bool  // mode 7 extra background mode
	ExtSynchro  bool  // usually 0, used with sfx chip
}

// 2100h - INIDISP - Display Control 1 (W)
func (ppu *PPU) inidisp(data uint8) {
	ppu.display.brightness = data & 0x0F
	ppu.display.forceBlank = data&0x80 != 0
}

// 212Ch - TM - Main Screen Designation (W)
func (ppu *PPU) tm(data uint8) {
	for i := uint8(0); i < 4; i++ {
		ppu.backgroundData.bg[i].mainScreen = data&(1<<i) != 0
	}
	ppu.oam.mainScreen = data&0x01 != 0
}

// 212Dh - TS - Sub Screen Designation (W)
func (ppu *PPU) ts(data uint8) {
	for i := uint8(0); i < 4; i++ {
		ppu.backgroundData.bg[i].subScreen = data&(1<<i) != 0
		data = data >> 1
	}
	ppu.oam.subScreen = data&0x01 != 0
}

// 2133h - SETINI - Display Control 2 (W)
func (ppu *PPU) setini(data uint8) {
	ppu.display.vScanning = data&0x01 != 0
	ppu.display.objVDisplay = data&0x02 != 0
	ppu.display.bgVDisplay = data&0x04 != 0
	ppu.display.hPseudoMode = data&0x08 != 0
	ppu.display.ExtBgMode = data&0x40 != 0
	ppu.display.ExtSynchro = data&0x80 != 0
}
