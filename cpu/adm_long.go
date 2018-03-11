package cpu

import "github.com/snes-emu/gose/utils"

// LONG addressing mode to use with JMP instruction only
func (cpu CPU) admLongJ() (uint16, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	MM := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+3)
	return utils.JoinUint16(LL, MM), HH
}

// LONG addressing mode
func (cpu CPU) admLong() (uint8, uint8) {
	laddr, haddr := cpu.admLongP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
}

// LONG addressing mode pointer
func (cpu CPU) admLongP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	MM := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+3)
	address := utils.JoinUint32(LL, MM, HH)
	return address, address + 1
}

// LONG,X addressing mode
func (cpu CPU) admLongX() (uint8, uint8) {
	laddr, haddr := cpu.admLongXP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
}

// LONG,X addressing mode pointer
func (cpu CPU) admLongXP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	MM := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+3)
	address := utils.JoinUint32(LL, MM, HH) + uint32(cpu.getXRegister())
	return address, address + 1
}
