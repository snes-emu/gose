package cpu

import (
	"github.com/snes-emu/gose/utils"
)

//opF4 pushes the next 16bit value into the stack
func (cpu *CPU) opF4(data uint16) {
	dataHi, dataLo := utils.SplitUint16(data)
	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
	cpu.cycles += 5
}
