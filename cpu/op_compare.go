package cpu

import "github.com/snes-emu/gose/utils"

// cmp16 does a 16bit comparison the accumulator to the data
func (cpu *CPU) cmp16(data uint16) {
	result := cpu.getCRegister() - data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getCRegister() >= data
}

// cmp8 does a 8bit comparison the accumulator to the data
func (cpu *CPU) cmp8(data uint8) {
	result := cpu.getARegister() - data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getARegister() >= data
}

// cmp compare the accumulator to the data handling the 16bit/8bit distinction
func (cpu *CPU) cmp(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.cmp8(dataLo)
	} else {
		cpu.cmp16(utils.JoinUint16(dataHi, dataLo))
	}
}

func (cpu *CPU) opC1() {

	dataHi, dataLo := cpu.admPDirectX()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opC3() {

	dataHi, dataLo := cpu.admStackS()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opC5() {

	dataHi, dataLo := cpu.admDirect()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opC7() {

	dataHi, dataLo := cpu.admBDirect()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opC9() {

	dataHi, dataLo := cpu.admImmediateM()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opCD() {

	dataHi, dataLo := cpu.admAbsolute()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opCF() {

	dataHi, dataLo := cpu.admLong()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opD1() {

	dataHi, dataLo := cpu.admPDirectY()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] + utils.BoolToUint16[cpu.xFlag]*(utils.BoolToUint16[cpu.pFlag]-1)
}

func (cpu *CPU) opD2() {

	dataHi, dataLo := cpu.admPDirect()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opD3() {

	dataHi, dataLo := cpu.admPStackSY()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opD5() {

	dataHi, dataLo := cpu.admDirectX()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opD7() {

	dataHi, dataLo := cpu.admBDirectY()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opD9() {

	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag]*(utils.BoolToUint16[cpu.pFlag]-1)
}

func (cpu *CPU) opDD() {

	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag]*(utils.BoolToUint16[cpu.pFlag]-1)
}

func (cpu *CPU) opDF() {

	dataHi, dataLo := cpu.admLongX()
	cpu.cmp(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}

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
		cpu.cpx16(utils.JoinUint16(dataHi, dataLo))
	}
}

func (cpu *CPU) opE0() {
	dataHi, dataLo := cpu.admImmediateX()
	cpu.cpx(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opE4() {
	dataHi, dataLo := cpu.admDirect()
	cpu.cpx(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opEC() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.cpx(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
}

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
		cpu.cpy16(utils.JoinUint16(dataHi, dataLo))
	}
}

func (cpu *CPU) opC0() {
	dataHi, dataLo := cpu.admImmediateX()
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
