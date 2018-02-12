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

func (cpu *CPU) opC0() {
	dataHi, dataLo := cpu.admImmediate()
	cpu.cpy(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opC4() {
	dataHi, dataLo := cpu.admDirect()
	cpu.cpy(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opCC() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.cpy(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
}
