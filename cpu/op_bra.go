package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op80() {
	offset := cpu.admRelative8()
	cpu.cycles += 3 + utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += offset + 2
}
