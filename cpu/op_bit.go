package cpu

import "github.com/snes-emu/gose/utils"

// bit16 performs a bitwise and for 16bits variables
func (cpu *CPU) bit16(data uint16, isImmediate bool) uint16 {
	result := cpu.getCRegister() & data

	// Last bit value
	if !isImmediate {
		cpu.nFlag = data&0x8000 != 0
		cpu.vFlag = data&0x4000 != 0
	}

	cpu.zFlag = result == 0

	return result
}

// bit8 performs a bitwise and for 8bits variables
func (cpu *CPU) bit8(data uint8, isImmediate bool) uint8 {
	result := cpu.getARegister() & data

	// Last bit value
	if !isImmediate {
		cpu.nFlag = data&0x80 != 0
		cpu.vFlag = data&0x40 != 0
	}

	cpu.zFlag = result == 0

	return result
}

// bit performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) bit(dataHi, dataLo uint8, isImmediate bool) {
	if cpu.mFlag {
		cpu.bit8(dataLo, isImmediate)
	} else {
		cpu.bit16(utils.JoinUint16(dataHi, dataLo), isImmediate)
	}
}

func (cpu *CPU) op24() {
	dataHi, dataLo := cpu.admDirect()
	cpu.bit(dataHi, dataLo, false)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op2C() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.bit(dataHi, dataLo, false)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op34() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.bit(dataHi, dataLo, false)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op3C() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.bit(dataHi, dataLo, false)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op89() {
	dataHi, dataLo := cpu.admImmediateM()
	cpu.bit(dataHi, dataLo, true)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}
