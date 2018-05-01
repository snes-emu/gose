package core

import (
	"github.com/snes-emu/gose/io"
	"github.com/snes-emu/gose/ppu"
	"github.com/snes-emu/gose/rom"
)

const bankNumber = 0x100
const offsetMask = 0xFFFF
const sramSize = 0x8000
const wramSize = 0x20000
const ioSize = 0x8000

const (
	lowWramRegion = iota
	ioRegisterRegion
	romRegion
	wramRegion
	sramRegion
)

// Memory struct containing SNES working RAM, cartridge static RAM, special hardware registers and default memory buffer for ROM
type Memory struct {
	main    [bankNumber][]uint8
	sram    [sramSize]uint8
	wram    [wramSize]uint8
	io      [ioSize]io.Register
	romType uint
	ppu     *ppu.PPU
	cpu     *CPU
}

// New creates a Memory struct and initialize it
func newMemory() *Memory {
	memory := &Memory{}
	for bank := 0; bank < bankNumber; bank++ {
		memory.main[bank] = make([]byte, 0x10000)
	}
	return memory
}

// LoadROM takes a memory buffer and load it into memory depending ROM type
func (memory *Memory) LoadROM(r rom.ROM) {

	// only LoROM for now
	switch r.Type {
	case rom.LoROM:
		memory.romType = rom.LoROM
		romSize := len(r.Data)
		for bank := 0x00; bank < 0x80; bank++ {
			for offset := 0x8000; offset < 0x10000; offset++ {
				if pos := offset + (bank-1)*0x8000; pos < romSize {
					memory.main[bank][offset] = r.Data[pos]
				}
			}
		}

		for bank := 0x80; bank < 0x100; bank++ {
			memory.main[bank] = memory.main[bank-0x80]
		}
	}
}

func (memory *Memory) Init() {
	for i := 0; i < ioSize; i++ {
		memory.io[i] = io.UnusedRegister
	}
	for i := 0; i < 0x40; i++ {
		memory.io[0x2100+i] = memory.ppu.Registers[i]
	}
}

//GetByte gets a byte by its complete address
func (memory Memory) GetByte(index uint32) uint8 {
	K := index >> 16
	offset := index & offsetMask
	return memory.GetByteBank(uint8(K), uint16(offset))
}

//GetByteBank gets a byte by memory bank and offset
func (memory Memory) GetByteBank(K uint8, offset uint16) uint8 {
	switch memory.romType {
	case rom.LoROM:
		if K < 0x40 || (0x7F < K && K < 0xC0) {
			if offset < 0x2000 {
				return memory.wram[offset]
			} else if 0x2133 < offset && offset < 0x2140 {
				return memory.ppu.Registers[offset-0x2100](0)
			}
		} else if offset < 0x8000 && ((0x6F < K && K < 0x7E) || (0xEF < K && K < 0xFE)) {
			return memory.sram[offset]
		} else if K > 0x7D && K < 0x80 {
			return memory.wram[(uint32(K)-0x7E)<<16+uint32(offset)]
		} else if 0xFD < K && offset < 0x8000 {
			return memory.sram[offset]
		}
		return memory.main[K][offset]
	default:
		return 0x00
	}
}

//SetByte sets a byte by its complete address
func (memory *Memory) SetByte(value uint8, index uint32) {
	K := index >> 16
	offset := index & offsetMask
	memory.SetByteBank(value, uint8(K), uint16(offset))
}

//SetByteBank sets a byte by memory bank and offset
func (memory *Memory) SetByteBank(value uint8, K uint8, offset uint16) {
	switch memory.romType {
	case rom.LoROM:
		if K < 0x40 || (0x7F < K && K < 0xC0) {
			if offset < 0x2000 {
				memory.wram[offset] = value
			} else if 0x20FF < offset && offset < 0x2134 {
				memory.ppu.Registers[offset-0x2100](value)
			}
		} else if offset < 0x8000 && ((0x6F < K && K < 0x7E) || (0xEF < K && K < 0xFE)) {
			memory.sram[offset] = value
		} else if K > 0x7D && K < 0x80 {
			memory.wram[(uint32(K)-0x7E)<<16+uint32(offset)] = value
		} else if 0xFD < K && offset < 0x8000 {
			memory.sram[offset] = value
		}
	}
}
