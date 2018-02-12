package cpu

import "github.com/snes-emu/gose/utils"

// ora16 performs a bitwise or for 16bits variables
func (cpu *CPU) ora16(data uint16) uint16 {

	return cpu.getCRegister() | data
}

// ora8 performs a bitwise or for 8bits variables
func (cpu *CPU) ora8(data uint8) uint8 {

	return cpu.getARegister() | data
}

// ora performs a bitwise or taking care of 16/8bits cases
func (cpu *CPU) ora(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.ora8(dataLo))
	} else {
		cpu.setCRegister(cpu.ora16(utils.ReadUint16(dataHi, dataLo)))
	}
}
