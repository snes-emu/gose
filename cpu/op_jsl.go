package cpu

import "github.com/snes-emu/gose/utils"

// jsl jumps to a subroutine long
func (cpu *CPU) jsl(addr uint16) {
	cpu.pushStack(cpu.getKRegister())
	haddr, laddr := utils.WriteUint16(cpu.getPCRegister() + 3)

	cpu.pushStack(haddr)
	cpu.pushStack(laddr)

	cpu.PC = addr
}
