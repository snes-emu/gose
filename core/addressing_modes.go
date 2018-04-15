package core

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

// ACCUMULATOR addressing mode
func (cpu CPU) admAccumulator() (uint8, uint8) {
	return cpu.getARegister(), cpu.getBRegister()
}

// DIRECT addressing mode otherwise
func (cpu CPU) admDirect() (uint8, uint8) {
	laddress, haddress := cpu.admDirectP()

	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

// DIRECT addressing mode pointer
func (cpu CPU) admDirectP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		address := utils.JoinUint32(LL, cpu.getDHRegister(), 0x00)
		return 0x00, address
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return laddress, haddress
}

// DIRECT addressing mode for "new" intructions (only use by PEI)
func (cpu CPU) admDirectNew() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

// DIRECT,X addressing mode otherwise
func (cpu CPU) admDirectX() (uint8, uint8) {
	laddress, haddress := cpu.admDirectXP()
	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

// DIRECT,X addressing mode pointer
func (cpu CPU) admDirectXP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		address := utils.JoinUint32(LL+cpu.getXLRegister(), cpu.getDHRegister(), 0x00)
		return address, 0x00
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister() + 1)
	return laddress, haddress
}

// DIRECT,X addressing mode otherwise pointer
func (cpu CPU) admDirectY() (uint8, uint8) {
	laddress, haddress := cpu.admDirectYP()
	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

// DIRECT,X addressing mode otherwise
func (cpu CPU) admDirectYP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
		address := utils.JoinUint32(LL+cpu.getYLRegister(), cpu.getDHRegister(), 0x00)
		return address, 0x00
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister() + 1)
	return laddress, haddress
}

// (DIRECT) addressing mode when e is 1 and DL is $00
func (cpu CPU) admPDirect8() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := utils.JoinUint32(LL, cpu.getDHRegister(), 0x00)
	haddress := utils.JoinUint32(LL+1, cpu.getDHRegister(), 0x00)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister())
	return cpu.memory.GetByte(pointer), cpu.memory.GetByte(pointer + 1)
}

// (DIRECT) addressing mode otherwise
func (cpu CPU) admPDirect() (uint8, uint8) {
	laddress, haddress := cpu.admPDirectP()
	return cpu.memory.GetByte(laddress), cpu.memory.GetByte(haddress)
}

// (DIRECT) addressing mode pointer
func (cpu CPU) admPDirectP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return laddress, haddress
}

// [DIRECT] addressing mode
func (cpu CPU) admBDirect() (uint8, uint8) {
	laddr, haddr := cpu.admBDirectP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
}

// [DIRECT] addressing mode pointer
func (cpu CPU) admBDirectP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := utils.JoinUint32(ll, mm, hh)
	return pointer, pointer + 1
}

// (DIRECT,X) addressing mode otherwise
func (cpu CPU) admPDirectX() (uint8, uint8) {
	laddr, haddr := cpu.admPDirectXP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)

}

// (DIRECT,X) addressing mode pointer
func (cpu CPU) admPDirectXP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		laddress := utils.JoinUint32(LL+cpu.getXLRegister(), cpu.getDHRegister(), 0x00)
		haddress := utils.JoinUint32(LL+cpu.getXLRegister()+1, cpu.getDHRegister(), 0x00)
		ll := cpu.memory.GetByte(laddress)
		hh := cpu.memory.GetByte(haddress)
		pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister())
		return pointer, pointer + 1
	}

	l := uint16(LL)
	laddress := uint32(cpu.getDRegister() + l + cpu.getXRegister())
	hadress := uint32(cpu.getDRegister() + l + cpu.getXRegister() + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister())
	return pointer, pointer + 1

}

// (DIRECT),Y addressing mode otherwise
func (cpu CPU) admPDirectY() (uint8, uint8) {
	laddr, haddr := cpu.admPDirectYP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
}

// (DIRECT),Y addressing mode pointer
func (cpu CPU) admPDirectYP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
		laddress := utils.JoinUint32(LL, cpu.getDHRegister(), 0x00)
		haddress := utils.JoinUint32(LL+1, cpu.getDHRegister(), 0x00)
		ll := cpu.memory.GetByte(laddress)
		hh := cpu.memory.GetByte(haddress)
		pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister()) + uint32(cpu.getYRegister())
		return pointer, pointer + 1
	}

	l := uint16(LL)
	cpu.pFlag = cpu.getDRegister()&0x00FF+l+1 > 0xFF
	laddress := uint32(cpu.getDRegister() + l)
	hadress := uint32(cpu.getDRegister() + l + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister()) + uint32(cpu.getYRegister())
	return pointer, pointer + 1
}

// [DIRECT],Y addressing mode
func (cpu CPU) admBDirectY() (uint8, uint8) {
	laddr, haddr := cpu.admBDirectYP()
	return cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
}

// [DIRECT],Y addressing mode pointer
func (cpu CPU) admBDirectYP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := utils.JoinUint32(ll, mm, hh) + uint32(cpu.getYRegister())
	return pointer, pointer + 1
}

// IMMEDIATE addressing mode with 16-bit/8-bit data depending on the m flag
func (cpu CPU) admImmediateM() (uint8, uint8) {

	if cpu.mFlag {
		return cpu.admImmediate8()
	}

	return cpu.admImmediate16()
}

// IMMEDIATE addressing mode with 16-bit/8-bit data depending on the x flag
func (cpu CPU) admImmediateX() (uint8, uint8) {

	if cpu.xFlag {
		return cpu.admImmediate8()
	}

	return cpu.admImmediate16()
}

// IMMEDIATE addressing mode with 8-bit data
func (cpu CPU) admImmediate8() (uint8, uint8) {
	return cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1), 0x00
}

// IMMEDIATE addressing mode with 16-bite data
func (cpu CPU) admImmediate16() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return LL, HH
}

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

// RELATIVE8 addressing mode
func (cpu CPU) admRelative8() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	if LL < 80 {
		return uint16(LL)
	}
	return uint16(LL) - 256

}

// RELATIVE16 addressing mode
func (cpu CPU) admRelative16() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return uint16(LL) + uint16(HH)<<8
}

// SOURCE,DESTINATION addressing mode
func (cpu CPU) admSourceDestination() (uint8, uint16, uint8, uint16) {
	SBank := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	DBank := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	var SAddress, DAddress uint16
	if cpu.xFlag {
		SAddress = uint16(cpu.getXLRegister())
		DAddress = uint16(cpu.getYLRegister())
	} else {
		SAddress = cpu.getXRegister()
		DAddress = cpu.getYRegister()
	}
	return SBank, SAddress, DBank, DAddress
}

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
