package cpu

import "github.com/snes-emu/gose/utils"

// sty16 stores the x register in the memory
func (cpu *CPU) sty16(haddr, laddr uint32) {

	dataHi, dataLo := utils.WriteUint16(cpu.getYRegister())

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
