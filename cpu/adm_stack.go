package cpu

import "github.com/snes-emu/gose/utils"

// STACK,S addressing mode
func (cpu CPU) admStackS() (uint8, uint8) {
	laddress, haddress := cpu.admStackSP()
	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

// STACK,S addressing mode pointer
func (cpu CPU) admStackSP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	return laddress, haddress
}

// (STACK,S),Y addressing mode
func (cpu CPU) admPStackSY() (uint8, uint8) {
	laddr, haddr := cpu.admPStackSYP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
}

// (STACK,S),Y addressing mode
func (cpu CPU) admPStackSYP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister()) + uint32(cpu.getYRegister())
	return pointer, pointer + 1
}
