package cpu

import (
	"github.com/snes-emu/gose/utils"
)

func (cpu *CPU) op40() {
	cpu.plp()
	addressLo := cpu.pullStack()
	addressHi := cpu.pullStack()
	cpu.PC = utils.ReadUint16(addressHi, addressLo)
	if !cpu.eFlag {
		cpu.K = cpu.pullStack()
	}
	cpu.cycles += 7 - utils.BoolToUint16[cpu.eFlag]

}
