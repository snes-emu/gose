package cpu

import "github.com/snes-emu/gose/utils"

// phx16 push the X register onto the stack
func (cpu *CPU) phx16() {
	dataHi, dataLo := utils.WriteUint16(cpu.getXRegister())

	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
}

// phx8 push the lower bit of the X register onto the stack
func (cpu *CPU) phx8() {
	data := cpu.getXLRegister()

	cpu.pushStack(data)
}

func (cpu *CPU) phx() {
	if cpu.xFlag {
		cpu.phx8()
	} else {
		cpu.phx16()
	}
}

func (cpu *CPU) opDA() {
	cpu.phx()
}
