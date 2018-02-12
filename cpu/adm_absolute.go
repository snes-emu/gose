package cpu

import "github.com/snes-emu/gose/utils"

// ABSOLUTE addressing mode to use only for JMP	and JSR instructions
func (cpu CPU) admAbsoluteJ() uint32 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return utils.ReadUint32(cpu.getKRegister(), HH, LL)
}

//ABSOLUTE addressing mode
func (cpu CPU) admAbsolute() (uint8, uint8) {
	haddress, laddress := cpu.admAbsoluteP()
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

//ABSOLUTE addressing mode pointer
func (cpu CPU) admAbsoluteP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return address + 1, address
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteX() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + uint32(cpu.getXRegister()) + 1), cpu.memory.GetByte(address + uint32(cpu.getXRegister()))
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteY() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + uint32(cpu.getYRegister()) + 1), cpu.memory.GetByte(address + uint32(cpu.getYRegister()))
}

// (ABSOLUTE) addressing mode
func (cpu CPU) admPAbsolute() uint32 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(0x00, HH, LL)
	return utils.ReadUint32(cpu.getKRegister(), cpu.memory.GetByte(address+1), cpu.memory.GetByte(address))
}

// [ABSOLUTE] addressing mode
func (cpu CPU) admBAbsolute() uint32 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(0x00, HH, LL)
	return utils.ReadUint32(cpu.memory.GetByte(address+2), cpu.memory.GetByte(address+1), cpu.memory.GetByte(address))
}

// (ABSOLUTE,X) addressing mode
func (cpu CPU) admPAbsoluteX() uint32 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(0x00, HH, LL)
	return utils.ReadUint32(cpu.getKRegister(), cpu.memory.GetByte(address+uint32(cpu.getXRegister())+1), cpu.memory.GetByte(address+uint32(cpu.getXRegister())))
}
