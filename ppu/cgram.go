package ppu

// 2121 - Color index (0..255). This is a WORD-address (2-byte steps), allowing to access 256 words (512 bytes). Writing to this register resets the 1st/2nd access flipflop (for 2122h/213Bh) to 1st access.
func (ppu *PPU) cgadd(addr uint8) uint8 {
	ppu.cgramAddr = 2 * uint16(addr)
	return 0
}

// 2122 - CGDATA - Palette CGRAM Data Write (W)
func (ppu *PPU) cgdata(data uint8) uint8 {
	if ppu.cgramAddr%2 == 0 {
		// Write to the temporary variable
		ppu.cgramLsb = data
	} else {
		ppu.cgram[ppu.cgramAddr-1] = ppu.cgramLsb
		ppu.cgram[ppu.cgramAddr] = data
	}
	ppu.cgramAddr = (ppu.cgramAddr + 1) % 512
	return 0
}

// 213B - RDCGRAM - Palette CGRAM Data Read (R)
func (ppu *PPU) rdcgram(_ uint8) uint8 {
	res := ppu.cgram[ppu.cgramAddr]
	ppu.cgramAddr = (ppu.cgramAddr + 1) % 512
	return res
}
