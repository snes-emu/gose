package cpu

import "github.com/snes-emu/gose/utils"

// ABSOLUTE addressing mode to use only for JMP	and JSR instructions
func (cpu CPU) admAbsoluteJ() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return utils.ReadUint16(HH, LL)
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
	haddress, laddress := cpu.admAbsoluteXP()
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// ABSOLUTE,X addressing mode pointer
func (cpu CPU) admAbsoluteXP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return address + uint32(cpu.getXRegister()) + 1, address + uint32(cpu.getXRegister())
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteY() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + uint32(cpu.getYRegister()) + 1), cpu.memory.GetByte(address + uint32(cpu.getYRegister()))
}

// (ABSOLUTE) addressing mode
func (cpu CPU) admPAbsoluteJ() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint16(HH, LL)
	return utils.ReadUint16(cpu.memory.GetByteBank(0x00, address+1), cpu.memory.GetByteBank(0x00, address))
}

// [ABSOLUTE] addressing mode
func (cpu CPU) admBAbsoluteJ() (uint8, uint16) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint16(HH, LL)
	return cpu.memory.GetByteBank(0x00, address+2), utils.ReadUint16(cpu.memory.GetByteBank(0x00, address+1), cpu.memory.GetByteBank(0x00, address))
}

// (ABSOLUTE,X) addressing mode
func (cpu CPU) admPAbsoluteXJ() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.ReadUint16(HH, LL) + cpu.getXRegister()
	return utils.ReadUint16(cpu.memory.GetByteBank(0x00, address+1), cpu.memory.GetByteBank(0x00, address))
}
