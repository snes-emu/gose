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
		result = cpu.getCRegister() + data + utils.BoolToUint16[cpu.cFlag]
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
		result = cpu.getARegister() + data + utils.BoolToUint8[cpu.cFlag]
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

func (cpu *CPU) opEF() {
	dataHi, dataLo := cpu.admLong()

	if cpu.mFlag {
		cpu.setARegister(cpu.sbc8(dataLo))
	} else {
		cpu.setCRegister(cpu.sbc16(utils.ReadUint16(dataHi, dataLo)))
	}

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opF2() {
	dataHi, dataLo := cpu.admPDirect()

	if cpu.mFlag {
		cpu.setARegister(cpu.sbc8(dataLo))
	} else {
		cpu.setCRegister(cpu.sbc16(utils.ReadUint16(dataHi, dataLo)))
	}

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opF5() {
	dataHi, dataLo := cpu.admDirectX()

	if cpu.mFlag {
		cpu.setARegister(cpu.sbc8(dataLo))
	} else {
		cpu.setCRegister(cpu.sbc16(utils.ReadUint16(dataHi, dataLo)))
	}

	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) opF9() {
	dataHi, dataLo := cpu.admAbsoluteX()

	if cpu.mFlag {
		cpu.setARegister(cpu.sbc8(dataLo))
	} else {
		cpu.setCRegister(cpu.sbc16(utils.ReadUint16(dataHi, dataLo)))
	}

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.mFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) opFF() {
	dataHi, dataLo := cpu.admLongX()

	if cpu.mFlag {
		cpu.setARegister(cpu.sbc8(dataLo))
	} else {
		cpu.setCRegister(cpu.sbc16(utils.ReadUint16(dataHi, dataLo)))
	}

	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}
