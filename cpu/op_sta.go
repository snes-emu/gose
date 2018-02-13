package cpu

import "github.com/snes-emu/gose/utils"

// sta16 stores the accumulator in the memory
func (cpu *CPU) sta16(haddr, laddr uint32) {

	dataHi, dataLo := utils.WriteUint16(cpu.getCRegister())

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
