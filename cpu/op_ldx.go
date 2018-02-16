package cpu

import "github.com/snes-emu/gose/utils"

// ldx16 load data into the x register
func (cpu *CPU) ldx16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data != 0

	cpu.setXRegister(data)
}

// ldx8 load data into the lower bits of the x register
func (cpu *CPU) ldx8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data != 0

	cpu.setXLRegister(data)
}

// ldx load data into the x register taking care of 16bit/8bit cases
func (cpu *CPU) ldx(dataHi, dataLo uint8) {
	if cpu.xFlag {
		cpu.ldx8(dataLo)
	} else {
		cpu.ldx16(utils.JoinUint16(dataHi, dataLo))
	}
}

func (cpu *CPU) opA2() {
	dataHi, dataLo := cpu.admImmediateX()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opA6() {
	dataHi, dataLo := cpu.admDirect()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opAE() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) opB6() {
	dataHi, dataLo := cpu.admDirectY()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opBE() {
	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 6 - 2*utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}
