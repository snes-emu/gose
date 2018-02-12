package cpu

import "github.com/snes-emu/gose/utils"

// cpx16 does a 16bit comparison of the X register with the data
func (cpu *CPU) cpx16(data uint16) {
	result := cpu.getXRegister() - data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getXRegister() >= data
}

// cpx8 does a 8bit comparison of the X register with the data
func (cpu *CPU) cpx8(data uint8) {
	result := cpu.getXLRegister() - data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getXLRegister() >= data
}

// cpx compare the X register to the data handling the 16bit/8bit distinction
func (cpu *CPU) cpx(dataHi, dataLo uint8) {
	if cpu.xFlag {
		cpu.cpx8(dataLo)
	} else {
		cpu.cpx16(utils.ReadUint16(dataHi, dataLo))
	}
}
