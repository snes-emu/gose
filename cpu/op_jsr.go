package cpu

import "github.com/snes-emu/gose/utils"

// jsr jumps to a subroutine
func (cpu *CPU) jsr(addr uint16) {
	haddr, laddr := utils.SplitUint16(cpu.getPCRegister() + 2)

	cpu.pushStack(haddr)
	cpu.pushStack(laddr)

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
	cpu.jsr(addr)
	cpu.cycles += 8
	cpu.PC += 3
}
