package apu

import (
	"fmt"

	"github.com/snes-emu/gose/log"

	"github.com/snes-emu/gose/io"
)

const APUIONum = 4

const (
	stateInit = iota
	stateTransfer
	stateRunning
)

type APU struct {
	state   int
	isReset bool

	addr          uint16
	cmd           uint8
	transferIndex uint8

	IO0       uint8
	IO1       uint8
	IO2       uint8
	IO3       uint8
	Registers [APUIONum]*io.Register
}

func New(rf *io.RegisterFactory) *APU {
	apu := &APU{}

	apu.Registers[0] = rf.NewRegister(apu.CPUIO0R, apu.CPUIO0W, "APUIO0")
	apu.Registers[1] = rf.NewRegister(apu.CPUIO1R, apu.CPUIO1W, "APUIO1")
	apu.Registers[2] = rf.NewRegister(apu.CPUIO2R, apu.CPUIO2W, "APUIO2")
	apu.Registers[3] = rf.NewRegister(apu.CPUIO3R, apu.CPUIO3W, "APUIO3")

	apu.reset()

	return apu
}

func (apu *APU) reset() {
	apu.isReset = true
	apu.transferIndex = 0x00
	apu.state = stateInit
	apu.IO0 = 0xAA
	apu.IO1 = 0xBB
}

func (apu *APU) stateInit(data uint8) {
	if apu.isReset {
		if data != 0xCC {
			log.Debug(fmt.Sprintf("stateInit: %X received while is reset\n", data))
			return
		}
		apu.isReset = false
	}
	apu.addr = uint16(apu.IO3)<<8 | uint16(apu.IO2)
	apu.cmd = apu.IO1
	if apu.cmd != 0 {
		apu.state = stateTransfer
	} else {
		apu.state = stateRunning
	}
	apu.IO0 = data

	log.Debug(fmt.Sprintf("stateInit: addr: %X, cmd: %X, kick %X\n", apu.addr, apu.cmd, apu.IO0))
}

func (apu *APU) stateTransfer(data uint8) {
	if apu.transferIndex == data {
		apu.transferIndex++
	} else if apu.transferIndex < data {
		// end of transfer
		apu.transferIndex = 0x00
		apu.state = stateInit
		apu.stateInit(data)
	}
	apu.IO0 = data

	log.Debug(fmt.Sprintf("stateTransfer: index: %X, kick: %X\n", apu.transferIndex, apu.IO0))
}

func (apu *APU) stateRunning(data uint8) {
	apu.reset()
	log.Debug(fmt.Sprintf("stateRunning: kick: %X\n", apu.IO0))
}

func (apu *APU) CPUIO0R() uint8 {
	return apu.IO0
}

func (apu *APU) CPUIO0W(data uint8) {
	//state machine to roughly emulate the APU boot rom: https://problemkaputt.de/fullsnes.htm#snesapumaincpucommunicationport
	switch apu.state {
	case stateInit:
		apu.stateInit(data)
	case stateTransfer:
		apu.stateTransfer(data)
	case stateRunning:
		apu.stateRunning(data)
	}
}

func (apu *APU) CPUIO1R() uint8 {
	return apu.IO1
}

func (apu *APU) CPUIO1W(data uint8) {
	apu.IO1 = data
}

func (apu *APU) CPUIO2R() uint8 {
	return apu.IO2
}

func (apu *APU) CPUIO2W(data uint8) {
	apu.IO2 = data
}

func (apu *APU) CPUIO3R() uint8 {
	return apu.IO3
}

func (apu *APU) CPUIO3W(data uint8) {
	apu.IO3 = data
}
