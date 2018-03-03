package ppu

// 2115 - VMAIN - VRAM Address Increment Mode (W)
func (ppu *PPU) vmain(data uint8) uint8 {
	ppu.vramIncrementMode = data&0x80 != 0

	incrementValues := map[uint8]uint8{
		0x00: 1, 0x01: 32, 0x10: 128, 0x11: 128,
	}

	ppu.vramIncrementAmount = incrementValues[data&0x3]
	ppu.vramAddrMapping = data & 0xc >> 2

	return 0
}
