package ppu

type cgram struct {
	bytes [0x200]byte // cgram represents the color graphics ram and stores the color palette with 256 color entries
	addr  uint16      // store the cgram address over 512 byte (not the Word addr !)
	lsb   uint8       // temporary variable for the cgdata register
}

// 2121 - Color index (0..255). This is a WORD-address (2-byte steps), allowing to access 256 words (512 bytes). Writing to this register resets the 1st/2nd access flipflop (for 2122h/213Bh) to 1st access.
func (ppu *PPU) cgadd(addr uint8) {
	ppu.cgram.addr = 2 * uint16(addr)
}

// 2122 - CGDATA - Palette CGRAM Data Write (W)
func (ppu *PPU) cgdata(data uint8) {
	if ppu.cgram.addr%2 == 0 {
		// Write to the temporary variable
		ppu.cgram.lsb = data
	} else {
		ppu.cgram.bytes[ppu.cgram.addr-1] = ppu.cgram.lsb
		ppu.cgram.bytes[ppu.cgram.addr] = data
	}
	ppu.cgram.addr = (ppu.cgram.addr + 1) % 512
}

// 213B - RDCGRAM - Palette CGRAM Data Read (R)
func (ppu *PPU) rdcgram() uint8 {
	res := ppu.cgram.bytes[ppu.cgram.addr]
	ppu.cgram.addr = (ppu.cgram.addr + 1) % 512
	return res
}
