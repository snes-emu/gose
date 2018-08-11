package apu

import (
	"fmt"

	"github.com/snes-emu/gose/io"
)

type APU struct {
	CommunicationPorts [4]*port
	Registers          [4]*io.Register
}

func New() *APU {
	apu := &APU{}
	for i := 0; i < len(apu.CommunicationPorts); i++ {
		apu.CommunicationPorts[i] = newPort()
		apu.Registers[i] = io.NewRegister(apu.CommunicationPorts[i].CPUOutput, apu.CommunicationPorts[i].CPUInput, fmt.Sprintf("APUFAKE%v", i))
	}
	apu.init()
	return apu
}

func (apu *APU) init() {
	apu.CommunicationPorts[0].apuInput(0xAA)
	apu.CommunicationPorts[1].apuInput(0xBB)
	apu.CommunicationPorts[2].init = true
	apu.CommunicationPorts[3].init = true

}
