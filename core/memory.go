package core

import (
	"github.com/snes-emu/gose/apu"
	"github.com/snes-emu/gose/io"
	"github.com/snes-emu/gose/rom"
)

const regionNumber = 0x1000
const bankNumber = 0x100
const offsetMask = 0xFFFF
const sramSize = 0x80000
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
	mmap    [regionNumber]uint
	main    [bankNumber][]uint8
	sram    [sramSize]uint8
	wram    [wramSize]uint8
	io      [ioSize]*io.Register
	romType uint
	apu     *apu.APU
	ppu     *PPU
	cpu     *CPU
}

// New creates a Memory struct and initialize it
func newMemory() *Memory {
	memory := &Memory{}
	for bank := 0; bank < bankNumber; bank++ {
		memory.main[bank] = make([]byte, 0x10000)
	}

	//set rom region as default
	for region := 0; region < regionNumber; region++ {
		memory.mmap[region] = romRegion
	}
	return memory
}

// LoadROM takes a memory buffer and load it into memory depending ROM type
func (memory *Memory) LoadROM(r rom.ROM) {
	romSize := len(r.Data)

	// set the ROM section according to: https://en.wikibooks.org/wiki/Super_NES_Programming/SNES_memory_map
	switch r.Type {
	case rom.LoROM:
		memory.romType = rom.LoROM
		for bank := 0x00; bank < 0x80; bank++ {
			for offset := 0x8000; offset < 0x10000; offset++ {
				if pos := offset + (bank-1)*0x8000; pos < romSize {
					memory.main[bank][offset] = r.Data[pos]
				}
			}
		}
	case rom.HiROM:
		memory.romType = rom.HiROM
		for bank := 0x00; bank < 0x40; bank++ {
			for offset := 0; offset < 0x10000; offset++ {
				if pos := offset + bank*0x10000; pos < romSize {
					memory.main[bank][offset] = r.Data[pos]
					memory.main[bank+0x40][offset] = r.Data[pos]
				}
			}
		}
	}

	//upper banks are mirrors of the lower ones
	for bank := 0x80; bank < 0x100; bank++ {
		memory.main[bank] = memory.main[bank-0x80]
	}
	memory.initMmap()
}

func (memory *Memory) initIo(rf *io.RegisterFactory) {
	for i := 0; i < ioSize; i++ {
		memory.io[i] = rf.NewRegister(nil, nil)
	}
	for i := 0; i < 0x40; i++ {
		memory.io[0x2100+i] = memory.ppu.Registers[i]
		memory.io[0x2140+i] = memory.apu.Registers[i%4]
	}
	for i := 0; i < 0x380; i++ {
		memory.io[0x4000+i] = memory.cpu.ioRegisters[i]
	}
}

func (memory *Memory) initMmap() {
	//map system reserved banks
	for bankIndex := 0x0; bankIndex < 0x40; bankIndex++ {
		for offset := 0x0; offset < 0x2; offset++ {
			memory.mmap[bankIndex<<4|offset] = lowWramRegion
			memory.mmap[(bankIndex+0x80)<<4|offset] = lowWramRegion
		}
		for offset := 0x2; offset < 0x8; offset++ {
			memory.mmap[bankIndex<<4|offset] = ioRegisterRegion
			memory.mmap[(bankIndex+0x80)<<4|offset] = ioRegisterRegion
		}
	}

	//map sram
	switch memory.romType {
	case rom.LoROM:
		for bankIndex := 0x70; bankIndex < 0x80; bankIndex++ {
			for offset := 0; offset < 0x8; offset++ {
				memory.mmap[bankIndex<<4|offset] = sramRegion
				memory.mmap[(bankIndex+0x80)<<4|offset] = sramRegion
			}
		}
	case rom.HiROM:
		//in HiROM sram is mapped here: overwrite the unused ioRegister regions
		for bankIndex := 0x20; bankIndex < 0x40; bankIndex++ {
			for offset := 0x6; offset < 0x8; offset++ {
				memory.mmap[bankIndex<<4|offset] = sramRegion
				memory.mmap[(bankIndex+0x80)<<4|offset] = sramRegion
			}
		}
	}

	//map work ram
	//note: in LoROM this will overwrite the last sram regions in the lower banks
	for bankIndex := 0x7E; bankIndex < 0x80; bankIndex++ {
		for offset := 0; offset < 0x10; offset++ {
			memory.mmap[bankIndex<<4|offset] = wramRegion
		}
	}
}

//GetByte gets a byte by its complete address
func (memory *Memory) GetByte(index uint32) uint8 {
	K := index >> 16
	offset := index & offsetMask
	return memory.GetByteBank(uint8(K), uint16(offset))
}

//GetByteBank gets a byte by memory bank and offset
func (memory *Memory) GetByteBank(K uint8, offset uint16) uint8 {
	switch memory.mmap[uint16(K)<<4|offset>>12] {
	case lowWramRegion:
		return memory.wram[offset]
	case ioRegisterRegion:
		return memory.io[offset].Read()
	case romRegion:
		return memory.main[K][offset]
	case wramRegion:
		return memory.wram[(uint32(K)-0x7E)<<16+uint32(offset)]
	case sramRegion:
		return memory.sram[offset]
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
	switch memory.mmap[uint16(K)<<4|offset>>12] {
	case lowWramRegion:
		memory.wram[offset] = value
	case ioRegisterRegion:
		memory.io[offset].Write(value)
	case wramRegion:
		memory.wram[(uint32(K)-0x7E)<<16+uint32(offset)] = value
	case sramRegion:
		memory.sram[offset] = value
	}
}
