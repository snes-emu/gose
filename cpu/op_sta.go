package cpu

import "github.com/snes-emu/gose/utils"

// sta16 stores the accumulator in the memory
func (cpu *CPU) sta16(haddr, laddr uint32) {

	dataHi, dataLo := utils.SplitUint16(cpu.getCRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sta8 stores the lower part of the accumulator in the memory
func (cpu *CPU) sta8(addr uint32) {

	cpu.memory.SetByte(cpu.getARegister(), addr)
}

// sta stores the accumulator in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sta(haddr, laddr uint32) {
	if cpu.mFlag {
		cpu.sta8(laddr)
	} else {
		cpu.sta16(haddr, laddr)
	}
}

func (cpu *CPU) op81() {
	haddr, laddr := cpu.admPDirectXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op83() {
	haddr, laddr := cpu.admStackSP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op85() {
	haddr, laddr := cpu.admDirectP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op87() {
	haddr, laddr := cpu.admBDirectP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8D() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op8F() {
	haddr, laddr := cpu.admLongP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) op91() {
	haddr, laddr := cpu.admPDirectYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op92() {
	haddr, laddr := cpu.admPDirectP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op93() {
	haddr, laddr := cpu.admPStackSYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op95() {
	haddr, laddr := cpu.admDirectXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op97() {
	haddr, laddr := cpu.admBDirectYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op99() {
	haddr, laddr := cpu.admAbsoluteYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9D() {
	haddr, laddr := cpu.admAbsoluteXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9F() {
	haddr, laddr := cpu.admLongXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}
