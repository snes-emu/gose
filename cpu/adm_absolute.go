package cpu

import "github.com/snes-emu/gose/utils"

// ABSOLUTE addressing mode to use only for JMP	and JSR instructions
func (cpu CPU) admAbsoluteJ() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return utils.JoinUint16(LL, HH)
}

//ABSOLUTE addressing mode
func (cpu CPU) admAbsolute() (uint8, uint8) {
	laddress, haddress := cpu.admAbsoluteP()
	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

//ABSOLUTE addressing mode pointer
func (cpu CPU) admAbsoluteP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.JoinUint32(LL, HH, cpu.getDBRRegister())
	return address, address + 1
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteX() (uint8, uint8) {
	laddress, haddress := cpu.admAbsoluteXP()
	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

// ABSOLUTE,X addressing mode pointer
func (cpu CPU) admAbsoluteXP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.JoinUint32(LL, HH, cpu.getDBRRegister())
	cpu.pFlag = uint16(LL)+0x00FF+cpu.getXRegister()+1 > 0xFF
	return address + uint32(cpu.getXRegister()), address + uint32(cpu.getXRegister()) + 1
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteY() (uint8, uint8) {
	laddr, haddr := cpu.admAbsoluteYP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
}

// ABSOLUTE,X addressing mode pointer
func (cpu CPU) admAbsoluteYP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.JoinUint32(LL, HH, cpu.getDBRRegister())
	cpu.pFlag = uint16(LL)+0x00FF+cpu.getYRegister()+1 > 0xFF
	return address + uint32(cpu.getYRegister()), address + uint32(cpu.getYRegister()) + 1
}

// (ABSOLUTE) addressing mode
func (cpu CPU) admPAbsoluteJ() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.JoinUint16(LL, HH)
	return utils.JoinUint16(cpu.memory.GetByteBank(0x00, address), cpu.memory.GetByteBank(0x00, address+1))
}

// [ABSOLUTE] addressing mode
func (cpu CPU) admBAbsoluteJ() (uint8, uint16) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.JoinUint16(LL, HH)
	return cpu.memory.GetByteBank(0x00, address+2), utils.JoinUint16(cpu.memory.GetByteBank(0x00, address), cpu.memory.GetByteBank(0x00, address+1))
}

// (ABSOLUTE,X) addressing mode
func (cpu CPU) admPAbsoluteXJ() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	address := utils.JoinUint16(LL, HH) + cpu.getXRegister()
	return utils.JoinUint16(cpu.memory.GetByteBank(0x00, address), cpu.memory.GetByteBank(0x00, address+1))
}
