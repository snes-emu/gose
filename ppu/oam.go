package ppu

// 2102 - OAMADDL
func (ppu *PPU) oamaddl(data uint8) uint8 {
	ppu.oamAddr = (ppu.oamLastWrittenAddr & 0x0100) | uint16(data)
	ppu.oamLastWrittenAddr = ppu.oamAddr
	return 0
}

// 2103 - OAMADDH
func (ppu *PPU) oamaddh(data uint8) uint8 {
	ppu.oamAddr = (uint16(data) << 8) | (ppu.oamLastWrittenAddr & 0x00ff)
	ppu.oamLastWrittenAddr = ppu.oamAddr
	return 0
}

// 2104 - OAMDATA - OAM Data Write (W)
func (ppu *PPU) oamdata(data uint8) uint8 {
	if ppu.oamFlip == 0 {
		// Write to the temporary variable
		ppu.oamLsb = data
	} else {
		// Remove the Obj Priority activation bit and keep only the b aaaaaaaa part
		addr := 2 * (ppu.oamAddr & 0x01ff)
		ppu.oam[addr] = ppu.oamLsb
		ppu.oam[addr+1] = data

		// Increment the address
		ppu.oamAddr++
	}
	// Change the oam flip
	ppu.oamFlip ^= 1
	return 0
}

// 2138 - RDOAM - OAM Data Read (R)
func (ppu *PPU) rdoam(_ uint8) uint8 {
	res := ppu.oam[2*(ppu.oamLastWrittenAddr&0x01ff)+ppu.oamFlip]
	// Increment address only if Flip value is 1 (end of a word reached)
	ppu.oamAddr += ppu.oamFlip
	ppu.oamFlip ^= 1
	return res
}
