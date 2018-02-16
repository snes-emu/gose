package cpu

import (
	"github.com/snes-emu/gose/utils"
)

//op62 pushes 16bit data into the stack, called thanks to the next 8bit value
func (cpu *CPU) op62(data uint16) {
	dataHi, dataLo := utils.SplitUint16(data)
	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
	cpu.cycles += 6
}
