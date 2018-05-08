package apu

import "github.com/snes-emu/gose/io"

type APU struct {
	CommunicationPorts [4]*port
	Registers          [4]*io.Register
}

func New() *APU {
	apu := &APU{}
	for i := 0; i < len(apu.CommunicationPorts); i++ {
		apu.CommunicationPorts[i] = newPort()
		apu.Registers[i] = io.NewRegister(apu.CommunicationPorts[i].APUOutput, apu.CommunicationPorts[i].APUInput)
	}
	return apu
}
