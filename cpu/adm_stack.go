package cpu

import "github.com/snes-emu/gose/utils"

// STACK,S addressing mode
func (cpu CPU) admStackS() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// (STACK,S),Y addressing mode
func (cpu CPU) admPStackSY() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := utils.ReadUint32(cpu.getDBRRegister(), hh, ll) + uint32(cpu.getYRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}
