package cpu

import (
	"github.com/snes-emu/gose/utils"
)

func (cpu *CPU) op1B() {
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	if cpu.eFlag {
		_, dataLo := utils.WriteUint16(cpu.C)
		cpu.S = utils.ReadUint16(0x01, dataLo)
	} else {
		cpu.S = cpu.C
	}
	cpu.cycles += 2
}
