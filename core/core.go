package core

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"strings"

	"github.com/snes-emu/gose/apu"
	"github.com/snes-emu/gose/io"
	"github.com/snes-emu/gose/log"
	"github.com/snes-emu/gose/render"
	"github.com/snes-emu/gose/rom"
	"go.uber.org/zap"
)

type BreakpointData struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data uint8  `json:"data"`
}

// Emulator gathers the components required for emulation (PPU, CPU, Memory)
type Emulator struct {
	CPU    *CPU
	Memory *Memory
	PPU    *PPU

	//state
	state     *state
	pauseChan chan chan bool
	stopChan  chan struct{}
	stepChan  chan int

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
		pauseChan:           make(chan chan bool, 1),
		stopChan:            make(chan struct{}),
		stepChan:            make(chan int),
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

func (e *Emulator) SetBreakpoint(addr uint32) {
	e.breakpoint = addr
}

func (e *Emulator) atBreakpoint() bool {
	return e.breakpoint != 0 && uint16(e.breakpoint&0xFFFF) == e.CPU.PC && uint8(e.breakpoint>>16) == e.CPU.K
}

func (e *Emulator) SetRegisterBreakpoint(registers string) {
	newBreakpoints := map[string]struct{}{}
	for _, reg := range strings.Split(registers, ",") {
		newBreakpoints[strings.ToUpper(strings.TrimSpace(reg))] = struct{}{}
	}
	e.registerBreakpoints = newBreakpoints
}

func (e *Emulator) atRegisterBreakpoint(register string) bool {
	_, ok := e.registerBreakpoints[register]
	return ok
}

func (e *Emulator) handleRegisterBreakpoint(name string, typ string, data uint8) {
	if !e.IsPaused() && e.atRegisterBreakpoint(name) {
		log.Debug(
			"breakpoint reached, pausing execution...",
			zap.String("register", name),
			zap.String("type", typ),
			zap.Uint8("data", data),
		)
		e.TogglePause()
		e.BreakpointCh <- BreakpointData{Name: name, Type: typ, Data: data}
	}
}

func (e *Emulator) pause(ch chan bool) {
	log.Debug("emulator paused")
	e.state.Pause()
	ch <- true
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
			case ch := <-e.pauseChan:
				e.pause(ch)
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
		case ch := <-e.pauseChan:
			e.pause(ch)
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

	case ch := <-e.pauseChan:
		e.state.Start()
		ch <- false

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

// TogglePause toggles a pause in execution
func (e *Emulator) TogglePause() chan bool {
	listenCh := make(chan bool, 1)
	e.pauseChan <- listenCh
	return listenCh
}

// TogglePause toggles a pause in execution
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
