package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op5B() {
	cpu.D = cpu.C
	// Last bit value
	cpu.nFlag = cpu.D&0x8000 != 0
	cpu.zFlag = cpu.D == 0
	cpu.cycles += 2
}

func (cpu *CPU) op1B() {
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	if cpu.eFlag {
		_, dataLo := utils.SplitUint16(cpu.C)
		cpu.S = utils.JoinUint16(0x01, dataLo)
	} else {
		cpu.S = cpu.C
	}
	cpu.cycles += 2
}

func (cpu *CPU) op7B() {
	cpu.C = cpu.D
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	cpu.cycles += 2
}

func (cpu *CPU) op3B() {
	cpu.C = cpu.S
	// Last bit value
	cpu.nFlag = cpu.S&0x8000 != 0
	cpu.zFlag = cpu.S == 0
	cpu.cycles += 2
}
