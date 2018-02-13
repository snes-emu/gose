package cpu

import "github.com/snes-emu/gose/utils"

// jmp jumps to the address specified by the addressing mode
func (cpu *CPU) jmp(addr uint16) {
	cpu.PC = addr
}

// jmpLong jumps to the address specified by the long addressing
func (cpu *CPU) jmpLong(addr uint32) {
	hiaddr, miaddr, loaddr := utils.WriteUint32(addr)

	cpu.K = hiaddr
	cpu.PC = utils.ReadUint16(miaddr, loaddr)
}
