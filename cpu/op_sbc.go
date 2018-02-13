package cpu

import "github.com/snes-emu/gose/utils"

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
		cpu.vFlag = (data^result)&^(data^cpu.getCRegister())&0x8000 != 0
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
		cpu.vFlag = (data^result)&^(data^cpu.getARegister())&0x80 != 0
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
		cpu.setCRegister(cpu.sbc16(utils.ReadUint16(dataHi, dataLo)))
	}
}

func (cpu *CPU) opE1() {

	dataHi, dataLo := cpu.admPDirectX()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opE3() {

	dataHi, dataLo := cpu.admStackS()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opE5() {

	dataHi, dataLo := cpu.admDirect()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opE7() {

	dataHi, dataLo := cpu.admBDirect()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opE9() {

	dataHi, dataLo := cpu.admImmediateM()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opED() {

	dataHi, dataLo := cpu.admAbsolute()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opF1() {

	dataHi, dataLo := cpu.admPDirectY()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) opF2() {
	dataHi, dataLo := cpu.admPDirect()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opF3() {

	dataHi, dataLo := cpu.admPStackSY()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opF5() {
	dataHi, dataLo := cpu.admDirectX()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opF7() {

	dataHi, dataLo := cpu.admDirectX()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opF9() {
	dataHi, dataLo := cpu.admAbsoluteX()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.mFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) opFD() {

	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.sbc(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]

}

func (cpu *CPU) opFF() {
	dataHi, dataLo := cpu.admLongX()

	cpu.sbc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}
