package ppu

// 2102 - OAMADDL
func (ppu *PPU) oamaddl(data uint8) uint8 {
	ppu.oamAddr = (ppu.oamLastWrittenAddr & 0x0200) | (uint16(data) << 1)
	ppu.oamLastWrittenAddr = ppu.oamAddr
	return 0
}

// 2103 - OAMADDH
func (ppu *PPU) oamaddh(data uint8) uint8 {
	ppu.oamPriorityBit = data&0x80 != 0
	ppu.oamAddr = (uint16(data) << 9) | (ppu.oamLastWrittenAddr & 0x01fe)
	ppu.oamLastWrittenAddr = ppu.oamAddr
	return 0
}

// 2104 - OAMDATA - OAM Data Write (W)
func (ppu *PPU) oamdata(data uint8) uint8 {
	if ppu.oamAddr%2 == 0 {
		// Write to the temporary variable
		ppu.oamLsb = data
	}
	if ppu.oamAddr > 0x1FF {
		ppu.oam[ppu.oamAddr] = data
	} else if ppu.oamAddr%2 == 1 {
		// Remove the Obj Priority activation bit and keep only the b aaaaaaaa part
		ppu.oam[ppu.oamAddr-1] = ppu.oamLsb
		ppu.oam[ppu.oamAddr] = data
	}
	// Increment the address
	ppu.oamAddr = (ppu.oamAddr + 1) % 544
	return 0
}

// 2138 - RDOAM - OAM Data Read (R)
func (ppu *PPU) rdoam(_ uint8) uint8 {
	res := ppu.oam[ppu.oamAddr]
	ppu.oamAddr = (ppu.oamAddr + 1) % 544
	return res
}