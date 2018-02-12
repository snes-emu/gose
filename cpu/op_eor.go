package cpu

import "github.com/snes-emu/gose/utils"

// eor16 performs a bitwise exclusive or for 16bits variables
func (cpu *CPU) eor16(data uint16) uint16 {
	result := cpu.getCRegister() ^ data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	return result
}

// eor8 performs a bitwise and for 8bits variables
func (cpu *CPU) eor8(data uint8) uint8 {
	result := cpu.getARegister() ^ data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	return result
}

// eor performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) eor(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.eor8(dataLo))
	} else {
		cpu.setCRegister(cpu.eor16(utils.ReadUint16(dataHi, dataLo)))
	}
}

func (cpu *CPU) op41() {
	dataHi, dataLo := cpu.admPDirectX()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op43() {
	dataHi, dataLo := cpu.admStackS()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op45() {
	dataHi, dataLo := cpu.admDirect()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op47() {
	dataHi, dataLo := cpu.admBDirect()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op49() {
	dataHi, dataLo := cpu.admImmediate()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op4D() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op4F() {
	dataHi, dataLo := cpu.admLong()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op51() {
	dataHi, dataLo := cpu.admPDirectY()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op52() {
	dataHi, dataLo := cpu.admPDirect()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op53() {
	dataHi, dataLo := cpu.admPStackSY()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op55() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op57() {
	dataHi, dataLo := cpu.admBDirectY()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op59() {
	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op5D() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op5F() {
	dataHi, dataLo := cpu.admLongX()
	cpu.eor(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}
