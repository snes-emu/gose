package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op6B() {
	PCLo := cpu.pullStack()
	PCHi := cpu.pullStack()
	cpu.K = cpu.pullStack()
	cpu.cycles += 6
	cpu.PC = utils.ReadUint16(PCHi, PCLo) + 1

}
