package ppu

import "fmt"

// vramAddrTranslation Performs the address translation
func (ppu PPU) vramAddrTranslation(addr uint16) uint16 {
	switch ppu.vramAddrMapping {
	case 0x0:
		// No remapping
		return addr
	case 0x1:
		// Remap addressing aaaaaaaaBBBccccc => aaaaaaaacccccBBB
		return (addr & 0xff00) | ((addr & 0xe0) >> 5) | ((addr & 0x1f) << 3)
	case 0x2:
		// Remap addressing aaaaaaaBBBcccccc => aaaaaaaccccccBBB
		return (addr & 0xfe00) | ((addr & 0x1c0) >> 6) | ((addr & 0x3F) << 3)
	case 0x3:
		// Remap addressing aaaaaaBBBccccccc => aaaaaacccccccBBB
		return (addr & 0xfc00) | ((addr & 0xb80) >> 7) | ((addr & 0x7F) << 3)

	default:
		panic(fmt.Sprintf("Unknown vram Addr mapping mode: %v", ppu.vramAddrMapping))
	}
}

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
