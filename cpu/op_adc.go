package cpu

import "github.com/snes-emu/gose/utils"

// adc16 performs an add with carry 16bit operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) adc16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in adc needs to be implemented")
		result = (cpu.getCRegister() & 0x000f) + (data & 0x000f) + utils.BoolToUint16[cpu.cFlag] + (cpu.getCRegister() & 0x00f0) + (data & 0x00f0) + (cpu.C & 0x0f00) + (data & 0x0f00) + (cpu.getCRegister() & 0xf000) + (data & 0xf000)
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

func (cpu *CPU) op61() {
	// TODO

	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}
