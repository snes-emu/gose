package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) tcd() {
	cpu.D = cpu.C
	// Last bit value
	cpu.nFlag = cpu.D&0x8000 != 0
	cpu.zFlag = cpu.D == 0
	cpu.cycles += 2
}

func (cpu *CPU) op5B() {
	cpu.tcd()
}

func (cpu *CPU) tcs() {
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

func (cpu *CPU) op1B() {
	cpu.tcs()
}

func (cpu *CPU) tdc() {
	cpu.C = cpu.D
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	cpu.cycles += 2
}

func (cpu *CPU) op7B() {
	cpu.tdc()
}

func (cpu *CPU) tsc() {
	cpu.C = cpu.S
	// Last bit value
	cpu.nFlag = cpu.S&0x8000 != 0
	cpu.zFlag = cpu.S == 0
	cpu.cycles += 2
}

func (cpu *CPU) op3B() {
	cpu.tsc()
}
