package core

import (
	"github.com/snes-emu/gose/log"
	"go.uber.org/zap"
	"strings"
)

type BreakpointData struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data uint8  `json:"data"`
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
	if e.atRegisterBreakpoint(name) {
		log.Info(
			"breakpoint reached, pausing execution...",
			zap.String("register", name),
			zap.String("type", typ),
			zap.Uint8("data", data),
		)
		// TODO: for now we just queue up DATA but issues can arise during DMAs:
		// it can occur that we are in a DMA and we already paused the emulator
		// however we don't handle interrupting a DMA from the debugger yet, so we just queue up the breakpoint data
		if !e.IsPaused() {
			// Set the emulator state to pause since we reached a breakpoint
			e.state.Pause()
			e.BreakpointCh <- BreakpointData{Name: name, Type: typ, Data: data}
		}
	}
}
