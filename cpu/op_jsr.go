package cpu

import "github.com/snes-emu/gose/utils"

// jsr jumps to a subroutine
func (cpu *CPU) jsr(addr uint16) {
	haddr, laddr := utils.WriteUint16(cpu.getPCRegister() + 2)

	cpu.pushStack(haddr)
	cpu.pushStack(laddr)

	cpu.jmp(addr)
}
