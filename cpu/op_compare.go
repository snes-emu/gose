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
func (cpu *CPU) cmp(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.cmp8(dataLo)
	} else {
		cpu.cmp16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opC1() {

	dataLo, dataHi := cpu.admPDirectX()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opC3() {

	dataLo, dataHi := cpu.admStackS()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opC5() {

	dataLo, dataHi := cpu.admDirect()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opC7() {

	dataLo, dataHi := cpu.admBDirect()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opC9() {

	dataLo, dataHi := cpu.admImmediateM()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opCD() {

	dataLo, dataHi := cpu.admAbsolute()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) opCF() {

	dataLo, dataHi := cpu.admLong()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) opD1() {

	dataLo, dataHi := cpu.admPDirectY()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] + utils.BoolToUint16[cpu.xFlag]*(utils.BoolToUint16[cpu.pFlag]-1)
	cpu.PC += 2
}

func (cpu *CPU) opD2() {

	dataLo, dataHi := cpu.admPDirect()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opD3() {

	dataLo, dataHi := cpu.admPStackSY()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opD5() {

	dataLo, dataHi := cpu.admDirectX()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opD7() {

	dataLo, dataHi := cpu.admBDirectY()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opD9() {

	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag]*(utils.BoolToUint16[cpu.pFlag]-1)
	cpu.PC += 3
}

func (cpu *CPU) opDD() {

	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag]*(utils.BoolToUint16[cpu.pFlag]-1)
	cpu.PC += 3
}

func (cpu *CPU) opDF() {

	dataLo, dataHi := cpu.admLongX()
	cpu.cmp(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
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
func (cpu *CPU) cpx(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.cpx8(dataLo)
	} else {
		cpu.cpx16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opE0() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.cpx(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opE4() {
	dataLo, dataHi := cpu.admDirect()
	cpu.cpx(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opEC() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.cpx(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
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
func (cpu *CPU) cpy(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.cpy8(dataLo)
	} else {
		cpu.cpy16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opC0() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.cpy(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opC4() {
	dataLo, dataHi := cpu.admDirect()
	cpu.cpy(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opCC() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.cpy(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}
