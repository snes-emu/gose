package core

// cgram represents the color graphics ram and stores the color palette with 256 color entries
// Colors are stored in 2-byte words like:
// _bbbbbgg gggrrrrr
// The last bit is not used
type cgram struct {
	bytes [0x200]byte // byte the raw 512 bytes (for 256 color entries)
	addr  uint16      // store the cgram address over 512 byte (not the Word addr !)
	// lsb temporary variable for the cgdata register, it's used when we want to write a new color:
	// - we first set the addr where we want to write (for instance the color number 67, we will have 61 * 2 = 122 stored in addr
	// - we call cgdata to write once (this will store the provided data in the lsb)
	// - we call cgdata one more time and then we will write (data << 8 | lsb) to the registry
	lsb uint8
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
		ppu.cgram.write(ppu.cgram.addr, ppu.cgram.lsb, data)
	}
	ppu.cgram.incrAddr()
}

// 213B - RDCGRAM - Palette CGRAM Data Read (R)
func (ppu *PPU) rdcgram() uint8 {
	res := ppu.cgram.read(ppu.cgram.addr)
	ppu.cgram.incrAddr()
	return res
}

func (cg *cgram) read(addr uint16) uint8 {
	return cg.bytes[addr]
}

func (cg *cgram) write(addr uint16, low uint8, high uint8) {
	// This could be (high & 0x7f) but since last bit is never used in color palette it's not an issue
	cg.bytes[addr] = high
	cg.bytes[addr-1] = low
}

func (cg *cgram) incrAddr() {
	cg.addr = (cg.addr + 1) % 512
}
