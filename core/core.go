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

	//state
	state     string
	pauseChan chan struct{}
	stopChan  chan struct{}
	stepChan  chan int
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

	return &Emulator{
		CPU:       cpu,
		Memory:    mem,
		PPU:       ppu,
		state:     "paused",
		pauseChan: make(chan struct{}),
		stopChan:  make(chan struct{}),
		stepChan:  make(chan int),
	}
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

// ReadROM open the rom at filename and load it in memory
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

func (e *Emulator) loop() {
	n := 0
	for {
		switch e.state {
		case "stopped":
			return
		case "paused":
			n = e.statePaused()
		case "started":
			e.stateStarted(n)

		}
	}
}

func (e *Emulator) stateStarted(n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			select {
			case <-e.pauseChan:
				e.state = "paused"
				return
			case <-e.stopChan:
				e.state = "stopped"
				return
			default:
				e.CPU.execOpcode()
			}
		}

		// go back to paused state if execution is finished
		e.state = "paused"
		return
	} else {
		for {
			select {
			case <-e.pauseChan:
				e.state = "paused"
				return
			case <-e.stopChan:
				e.state = "stopped"
				return
			default:
				e.CPU.execOpcode()
			}
		}
	}
}

func (e *Emulator) statePaused() int {
	select {
	case <-e.stopChan:
		e.state = "stopped"

	case <-e.pauseChan:
		e.state = "started"

	case n := <-e.stepChan:
		e.state = "started"
		return n
	}
	return 0
}

// Start the main emulator loop
func (e *Emulator) Start() {
	e.state = "started"
	go e.loop()
}

func (e *Emulator) TogglePause() {
	e.pauseChan <- struct{}{}
}

func (e *Emulator) Stop() {
	close(e.stopChan)
}

func (e *Emulator) Step(n int) {
	e.stepChan <- n
}
