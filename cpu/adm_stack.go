package cpu

import "github.com/snes-emu/gose/utils"

// STACK,S addressing mode
func (cpu CPU) admStackS(LL uint8) (uint8, uint8) {
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// (STACK,S),Y addressing mode
func (cpu CPU) admStackSY(LL uint8) (uint8, uint8) {
	laddress := uint32(cpu.getSRegister() + uint16(LL))
	haddress := uint32(cpu.getSRegister() + uint16(LL) + 1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := utils.ReadUint32(cpu.getDBRRegister(), hh, ll) + uint32(cpu.getYRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}
