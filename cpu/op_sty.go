package cpu

import "github.com/snes-emu/gose/utils"

// sty16 stores the x register in the memory
func (cpu *CPU) sty16(haddr, laddr uint32) {

	dataHi, dataLo := utils.SplitUint16(cpu.getYRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sty8 stores the lower part of the x register in the memory
func (cpu *CPU) sty8(addr uint32) {

	cpu.memory.SetByte(cpu.getYLRegister(), addr)
}

// sty stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sty(haddr, laddr uint32) {
	if cpu.xFlag {
		cpu.sty8(laddr)
	} else {
		cpu.sty16(haddr, laddr)
	}
}

func (cpu *CPU) op84() {
	haddr, laddr := cpu.admDirectP()
	cpu.sty(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8C() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.sty(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) op94() {
	haddr, laddr := cpu.admDirectXP()
	cpu.sty(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}
