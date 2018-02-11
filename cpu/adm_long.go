package cpu

import "github.com/snes-emu/gose/utils"

// LONG addressing mode to use with JMP instruction only
func (cpu CPU) admLongJ(LL uint8, MM uint8, HH uint8) uint32 {
	return utils.ReadUint32(HH, MM, LL)
}

// LONG addressing mode
func (cpu CPU) admLong(LL uint8, MM uint8, HH uint8) (uint8, uint8) {
	address := utils.ReadUint32(HH, MM, LL)
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

// LONG,X addressing mode
func (cpu CPU) admLongX(LL uint8, MM uint8, HH uint8) (uint8, uint8) {
	address := utils.ReadUint32(HH, MM, LL) + uint32(cpu.getXRegister())
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}
