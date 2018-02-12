package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op70() {
	offset := cpu.admRelative8()
	t := cpu.vFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}
