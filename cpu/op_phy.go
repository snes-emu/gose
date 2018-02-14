package cpu

import "github.com/snes-emu/gose/utils"

// phy16 push the Y register onto the stack
func (cpu *CPU) phy16() {
	dataHi, dataLo := utils.WriteUint16(cpu.getYRegister())

	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
}

// phy8 push the lower bit of the Y register onto the stack
func (cpu *CPU) phy8() {
	data := cpu.getYLRegister()

	cpu.pushStack(data)
}

func (cpu *CPU) phy() {
	if cpu.xFlag {
		cpu.phy8()
	} else {
		cpu.phy16()
	}
}
