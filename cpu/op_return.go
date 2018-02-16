package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op40() {
	cpu.plp()
	addressLo := cpu.pullStack()
	addressHi := cpu.pullStack()
	cpu.PC = utils.JoinUint16(addressHi, addressLo)
	if !cpu.eFlag {
		cpu.K = cpu.pullStack()
	}
	cpu.cycles += 7 - utils.BoolToUint16[cpu.eFlag]

}

func (cpu *CPU) op6B() {
	PCLo := cpu.pullStack()
	PCHi := cpu.pullStack()
	cpu.K = cpu.pullStack()
	cpu.cycles += 6
	cpu.PC = utils.JoinUint16(PCHi, PCLo) + 1

}

func (cpu *CPU) op60() {
	PCLo := cpu.pullStack()
	PCHi := cpu.pullStack()
	cpu.cycles += 6
	cpu.PC = utils.JoinUint16(PCHi, PCLo) + 1
}
