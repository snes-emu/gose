package cpu

import "github.com/snes-emu/gose/utils"

// pha16 push the accumulator onto the stack
func (cpu *CPU) pha16() {
	dataHi, dataLo := utils.SplitUint16(cpu.getCRegister())

	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
}

// pha8 push the lower bit of the accumulator onto the stack
func (cpu *CPU) pha8() {
	data := cpu.getARegister()

	cpu.pushStack(data)
}

func (cpu *CPU) pha() {
	if cpu.mFlag {
		cpu.pha8()
	} else {
		cpu.pha16()
	}
}

func (cpu *CPU) op48() {
	cpu.pha()
}
