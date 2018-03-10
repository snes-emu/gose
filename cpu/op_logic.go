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
func (cpu *CPU) and(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.and8(dataLo))
	} else {
		cpu.setCRegister(cpu.and16(utils.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op21() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op23() {
	dataLo, dataHi := cpu.admStackS()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op25() {
	dataLo, dataHi := cpu.admDirect()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op27() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op29() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op2D() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op2F() {
	dataLo, dataHi := cpu.admLong()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) op31() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) op32() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op33() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op35() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op37() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op39() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op3D() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op3F() {
	dataLo, dataHi := cpu.admLongX()
	cpu.and(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

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
func (cpu *CPU) eor(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.eor8(dataLo))
	} else {
		cpu.setCRegister(cpu.eor16(utils.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op41() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op43() {
	dataLo, dataHi := cpu.admStackS()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op45() {
	dataLo, dataHi := cpu.admDirect()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op47() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op49() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op4D() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op4F() {
	dataLo, dataHi := cpu.admLong()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) op51() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) op52() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op53() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op55() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op57() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op59() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op5D() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op5F() {
	dataLo, dataHi := cpu.admLongX()
	cpu.eor(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

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
func (cpu *CPU) ora(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.ora8(dataLo))
	} else {
		cpu.setCRegister(cpu.ora16(utils.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op01() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op03() {
	dataLo, dataHi := cpu.admStackS()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op05() {
	dataLo, dataHi := cpu.admDirect()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op07() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op09() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op0D() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op0F() {
	dataLo, dataHi := cpu.admLong()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) op11() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) op12() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op13() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op15() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op17() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op19() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op1D() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) op1F() {
	dataLo, dataHi := cpu.admLongX()
	cpu.ora(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}
