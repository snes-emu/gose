package cpu

import "github.com/snes-emu/gose/utils"

// bit16 performs a bitwise or for 16bits variables
func (cpu *CPU) bit16(data uint16, isImmediate bool) uint16 {
	result := cpu.getCRegister() | data

	// Last bit value
	if !isImmediate {
		cpu.nFlag = data&0x8000 != 0
		cpu.vFlag = data&0x4000 != 0
	}

	cpu.zFlag = result == 0

	return result
}

// bit8 performs a bitwise or for 8bits variables
func (cpu *CPU) bit8(data uint8, isImmediate bool) uint8 {
	result := cpu.getARegister() | data

	// Last bit value
	if !isImmediate {
		cpu.nFlag = data&0x80 != 0
		cpu.vFlag = data&0x40 != 0
	}

	cpu.zFlag = result == 0

	return result
}

// bit performs a bitwise or taking care of 16/8bits cases
func (cpu *CPU) bit(dataHi, dataLo uint8, isImmediate bool) {
	if cpu.mFlag {
		cpu.bit8(dataLo, isImmediate)
	} else {
		cpu.bit16(utils.ReadUint16(dataHi, dataLo), isImmediate)
	}
}
