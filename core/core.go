package core

import (
	"archive/zip"
	"io/ioutil"
	"os"

	"github.com/snes-emu/gose/apu"
	"github.com/snes-emu/gose/io"
	"github.com/snes-emu/gose/log"
	"github.com/snes-emu/gose/render"
	"github.com/snes-emu/gose/rom"
	"go.uber.org/zap"
)

// Emulator gathers the components required for emulation (PPU, CPU, Memory)
type Emulator struct {
	CPU    *CPU
	Memory *Memory
	PPU    *PPU

	//state
	state       *state
	pauseChan   chan struct{}
	resumeChan  chan struct{}
	stopChan    chan struct{}
	stepChan    chan int
	notifyPause chan struct{}

	//debugging
	registerBreakpoints map[string]struct{}
	breakpoint          uint32
	BreakpointCh        chan BreakpointData
	debug               bool
}

// New creates a new Emulator (creating the underlying components)
func New(renderer render.Renderer, debug bool) *Emulator {
	state := NewState()
	state.Pause()

	e := &Emulator{
		state:               state,
		pauseChan:           make(chan struct{}, 1),
		resumeChan:          make(chan struct{}),
		stopChan:            make(chan struct{}),
		stepChan:            make(chan int),
		notifyPause:         make(chan struct{}),
		debug:               debug,
		registerBreakpoints: map[string]struct{}{},
		BreakpointCh:        make(chan BreakpointData),
	}

	var rf *io.RegisterFactory
	if debug {
		rf = io.NewRegisterFactoryWithHook(e.handleRegisterBreakpoint)
	} else {
		rf = io.NewRegisterFactory()
	}

	apu := apu.New(rf)
	ppu := newPPU(renderer, rf)
	mem := newMemory()
	cpu := newCPU(mem, rf)

	cpu.ppu = ppu
	ppu.cpu = cpu

	mem.cpu = cpu
	mem.ppu = ppu
	mem.apu = apu
	mem.initIo(rf)

	e.Memory = mem
	e.CPU = cpu
	e.PPU = ppu

	return e
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

	e.PPU.renderer.SetRomTitle(rom.Title)
	e.Memory.LoadROM(*rom)
	e.CPU.Init()
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

func (e *Emulator) step() bool {
	select {
	case <-e.pauseChan:
		e.state.Pause()
		return false
	case <-e.stopChan:
		e.state.Stop()
		return false
	case <-e.resumeChan:
		e.CPU.execOpcode()
		if e.atBreakpoint() {
			e.state.Pause()
			return false
		}
	default:
		e.CPU.execOpcode()
		if e.atBreakpoint() {
			e.state.Pause()
			return false
		}
	}

	return true
}

func (e *Emulator) stateStarted(n int) {
	if n > 0 {
		for i := 0; i < n; i++ {
			if !e.step() {
				return
			}
		}

		// go back to paused state if execution is finished
		e.state.Pause()
		return
	}
	for e.step() {
	}
}

func (e *Emulator) statePaused() int {
	log.Info("entering paused state")

	//notify we entered the pause state in case someone is listening
	select {
	case e.notifyPause <- struct{}{}:
	default:
	}

	select {
	case <-e.stopChan:
		e.state.Stop()

	case <-e.resumeChan:
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
	initState := started
	if e.debug {
		initState = paused
	}
	e.startState(initState)
}

// Pause pauses the execution
func (e *Emulator) Pause() {
	e.pauseChan <- struct{}{}
	log.Info("execution paused")
}

// Resume resumes the execution
func (e *Emulator) Resume() {
	e.resumeChan <- struct{}{}
	log.Info("execution resumed")
}

// IsPaused checks if the execution is paused
func (e *Emulator) IsPaused() bool {
	return e.state.Status() == paused
}

// Stop stops the emulation
func (e *Emulator) Stop() {
	close(e.stopChan)
}

// Step continues the execution for the given number of steps (if given 0 it will loop until a pause is triggered or the emulator is stopped)
func (e *Emulator) Step(n int) {
	e.stepChan <- n
}

func (e *Emulator) WaitPaused() {
	if e.state.Status() == paused {
		return
	}

	<-e.notifyPause
}
