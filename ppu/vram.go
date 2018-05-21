package ppu

import (
	"fmt"

	"github.com/snes-emu/gose/bit"
)

type vram struct {
	bytes           [0x10000]byte // vram represents the VideoRAM (64KB)
	incrementMode   bool          // vram Address increment mode
	incrementAmount uint16        // vram Address increment amount
	addrMapping     uint8         // vram Address remaping (4 mode available)
	addr            uint16        // The vram addr (a word address !)
	cache           uint16        // a cache value used when using vmaddl/vmaddh registers that stores the vram value at the new address
}

// getvram.addr returns the vram addr performing the address translation
func (ppu *PPU) getvramAddr() uint16 {
	switch ppu.vram.addrMapping {
	case 0x0:
		// No remapping
		return ppu.vram.addr
	case 0x1:
		// Remap addressing aaaaaaaaBBBccccc => aaaaaaaacccccBBB
		return (ppu.vram.addr & 0xff00) | ((ppu.vram.addr & 0xe0) >> 5) | ((ppu.vram.addr & 0x1f) << 3)
	case 0x2:
		// Remap addressing aaaaaaaBBBcccccc => aaaaaaaccccccBBB
		return (ppu.vram.addr & 0xfe00) | ((ppu.vram.addr & 0x1c0) >> 6) | ((ppu.vram.addr & 0x3F) << 3)
	case 0x3:
		// Remap addressing aaaaaaBBBccccccc => aaaaaacccccccBBB
		return (ppu.vram.addr & 0xfc00) | ((ppu.vram.addr & 0xb80) >> 7) | ((ppu.vram.addr & 0x7F) << 3)

	default:
		panic(fmt.Sprintf("Unknown vram Addr mapping mode: %v", ppu.vram.addrMapping))
	}
}

// 2115 - VMAIN - VRAM Address Increment Mode (W)
func (ppu *PPU) vmain(data uint8) {
	ppu.vram.incrementMode = data&0x80 != 0

	incrementValues := map[uint8]uint16{
		0x00: 1, 0x01: 32, 0x10: 128, 0x11: 128,
	}

	ppu.vram.incrementAmount = incrementValues[data&0x3]
	ppu.vram.addrMapping = data & 0xc >> 2
}

// 2116 - VMADDL - VRAM Address (lower 8bit) (W)
func (ppu *PPU) vmaddl(data uint8) {
	ppu.vram.addr = (ppu.vram.addr & 0xff00) | uint16(data)
	newAddr := ppu.getvramAddr()
	ppu.vram.cache = bit.JoinUint16(ppu.vram.bytes[2*newAddr+1], ppu.vram.bytes[2*newAddr])
}

// 2117 - VMADDH - VRAM Address (upper 8bit) (W)
func (ppu *PPU) vmaddh(data uint8) {
	ppu.vram.addr = bit.JoinUint16(0x00, data) | (ppu.vram.addr & 0x0ff)
	newAddr := ppu.getvramAddr()
	ppu.vram.cache = bit.JoinUint16(ppu.vram.bytes[2*newAddr+1], ppu.vram.bytes[2*newAddr])
}

// 2118 - VMDATAL - VRAM Data Write (lower 8bit) (W)
func (ppu *PPU) vmdatal(data uint8) {

	ppu.vram.bytes[2*ppu.getvramAddr()] = data

	if ppu.vram.incrementMode {
		ppu.vram.addr += ppu.vram.incrementAmount
	}
}

// 2119 - VMDATAH - VRAM Data Write (upper 8bit) (W)
func (ppu *PPU) vmdatah(data uint8) {
	ppu.vram.bytes[2*ppu.getvramAddr()+1] = data

	if !ppu.vram.incrementMode {
		ppu.vram.addr += ppu.vram.incrementAmount
	}
}

// 2139 - RDVRAML - VRAM Data Read (lower 8bit) (R)
func (ppu *PPU) rdvraml() uint8 {
	res := ppu.vram.bytes[2*ppu.getvramAddr()]

	if ppu.vram.incrementMode {
		ppu.vram.addr += ppu.vram.incrementAmount
		newAddr := ppu.getvramAddr()
		ppu.vram.cache = bit.JoinUint16(ppu.vram.bytes[2*newAddr+1], ppu.vram.bytes[2*newAddr])
	}

	return res
}

// 213A - RDVRAMH - VRAM Data Read (upper 8bit) (R)
func (ppu *PPU) rdvramh() uint8 {
	res := ppu.vram.bytes[2*ppu.getvramAddr()+1]

	if !ppu.vram.incrementMode {
		ppu.vram.addr += ppu.vram.incrementAmount
		newAddr := ppu.getvramAddr()
		ppu.vram.cache = bit.JoinUint16(ppu.vram.bytes[2*newAddr+1], ppu.vram.bytes[2*newAddr])
	}

	return res
}
