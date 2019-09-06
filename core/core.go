package core

import (
	"archive/zip"
	"io/ioutil"
	"os"

	"github.com/snes-emu/gose/apu"
	"github.com/snes-emu/gose/rom"
	"go.uber.org/zap"
)

// Emulator gathers the components required for emulation (PPU, CPU, Memory)
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

// New creates a new Emulator (creating the underlying components)
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
	lg := zap.L()
	buf, err := readFile(filename)
	if err != nil {
		lg.Fatal("error when reading rom file", zap.Error(err))
	}

	rom, err := rom.ParseROM(buf)
	if err != nil {
		lg.Fatal("an error occured while parsing the ROM", zap.Error(err))
	}
	lg.Info("sucess parsing rom", zap.String("name", rom.Title))

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
	}
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

// TogglePause toggles a pause in execution
func (e *Emulator) TogglePause() {
	e.pauseChan <- struct{}{}
}

// Stop stops the emulation
func (e *Emulator) Stop() {
	close(e.stopChan)
}

// Step continues the execution for the given number of steps (if given 0 it will loop until a pause is triggered or the emulator is stopped)
func (e *Emulator) Step(n int) {
	e.stepChan <- n
}
