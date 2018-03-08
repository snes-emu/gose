package cpu

import "github.com/snes-emu/gose/utils"

// adc16 performs an add with carry 16bit operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) adc16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in adc needs to be implemented")
		//result = (cpu.getCRegister() & 0x000f) + (data & 0x000f) + utils.BoolToUint16[cpu.cFlag] + (cpu.getCRegister() & 0x00f0) + (data & 0x00f0) + (cpu.C & 0x0f00) + (data & 0x0f00) + (cpu.getCRegister() & 0xf000) + (data & 0xf000)

	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getCRegister() + data + utils.BoolToUint16[cpu.cFlag]
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Signed artihmetic overflow
		cpu.vFlag = (data^result)&^(data^cpu.getCRegister())&0x8000 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = (result < cpu.getCRegister()) || (result == cpu.getCRegister() && cpu.cFlag)

	}

	return result
}

// adc8 performs an add with carry 8bit operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) adc8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in adc needs to be implemented")
	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getARegister() + data + utils.BoolToUint8[cpu.cFlag]
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Signed artihmetic overflow
		cpu.vFlag = (data^result)&^(data^cpu.getARegister())&0x80 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = (result < cpu.getARegister()) || (result == cpu.getARegister() && cpu.cFlag)

	}

	return result
}

// sbc performs a substract with carry operation handling the 8/16 bit cases
func (cpu *CPU) adc(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.adc8(dataLo))
	} else {
		cpu.setCRegister(cpu.adc16(utils.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op61() {

	dataHi, dataLo := cpu.admPDirectX()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op63() {

	dataHi, dataLo := cpu.admStackS()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op65() {

	dataHi, dataLo := cpu.admDirect()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op67() {

	dataHi, dataLo := cpu.admBDirect()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op69() {

	dataHi, dataLo := cpu.admImmediateM()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op6D() {

	dataHi, dataLo := cpu.admAbsolute()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op6F() {

	dataHi, dataLo := cpu.admLong()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) op71() {

	dataHi, dataLo := cpu.admPDirectY()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) op72() {
	dataHi, dataLo := cpu.admPDirect()

	cpu.adc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op73() {

	dataHi, dataLo := cpu.admPStackSY()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op75() {
	dataHi, dataLo := cpu.admDirectX()

	cpu.adc(dataHi, dataLo)

	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op77() {

	dataHi, dataLo := cpu.admDirectX()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op79() {
	dataHi, dataLo := cpu.admAbsoluteX()

	cpu.adc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.mFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op7D() {

	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op7F() {
	dataHi, dataLo := cpu.admLongX()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

// sbc16 performs a substract with carry 16bit operation the formula is: accumulator = accumulator - data - 1 + carry
func (cpu *CPU) sbc16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in sbc needs to be implemented")

	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getCRegister() - data - 1 + utils.BoolToUint16[cpu.cFlag]
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Signed artihmetic overflow
		cpu.vFlag = ((data+1)^cpu.getCRegister())&^((data+1)^result)&0x8000 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = cpu.getCRegister() >= data

	}

	return result
}

// sbc8 performs a substract with carry 8bit operation the formula is: accumulator = accumulator - data - 1 + carry
func (cpu *CPU) sbc8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in sbc needs to be implemented")

	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getARegister() - data - 1 + utils.BoolToUint8[cpu.cFlag]
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Signed artihmetic overflow
		cpu.vFlag = ((data+1)^cpu.getARegister())&^((data+1)^result)&0x80 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = cpu.getARegister() >= data

	}

	return result
}

// sbc performs a substract with carry operation handling the 8/16 bit cases
func (cpu *CPU) sbc(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.sbc8(dataLo))
	} else {
		cpu.setCRegister(cpu.sbc16(utils.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) opE1() {

	dataHi, dataLo := cpu.admPDirectX()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opE3() {

	dataHi, dataLo := cpu.admStackS()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opE5() {

	dataHi, dataLo := cpu.admDirect()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opE7() {

	dataHi, dataLo := cpu.admBDirect()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opE9() {

	dataHi, dataLo := cpu.admImmediateM()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opED() {

	dataHi, dataLo := cpu.admAbsolute()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) opEF() {

	dataHi, dataLo := cpu.admLong()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) opF1() {

	dataHi, dataLo := cpu.admPDirectY()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) opF2() {
	dataHi, dataLo := cpu.admPDirect()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opF3() {

	dataHi, dataLo := cpu.admPStackSY()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opF5() {
	dataHi, dataLo := cpu.admDirectX()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opF7() {

	dataHi, dataLo := cpu.admDirectX()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opF9() {
	dataHi, dataLo := cpu.admAbsoluteX()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.mFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opFD() {

	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opFF() {
	dataHi, dataLo := cpu.admLongX()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}
