package ppu

type status struct {
	hCounterLatch  uint16 // Stores the h counter when latched is performed
	vCounterLatch  uint16 // Stores the v counter when latched is performed
	ophctFlip      bool   // H counter latch read flip
	opvctFlip      bool   // V counter latch read flip
	timeOver       bool   // Set when more than 34 sprite tiles are rendrered on one line
	rangeOver      bool   // Set when more than 32 sprites are rendered on one line
	palMode        bool   // PAL mode framerate
	latchedData    bool   // New latched data indicator
	interlaceFrame bool   // Interlace mode current frame
}

func (ppu *PPU) latchCounter() {
	//TODO check if bit 7 of $4201 is set before latching
	ppu.status.hCounterLatch = ppu.HCounter
	ppu.status.vCounterLatch = ppu.VCounter
	ppu.status.latchedData = true
}

// 2137h - SLHV - Latch H/V-Counter by Software (R)
func (ppu *PPU) slhv(data uint8) uint8 {
	ppu.latchCounter()
	return 0
}

// 213Ch - OPHCT - Horizontal Counter Latch (R)
func (ppu *PPU) ophct(data uint8) uint8 {
	var result uint8
	if ppu.status.ophctFlip {
		result = uint8(ppu.status.hCounterLatch >> 8)
	} else {
		result = uint8(ppu.status.hCounterLatch)
	}
	ppu.status.ophctFlip = !ppu.status.ophctFlip
	return result
}

// 213Dh - OPVCT - Vertical Counter Latch (R)
func (ppu *PPU) opvct(data uint8) uint8 {
	var result uint8
	if ppu.status.opvctFlip {
		result = uint8(ppu.status.vCounterLatch >> 8)
	} else {
		result = uint8(ppu.status.vCounterLatch)
	}
	ppu.status.opvctFlip = !ppu.status.opvctFlip
	return result
}

// 213Eh - STAT77 - PPU1 Status and Version Number (R)
func (ppu *PPU) stat77(data uint8) uint8 {
	var result uint8 = 1 // PPU1 5C77 Version Number
	if ppu.status.rangeOver {
		result += 0x40
	}
	if ppu.status.timeOver {
		result += 0x80
	}
	return result
}

// 213Fh - STAT78 - PPU2 Status and Version Number (R)
func (ppu *PPU) stat78(data uint8) uint8 {
	var result uint8 = 2 // PPU2 5C78 Version Number
	if ppu.status.palMode {
		result += 0x10
	}
	if ppu.status.latchedData {
		result += 0x40
	}
	if ppu.status.interlaceFrame {
		result += 0x80
	}
	ppu.status.latchedData = false
	ppu.status.ophctFlip = false
	ppu.status.opvctFlip = false
	return result
}
