package cpu

import "github.com/snes-emu/gose/utils"

// ora16 performs a bitwise or for 16bits variables
func (cpu *CPU) ora16(data uint16) uint16 {
	result := cpu.getCRegister() | data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	return result
}

// ora8 performs a bitwise or for 8bits variables
func (cpu *CPU) ora8(data uint8) uint8 {
	result := cpu.getARegister() | data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	return result
}

// ora performs a bitwise or taking care of 16/8bits cases
func (cpu *CPU) ora(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.ora8(dataLo))
	} else {
		cpu.setCRegister(cpu.ora16(utils.ReadUint16(dataHi, dataLo)))
	}
}

func (cpu *CPU) op01() {
	dataHi, dataLo := cpu.admPDirectX()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op03() {
	dataHi, dataLo := cpu.admStackS()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op05() {
	dataHi, dataLo := cpu.admDirect()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op07() {
	dataHi, dataLo := cpu.admBDirect()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op09() {
	dataHi, dataLo := cpu.admImmediateM()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op0D() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op0F() {
	dataHi, dataLo := cpu.admLong()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op11() {
	dataHi, dataLo := cpu.admPDirectY()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op12() {
	dataHi, dataLo := cpu.admPDirect()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op13() {
	dataHi, dataLo := cpu.admPStackSY()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op15() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op17() {
	dataHi, dataLo := cpu.admBDirectY()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op19() {
	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op1D() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
}

func (cpu *CPU) op1F() {
	dataHi, dataLo := cpu.admLongX()
	cpu.ora(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
}
