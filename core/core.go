package core

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

func readFile(src string) ([]byte, error) {
	r, err := zip.OpenReader(src)

	if err != nil {
		// Probably not a zip file try to open it as a .smc directly
		f, err := os.Open(src)
		defer f.Close()

		if err != nil {
			return nil, err
		}

		return ioutil.ReadAll(f)
	}

	defer r.Close()

	var rom zip.File

	for _, f := range r.File {
		// Use the biggest file
		if f.FileInfo().Size() > rom.FileInfo().Size() {
			rom = *f
		}
	}

	f, err := rom.Open()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

func (e *Emulator) ReadROM(filename string) {
	buf, err := readFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	rom, err := rom.ParseROM(buf)
	fmt.Println(rom.Title)

	if err != nil {
		log.Fatalf("There were a problem while importing the ROM: %v", err)
	}

	e.Memory.LoadROM(*rom)
	e.CPU.Init()
}
