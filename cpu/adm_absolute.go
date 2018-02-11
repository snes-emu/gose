package cpu

import "github.com/snes-emu/gose/utils"

// ABSOLUTE addressing mode to use only for JMP	and JSR instructions
func (cpu CPU) admAbsoluteJ(HH uint8, LL uint8) uint32 {
	return utils.ReadUint32(cpu.getKRegister(), HH, LL)
}

//ABSOLUTE addressing mode
func (cpu CPU) admAbsolute(HH uint8, LL uint8) (uint8, uint8) {
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteX(HH uint8, LL uint8) (uint8, uint8) {
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + uint32(cpu.getXRegister()) + 1), cpu.memory.GetByte(address + uint32(cpu.getXRegister()))
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteY(HH uint8, LL uint8) (uint8, uint8) {
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + uint32(cpu.getYRegister()) + 1), cpu.memory.GetByte(address + uint32(cpu.getYRegister()))
}

// (ABSOLUTE) addressing mode
func (cpu CPU) admPAbsolute(HH uint8, LL uint8) uint32 {
	address := utils.ReadUint32(0x00, HH, LL)
	return utils.ReadUint32(cpu.getKRegister(), cpu.memory.GetByte(address+1), cpu.memory.GetByte(address))
}

// [ABSOLUTE] addressing mode
func (cpu CPU) admBAbsolute(HH uint8, LL uint8) uint32 {
	address := utils.ReadUint32(0x00, HH, LL)
	return utils.ReadUint32(cpu.memory.GetByte(address+2), cpu.memory.GetByte(address+1), cpu.memory.GetByte(address))
}

// (ABSOLUTE,X) addressing mode
func (cpu CPU) admPAbsoluteX(HH uint8, LL uint8) uint32 {
	address := utils.ReadUint32(0x00, HH, LL)
	return utils.ReadUint32(cpu.getKRegister(), cpu.memory.GetByte(address+uint32(cpu.getXRegister())+1), cpu.memory.GetByte(address+uint32(cpu.getXRegister())))
}
