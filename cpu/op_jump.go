package cpu

import "github.com/snes-emu/gose/utils"

// jmp jumps to the address specified by the addressing mode
func (cpu *CPU) jmp(addr uint16) {
	cpu.PC = addr
}

// jmpLong jumps to the address specified by the long addressing
func (cpu *CPU) jmpLong(laddr uint16, haddr uint8) {
	cpu.K = haddr
	cpu.PC = laddr
}

func (cpu *CPU) op4C() {
	addr := cpu.admAbsoluteJ()
	cpu.jmp(addr)
	cpu.cycles += 3
	cpu.PC += 3
}

func (cpu *CPU) op5C() {
	laddr, haddr := cpu.admLongJ()
	cpu.jmpLong(laddr, haddr)
	cpu.cycles += 4
	cpu.PC += 4
}

func (cpu *CPU) op6C() {
	addr := cpu.admPAbsoluteJ()
	cpu.jmp(addr)
	cpu.cycles += 5
	cpu.PC += 3
}

func (cpu *CPU) op7C() {
	addr := cpu.admPAbsoluteXJ()
	cpu.jmp(addr)
	cpu.cycles += 6
	cpu.PC += 3
}

func (cpu *CPU) opDC() {
	laddr, haddr := cpu.admBAbsoluteJ()
	cpu.jmpLong(haddr, laddr)
	cpu.cycles += 6
	cpu.PC += 3
}

// jsl jumps to a subroutine long
func (cpu *CPU) jsl(laddr uint16, haddr uint8) {
	laddr2, haddr2 := utils.SplitUint16(cpu.getPCRegister() + 3)
	cpu.pushStackNew24(laddr2, haddr2, cpu.getKRegister())

	cpu.jmpLong(laddr, haddr)
}

func (cpu *CPU) op22() {
	laddr, haddr := cpu.admLongJ()
	cpu.jsl(laddr, haddr)
	cpu.cycles += 3
	cpu.PC += 3
}

// jsr jumps to a subroutine
func (cpu *CPU) jsr(addr uint16) {
	laddr, haddr := utils.SplitUint16(cpu.getPCRegister() + 2)

	cpu.pushStack(haddr)
	cpu.pushStack(laddr)

	cpu.jmp(addr)
}

// jsr jumps to a subroutine for new addressing mode
func (cpu *CPU) jsrNew(addr uint16) {
	laddr, haddr := utils.SplitUint16(cpu.getPCRegister() + 2)

	cpu.pushStackNew16(laddr, haddr)

	cpu.jmp(addr)
}

func (cpu *CPU) op20() {
	addr := cpu.admAbsoluteJ()
	cpu.jsr(addr)
	cpu.cycles += 6
	cpu.PC += 3
}

func (cpu *CPU) opFC() {
	addr := cpu.admPAbsoluteXJ()
	cpu.jsrNew(addr)
	cpu.cycles += 8
	cpu.PC += 3
}
