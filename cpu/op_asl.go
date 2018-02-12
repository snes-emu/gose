package cpu

import "github.com/snes-emu/gose/utils"

// asl16 performs a left shift on the 16 bit accumulator
func (cpu *CPU) asl16(data uint16) uint16 {
	result := cpu.getCRegister() << 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x8000 != 0

	// Last asl value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	return result
}

// asl8 performs a left shift on the lower 8 bit accumulator
func (cpu *CPU) asl8(data uint8) uint8 {
	result := cpu.getARegister() << 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getARegister()&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	return result
}

// asl performs a left shift taking care of 16/8bits cases
func (cpu *CPU) asl(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.asl8(dataLo)
	} else {
		cpu.asl16(utils.ReadUint16(dataHi, dataLo))
	}
}

func (cpu *CPU) op06() {
	dataHi, dataLo := cpu.admDirect()
	cpu.asl(dataHi, dataLo)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op0A() {
	dataHi, dataLo := cpu.admAccumulator()
	cpu.asl(dataHi, dataLo)
	cpu.cycles += 2
}

func (cpu *CPU) op0E() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.asl(dataHi, dataLo)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) op16() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.asl(dataHi, dataLo)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

func (cpu *CPU) op1E() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.asl(dataHi, dataLo)
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
}
