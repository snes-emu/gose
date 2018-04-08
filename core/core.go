package core

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/snes-emu/gose/ppu"
	"github.com/snes-emu/gose/rom"
)

type Emulator struct {
	CPU    *CPU
	Memory *Memory
	PPU    *ppu.PPU
}

func New() *Emulator {
	ppu := ppu.New()
	mem := newMemory()
	cpu := newCPU(mem)

	mem.cpu = cpu

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
}
