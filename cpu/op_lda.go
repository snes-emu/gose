package cpu

import "github.com/snes-emu/gose/utils"

// lda16 load data into the lower bits of the accumulator
func (cpu *CPU) lda16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data != 0

	cpu.setCRegister(data)
}

// lda8 load data into the lower bits of the accumulator
func (cpu *CPU) lda8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data != 0

	cpu.setARegister(data)
}

// lda load data into the accumulator
func (cpu *CPU) lda(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.lda8(dataLo)
	} else {
		cpu.lda16(utils.ReadUint16(dataHi, dataLo))
	}
}

func (cpu *CPU) opA1() {
	dataHi, dataLo := cpu.admPDirectX()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA3() {
	dataHi, dataLo := cpu.admStackS()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opA5() {
	dataHi, dataLo := cpu.admDirect()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA7() {
	dataHi, dataLo := cpu.admBDirect()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA9() {
	dataHi, dataLo := cpu.admImmediateM()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opAD() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) opAF() {
	dataHi, dataLo := cpu.admLong()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) opB1() {
	dataHi, dataLo := cpu.admPDirectY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) opB2() {
	dataHi, dataLo := cpu.admPDirect()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB3() {
	dataHi, dataLo := cpu.admPStackSY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opB5() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB7() {
	dataHi, dataLo := cpu.admBDirectY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB9() {
	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opBD() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opBF() {
	dataHi, dataLo := cpu.admLongX()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}
