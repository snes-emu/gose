package cpu

import "github.com/snes-emu/gose/utils"

// stz16 stores 0 in the memory
func (cpu *CPU) stz16(haddr, laddr uint32) {

	dataHi, dataLo := utils.WriteUint16(0x0000)

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stz8 stores 0 in the memory
func (cpu *CPU) stz8(addr uint32) {

	cpu.memory.SetByte(0x00, addr)
}

// stz stores 0 in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stz(haddr, laddr uint32) {
	if cpu.xFlag {
		cpu.stz8(laddr)
	} else {
		cpu.stz16(haddr, laddr)
	}
}

func (cpu *CPU) op64() {
	haddr, laddr := cpu.admDirectP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op74() {
	haddr, laddr := cpu.admDirectXP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op9C() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9E() {
	haddr, laddr := cpu.admAbsoluteXP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}
