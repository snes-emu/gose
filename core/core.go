package core

import (
	"archive/zip"
	"github.com/snes-emu/gose/log"
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
	state     *state
	pauseChan chan struct{}
	stopChan  chan struct{}
	stepChan  chan int

	//debugging
	breakpoint uint32
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

	state := NewState()
	state.Pause()

	return &Emulator{
		CPU:       cpu,
		Memory:    mem,
		PPU:       ppu,
		state:     state,
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
		log.Fatal("error when reading rom file", zap.Error(err))
	}

	rom, err := rom.ParseROM(buf)
	if err != nil {
		log.Fatal("an error occurred while parsing the ROM", zap.Error(err))
	}
	log.Info("success parsing rom", zap.String("name", rom.Title))

	e.Memory.LoadROM(*rom)
	e.CPU.Init()
}

func (e *Emulator) SetBreakpoint(addr uint32) {
	e.breakpoint = addr
}

func (e *Emulator) atBreakpoint() bool {
	return e.breakpoint != 0 && uint16(e.breakpoint&0xFFFF) == e.CPU.PC && uint8(e.breakpoint>>16) == e.CPU.K
}

func (e *Emulator) loop() {
	n := 0
	for {
		switch e.state.Status() {
		case stopped:
			return
		case paused:
			n = e.statePaused()
		case started:
			e.stateStarted(n)

		}
	}
}

func (e *Emulator) stateStarted(n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			select {
			case <-e.pauseChan:
				e.state.Pause()
				return
			case <-e.stopChan:
				e.state.Stop()
				return
			default:
				e.CPU.execOpcode()
				if e.atBreakpoint() {
					e.state.Pause()
					return
				}
			}
		}

		// go back to paused state if execution is finished
		e.state.Pause()
		return
	}
	for {
		select {
		case <-e.pauseChan:
			e.state.Pause()
			return
		case <-e.stopChan:
			e.state.Stop()
			return
		default:
			e.CPU.execOpcode()
			if e.atBreakpoint() {
				e.state.Pause()
				return
			}
		}
	}
}

func (e *Emulator) statePaused() int {
	select {
	case <-e.stopChan:
		e.state.Stop()

	case <-e.pauseChan:
		e.state.Start()

	case n := <-e.stepChan:
		e.state.Start()
		return n
	}
	return 0
}

func (e *Emulator) startState(status stateStatus) {
	e.state.Start()
	e.state.SetStatus(status)
	go e.loop()
}

// Start the main emulator loop
func (e *Emulator) Start() {
	e.startState(started)
}

// StartPaused the main emulator loop
func (e *Emulator) StartPaused() {
	e.startState(paused)
}

// TogglePause toggles a pause in execution
func (e *Emulator) TogglePause() {
	e.pauseChan <- struct{}{}
}

// TogglePause toggles a pause in execution
func (e *Emulator) IsPaused() bool {
	return e.state.Status() == paused
}

// Stop stops the emulation
func (e *Emulator) Stop() {
	// Dump the palette somewhere
	{
		f, err := os.Create("/tmp/palette.snes")
		if err != nil {
			log.Error("failed to dump color palette", zap.Error(err))
		} else {
			defer f.Close()
			f.Write(e.PPU.cgram.bytes[:])
		}
	}
	{
		f, err := os.Create("/tmp/vram.snes")
		if err != nil {
			log.Error("failed to dump color palette", zap.Error(err))
		} else {
			defer f.Close()
			f.Write(e.PPU.vram.bytes[:])
		}
	}
	close(e.stopChan)
}

// Step continues the execution for the given number of steps (if given 0 it will loop until a pause is triggered or the emulator is stopped)
func (e *Emulator) Step(n int) {
	e.stepChan <- n
}
