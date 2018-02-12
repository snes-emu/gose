package cpu

import "github.com/snes-emu/gose/utils"

// eor16 performs a bitwise exclusive or for 16bits variables
func (cpu *CPU) eor16(data uint16) uint16 {

	return cpu.getCRegister() ^ data
}

// eor8 performs a bitwise and for 8bits variables
func (cpu *CPU) eor8(data uint8) uint8 {

	return cpu.getARegister() ^ data
}

// eor performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) eor(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.eor8(dataLo))
	} else {
		cpu.setCRegister(cpu.eor16(utils.ReadUint16(dataHi, dataLo)))
	}
}
