package cpu

import "github.com/snes-emu/gose/utils"

// stx16 stores the x register in the memory
func (cpu *CPU) stx16(haddr, laddr uint32) {

	dataHi, dataLo := utils.SplitUint16(cpu.getXRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stx8 stores the lower part of the x register in the memory
func (cpu *CPU) stx8(addr uint32) {

	cpu.memory.SetByte(cpu.getXLRegister(), addr)
}

// stx stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stx(haddr, laddr uint32) {
	if cpu.xFlag {
		cpu.stx8(laddr)
	} else {
		cpu.stx16(haddr, laddr)
	}
}

func (cpu *CPU) op86() {
	haddr, laddr := cpu.admDirectP()
	cpu.stx(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8E() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.stx(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) op96() {
	haddr, laddr := cpu.admDirectYP()
	cpu.stx(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}
