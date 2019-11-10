package core

import (
	"testing"

	"github.com/snes-emu/gose/rom"

	"github.com/stretchr/testify/assert"
)

func TestSramRegionLorom(t *testing.T) {
	mem := newMemory()
	mem.romType = rom.LoROM
	mem.initMmap()

	for bank := 0x70; bank < 0x7E; bank++ {
		assert.Equal(t, sramRegion, mem.mmap[bank<<4], "bank: 0x%X", bank)
	}
	for bank := 0xF0; bank > 0x100; bank++ {
		assert.Equal(t, sramRegion, mem.mmap[bank<<4])
	}
}

func TestSramRegionHirom(t *testing.T) {
	mem := newMemory()
	mem.romType = rom.HiROM
	mem.initMmap()

	for bank := 0x20; bank < 0x3F; bank++ {
		assert.Equal(t, sramRegion, mem.mmap[bank<<4|0x6], "bank: 0x%X", bank)
	}
	for bank := 0xA0; bank < 0xBF; bank++ {
		assert.Equal(t, sramRegion, mem.mmap[bank<<4|0x6], "bank: 0x%X", bank)
	}
}

func TestSramGetSet(t *testing.T) {
	mem := newMemory()
	mem.romType = rom.LoROM
	mem.sram = make([]uint8, sramSize)
	mem.initMmap()

	value := uint8(0xDE)
	bank := uint8(0x70)
	offset := uint16(0x2000)

	mem.SetByteBank(value, bank, offset)

	assert.Equal(t, value, mem.GetByteBank(bank, offset))
}

func TestLowWramGetSet(t *testing.T) {
	mem := newMemory()
	mem.romType = rom.LoROM
	mem.initMmap()

	value := uint8(0xDE)
	bank := uint8(0x10)
	offset := uint16(0x1000)

	mem.SetByteBank(value, bank, offset)

	assert.Equal(t, value, mem.GetByteBank(bank, offset))
}

func TestWramGetSet(t *testing.T) {
	mem := newMemory()
	mem.romType = rom.LoROM
	mem.initMmap()

	value := uint8(0xDE)
	bank := uint8(0x7E)
	offset := uint16(0x1000)

	mem.SetByteBank(value, bank, offset)

	assert.Equal(t, value, mem.GetByteBank(bank, offset))
}

func TestROMGetSet(t *testing.T) {
	mem := newMemory()
	mem.romType = rom.LoROM
	mem.initMmap()

	value := uint8(0xDE)
	bank := uint8(0x20)
	offset := uint16(0x9000)

	mem.SetByteBank(value, bank, offset)

	//cannot write to rom region
	assert.Equal(t, uint8(0x00), mem.GetByteBank(bank, offset))
}

func TestRomMappingLorom(t *testing.T) {
	//create a 3M rom
	bytes := make([]byte, 3*1024*1024)
	name := []byte("Test ROM             ")
	for i, ch := range name {
		bytes[0x7FC0+i] = ch
	}

	value := uint8(0xFE)
	offset := uint16(0x4321)
	bytes[offset] = value

	rom, err := rom.ParseROM(bytes)
	assert.Nil(t, err)

	mem := newMemory()
	mem.LoadROM(*rom)

	assert.Equal(t, value, mem.GetByteBank(0x00, 0x8000+offset))
}

func TestRomMappingHirom(t *testing.T) {
	//create a 3M rom
	bytes := make([]byte, 3*1024*1024)
	name := []byte("Test ROM             ")
	for i, ch := range name {
		bytes[0xFFC0+i] = ch
	}

	value := uint8(0xFE)
	offset := uint16(0x4321)
	bytes[offset] = value

	rom, err := rom.ParseROM(bytes)
	assert.Nil(t, err)

	mem := newMemory()
	mem.LoadROM(*rom)

	assert.Equal(t, value, mem.GetByteBank(0x40, offset))
}
