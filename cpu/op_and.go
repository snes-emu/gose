package cpu

import "github.com/snes-emu/gose/utils"

// and16 performs a bitwise and for 16bits variables
func (cpu *CPU) and16(data uint16) uint16 {
	result := cpu.getCRegister() & data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	return result
}

// and8 performs a bitwise and for 8bits variables
func (cpu *CPU) and8(data uint8) uint8 {
	result := cpu.getARegister() & data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	return result
}

// and performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) and(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.and8(dataLo))
	} else {
		cpu.setCRegister(cpu.and16(utils.ReadUint16(dataHi, dataLo)))
	}
}

func (cpu *CPU) op21() {
	dataHi, dataLo := cpu.admPDirectX()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op23() {
	dataHi, dataLo := cpu.admStackS()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op25() {
	dataHi, dataLo := cpu.admDirect()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op27() {
	dataHi, dataLo := cpu.admBDirect()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op29() {
	dataHi, dataLo := cpu.admImmediate()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op2D() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op2F() {
	dataHi, dataLo := cpu.admLong()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op31() {
	dataHi, dataLo := cpu.admPDirectY()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op32() {
	dataHi, dataLo := cpu.admPDirect()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op33() {
	dataHi, dataLo := cpu.admPStackSY()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op35() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op37() {
	dataHi, dataLo := cpu.admBDirectY()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op39() {
	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op3D() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op3F() {
	dataHi, dataLo := cpu.admLongX()
	cpu.and(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}
