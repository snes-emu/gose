package cpu

import "github.com/snes-emu/gose/utils"

// STACK,S addressing mode
func (cpu CPU) admStackS() (uint8, uint8) {
	haddress, laddress := cpu.admStackSP()
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// STACK,S addressing mode pointer
func (cpu CPU) admStackSP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	return haddress, laddress
}

// (STACK,S),Y addressing mode
func (cpu CPU) admPStackSY() (uint8, uint8) {
	haddr, laddr := cpu.admPStackSYP()
	return cpu.memory.GetByte(haddr), cpu.memory.GetByte(laddr)
}

// (STACK,S),Y addressing mode
func (cpu CPU) admPStackSYP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := utils.JoinUint32(cpu.getDBRRegister(), hh, ll) + uint32(cpu.getYRegister())
	return pointer + 1, pointer
}
