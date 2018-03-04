package ppu

import "fmt"

// getVramAddr returns the vram addr performing the address translation
func (ppu PPU) getVramAddr() uint16 {
	switch ppu.vramAddrMapping {
	case 0x0:
		// No remapping
		return ppu.vramAddr
	case 0x1:
		// Remap addressing aaaaaaaaBBBccccc => aaaaaaaacccccBBB
		return (ppu.vramAddr & 0xff00) | ((ppu.vramAddr & 0xe0) >> 5) | ((ppu.vramAddr & 0x1f) << 3)
	case 0x2:
		// Remap addressing aaaaaaaBBBcccccc => aaaaaaaccccccBBB
		return (ppu.vramAddr & 0xfe00) | ((ppu.vramAddr & 0x1c0) >> 6) | ((ppu.vramAddr & 0x3F) << 3)
	case 0x3:
		// Remap addressing aaaaaaBBBccccccc => aaaaaacccccccBBB
		return (ppu.vramAddr & 0xfc00) | ((ppu.vramAddr & 0xb80) >> 7) | ((ppu.vramAddr & 0x7F) << 3)

	default:
		panic(fmt.Sprintf("Unknown vram Addr mapping mode: %v", ppu.vramAddrMapping))
	}
}

// 2115 - VMAIN - VRAM Address Increment Mode (W)
func (ppu *PPU) vmain(data uint8) uint8 {
	ppu.vramIncrementMode = data&0x80 != 0

	incrementValues := map[uint8]uint16{
		0x00: 1, 0x01: 32, 0x10: 128, 0x11: 128,
	}

	ppu.vramIncrementAmount = incrementValues[data&0x3]
	ppu.vramAddrMapping = data & 0xc >> 2

	return 0
}

// 2118 - VMDATAL - VRAM Data Write (lower 8bit) (W)
func (ppu *PPU) vmdatal(data uint8) uint8 {

	ppu.vram[2*ppu.getVramAddr()] = data

	if ppu.vramIncrementMode {
		ppu.vramAddr += ppu.vramIncrementAmount
	}

	return 0
}

// 2119 - VMDATAH - VRAM Data Write (upper 8bit) (W)
func (ppu *PPU) vmdatah(data uint8) uint8 {
	ppu.vram[2*ppu.getVramAddr()+1] = data

	if !ppu.vramIncrementMode {
		ppu.vramAddr += ppu.vramIncrementAmount
	}

	return 0
}
