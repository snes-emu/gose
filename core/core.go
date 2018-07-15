package core

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/snes-emu/gose/apu"
	"github.com/snes-emu/gose/rom"
)

type Emulator struct {
	CPU    *CPU
	Memory *Memory
	PPU    *PPU
}

func New() *Emulator {
	apu := apu.New()
	ppu := newPPU()
	mem := newMemory()
	cpu := newCPU(mem)

	cpu.ppu = ppu
	ppu.cpu = cpu

	mem.cpu = cpu
	mem.ppu = ppu
	mem.apu = apu
	mem.initIo()

	return &Emulator{cpu, mem, ppu}
}

func (e *Emulator) ReadROM(filename string) {
	ROM, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	rom, err := rom.ParseROM(ROM)
	fmt.Println(rom.Title)

	if err != nil {
		log.Fatalf("There were a problem while importing the ROM: %v", err)
	}

	e.Memory.LoadROM(*rom)
	e.CPU.Init()
}
