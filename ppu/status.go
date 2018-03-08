package ppu

type status struct {
	hCounterLatch uint16
	vCounterLatch uint16
	ophctFlip     bool
	opvctFlip     bool
}

// 2137h - SLHV - Latch H/V-Counter by Software (R)
func (ppu *PPU) slhv(data uint8) uint8 {
	//TODO check if bit 7 of $4201 is set before latching
	ppu.status.hCounterLatch = ppu.hCounter
	ppu.status.vCounterLatch = ppu.vCounter
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
