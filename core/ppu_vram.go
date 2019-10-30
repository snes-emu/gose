package core

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
	prefetch        uint16        // a cache value used when using vmaddl/vmaddh registers that stores the vram value at the new address
}

// 2115 - VMAIN - VRAM Address Increment Mode (W)
// Address increment mode:
// 0 => increment after writing into 2118/ reading from 2139
// 1 => increment after writing into 2119/ reading from 213a
func (ppu *PPU) vmain(data uint8) {
	ppu.vram.incrementMode = data&0x80 != 0

	ppu.vram.incrementAmount = incrAmount(data)
	ppu.vram.addrMapping = data & 0xc >> 2
}

// 2116 - VMADDL - VRAM Address (lower 8bit) (W)
func (ppu *PPU) vmaddl(data uint8) {
	ppu.vram.addr = (ppu.vram.addr & 0xff00) | uint16(data)
	ppu.vram.prefetchWord()
}

// 2117 - VMADDH - VRAM Address (upper 8bit) (W)
func (ppu *PPU) vmaddh(data uint8) {
	ppu.vram.addr = bit.JoinUint16(0x00, data) | (ppu.vram.addr & 0x0ff)
	ppu.vram.prefetchWord()
}

// 2118 - VMDATAL - VRAM Data Write (lower 8bit) (W)
func (ppu *PPU) vmdatal(data uint8) {
	ppu.vram.bytes[2*ppu.vram.getAddr()] = data

	if !ppu.vram.incrementMode {
		// No prefetching done there
		ppu.vram.addr += ppu.vram.incrementAmount
	}
}

// 2119 - VMDATAH - VRAM Data Write (upper 8bit) (W)
func (ppu *PPU) vmdatah(data uint8) {
	ppu.vram.bytes[2*ppu.vram.getAddr()+1] = data

	if ppu.vram.incrementMode {
		// No prefetching done there
		ppu.vram.addr += ppu.vram.incrementAmount
	}
}

// 2139 - RDVRAML - VRAM Data Read (lower 8bit) (R)
// This reads from the prefetch value and not the VRAM data directly !
// From: https://problemkaputt.de/fullsnes.htm#snesppucolorpalettememorycgramanddirectcolors
// Reading from these registers returns the LSB or MSB of an internal 16bit prefetch register. Depending on the Increment Mode the address does (or doesn't) get automatically incremented after the read.
// The prefetch register is filled with data from the currently addressed VRAM word (with optional Address Translation applied) upon two situations:
//
// Prefetch occurs AFTER changing the VRAM address (by writing 2116h/17h).
// Prefetch occurs BEFORE incrementing the VRAM address (by reading 2139h/3Ah).
//
// The "Prefetch BEFORE Increment" effect is some kind of a hardware glitch (Prefetch AFTER Increment would be more useful). Increment/Prefetch in detail:
//
// 1st  Send a byte from OLD prefetch value to the CPU        ;-this always
// 2nd  Load NEW value from OLD address into prefetch register;\these only if
// 3rd  Increment address so it becomes the NEW address       ;/increment occurs
//
// Increments caused by writes to 2118h/19h don't do any prefetching (the prefetch register is left totally unchanged by writes).
// In practice: After changing the VRAM address (via 2116h/17h), the first byte/word will be received twice, further values are received from properly increasing addresses (as a workaround: issue a dummy-read that ignores the 1st or 2nd value).
func (ppu *PPU) rdvraml() uint8 {
	res := bit.LowByte(ppu.vram.prefetch)

	if !ppu.vram.incrementMode {
		ppu.vram.prefetchAndIncrAddr()
	}

	return res
}

// 213A - RDVRAMH - VRAM Data Read (upper 8bit) (R)
// This reads from the prefetch value and not the VRAM data directly !
func (ppu *PPU) rdvramh() uint8 {
	res := bit.HighByte(ppu.vram.prefetch)

	if ppu.vram.incrementMode {
		ppu.vram.prefetchAndIncrAddr()
	}

	return res
}

func (vr *vram) prefetchWord() {
	newAddr := vr.getAddr()
	vr.prefetch = bit.JoinUint16(vr.bytes[2*newAddr+1], vr.bytes[2*newAddr])
}

func (vr *vram) prefetchAndIncrAddr() {
	vr.prefetchWord()
	vr.addr += vr.incrementAmount
}

// getAddr returns the vram addr performing the address translation
func (vr *vram) getAddr() uint16 {
	switch vr.addrMapping {
	case 0x0:
		// No remapping
		return vr.addr
	case 0x1:
		// Remap addressing aaaaaaaaBBBccccc => aaaaaaaacccccBBB
		return (vr.addr & 0xff00) | ((vr.addr & 0xe0) >> 5) | ((vr.addr & 0x1f) << 3)
	case 0x2:
		// Remap addressing aaaaaaaBBBcccccc => aaaaaaaccccccBBB
		return (vr.addr & 0xfe00) | ((vr.addr & 0x1c0) >> 6) | ((vr.addr & 0x3F) << 3)
	case 0x3:
		// Remap addressing aaaaaaBBBccccccc => aaaaaacccccccBBB
		return (vr.addr & 0xfc00) | ((vr.addr & 0xb80) >> 7) | ((vr.addr & 0x7F) << 3)

	default:
		panic(fmt.Sprintf("Unknown vram Addr mapping mode: %v", vr.addrMapping))
	}
}

func incrAmount(nb uint8) uint16 {
	switch nb & 0x3 {
	case 0:
		return 1
	case 1:
		return 32
	case 2, 3:
		return 128
	}
	panic("invalid increment amount index")
}
