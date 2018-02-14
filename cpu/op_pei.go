package cpu

import (
	"github.com/snes-emu/gose/utils"
)

//opD4 pushes 16bit data into the stack, called thanks to the next 8bit value
func (cpu *CPU) opD4() {
	dataHi, dataLo := cpu.admDirect()
	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
	cpu.cycles += 6 + utils.BoolToUint16[cpu.getDLRegister() == 0]
}
