package cpu

import "github.com/snes-emu/gose/utils"

// LONG addressing mode to use with JMP instruction only
func (cpu CPU) admLongJ() (uint8, uint16) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	MM := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+3)
	return HH, utils.ReadUint16(MM, LL)
}

// LONG addressing mode
func (cpu CPU) admLong() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	MM := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+3)
	address := utils.ReadUint32(HH, MM, LL)
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

// LONG,X addressing mode
func (cpu CPU) admLongX() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	MM := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+3)
	address := utils.ReadUint32(HH, MM, LL) + uint32(cpu.getXRegister())
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}
