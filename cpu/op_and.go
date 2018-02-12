package cpu

import "github.com/snes-emu/gose/utils"

// and16 performs a bitwise and for 16bits variables
func (cpu *CPU) and16(data uint16) uint16 {

	return cpu.getCRegister() & data
}

// and8 performs a bitwise and for 8bits variables
func (cpu *CPU) and8(data uint8) uint8 {

	return cpu.getARegister() & data
}

// and performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) and(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.and8(dataLo))
	} else {
		cpu.setCRegister(cpu.and16(utils.ReadUint16(dataHi, dataLo)))
	}
}
