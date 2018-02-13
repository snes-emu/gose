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
		cpu.setCRegister(cpu.adc16(utils.ReadUint16(dataHi, dataLo)))
	}
}

func (cpu *CPU) op61() {

	dataHi, dataLo := cpu.admPDirectX()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op63() {

	dataHi, dataLo := cpu.admStackS()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op65() {

	dataHi, dataLo := cpu.admDirect()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op67() {

	dataHi, dataLo := cpu.admBDirect()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op69() {

	dataHi, dataLo := cpu.admImmediateM()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op6D() {

	dataHi, dataLo := cpu.admAbsolute()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op71() {

	dataHi, dataLo := cpu.admPDirectY()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op72() {
	dataHi, dataLo := cpu.admPDirect()

	cpu.adc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op73() {

	dataHi, dataLo := cpu.admPStackSY()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op75() {
	dataHi, dataLo := cpu.admDirectX()

	cpu.adc(dataHi, dataLo)

	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op77() {

	dataHi, dataLo := cpu.admDirectX()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op79() {
	dataHi, dataLo := cpu.admAbsoluteX()

	cpu.adc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.mFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op7D() {

	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.adc(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]

}

func (cpu *CPU) op7F() {
	dataHi, dataLo := cpu.admLongX()

	cpu.adc(dataHi, dataLo)

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}
