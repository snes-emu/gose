package cpu

import "github.com/snes-emu/gose/utils"

// cpy16 does a 16bit comparison of the Y register with the data
func (cpu *CPU) cpy16(data uint16) {
	result := cpu.getYRegister() - data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getYRegister() >= data
}

// cpy8 does a 8bit comparison of the Y register with the data
func (cpu *CPU) cpy8(data uint8) {
	result := cpu.getYLRegister() - data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getYLRegister() >= data
}

// cpy compare the Y register to the data handling the 16bit/8bit distinction
func (cpu *CPU) cpy(dataHi, dataLo uint8) {
	if cpu.xFlag {
		cpu.cpy8(dataLo)
	} else {
		cpu.cpy16(utils.ReadUint16(dataHi, dataLo))
	}
}
