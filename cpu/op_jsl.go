package cpu

import "github.com/snes-emu/gose/utils"

// jsl jumps to a subroutine long
func (cpu *CPU) jsl(haddr uint8, laddr uint16) {
	cpu.pushStack(cpu.getKRegister())
	hiaddr, loaddr := utils.WriteUint16(cpu.getPCRegister() + 3)

	cpu.pushStack(hiaddr)
	cpu.pushStack(loaddr)

	cpu.jmpLong(haddr, laddr)
}

func (cpu *CPU) op22() {
	haddr, laddr := cpu.admLongJ()
	cpu.jsl(haddr, laddr)
	cpu.cycles += 3
	cpu.PC += 3
}
