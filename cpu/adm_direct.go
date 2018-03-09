package cpu

import "github.com/snes-emu/gose/utils"

// DIRECT addressing mode otherwise
func (cpu CPU) admDirect() (uint8, uint8) {
	haddress, laddress := cpu.admDirectP()

	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
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
	return haddress, laddress
}

// DIRECT addressing mode for "new" intructions (only use by PEI)
func (cpu CPU) admDirectNew() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// DIRECT,X addressing mode otherwise
func (cpu CPU) admDirectX() (uint8, uint8) {
	haddress, laddress := cpu.admDirectXP()
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// DIRECT,X addressing mode pointer
func (cpu CPU) admDirectXP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		address := utils.JoinUint32(LL+cpu.getXLRegister(), cpu.getDHRegister(), 0x00)
		return 0x00, address
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister() + 1)
	return haddress, laddress
}

// DIRECT,X addressing mode otherwise pointer
func (cpu CPU) admDirectY() (uint8, uint8) {
	haddress, laddress := cpu.admDirectYP()
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// DIRECT,X addressing mode otherwise
func (cpu CPU) admDirectYP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
		address := utils.JoinUint32(LL+cpu.getYLRegister(), cpu.getDHRegister(), 0x00)
		return 0x00, address
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister() + 1)
	return haddress, laddress
}

// (DIRECT) addressing mode when e is 1 and DL is $00
func (cpu CPU) admPDirect8() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := utils.JoinUint32(LL, cpu.getDHRegister(), 0x00)
	haddress := utils.JoinUint32(LL+1, cpu.getDHRegister(), 0x00)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT) addressing mode otherwise
func (cpu CPU) admPDirect() (uint8, uint8) {
	haddress, laddress := cpu.admPDirectP()
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// (DIRECT) addressing mode pointer
func (cpu CPU) admPDirectP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return haddress, laddress
}

// [DIRECT] addressing mode
func (cpu CPU) admBDirect() (uint8, uint8) {
	haddr, laddr := cpu.admBDirectP()
	return cpu.memory.GetByte(haddr), cpu.memory.GetByte(laddr)
}

// [DIRECT] addressing mode pointer
func (cpu CPU) admBDirectP() (uint32, uint32) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := utils.JoinUint32(ll, mm, hh)
	return pointer + 1, pointer
}

// (DIRECT,X) addressing mode otherwise
func (cpu CPU) admPDirectX() (uint8, uint8) {
	haddr, laddr := cpu.admPDirectXP()
	return cpu.memory.GetByte(haddr), cpu.memory.GetByte(laddr)

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
		return pointer + 1, pointer
	}

	l := uint16(LL)
	laddress := uint32(cpu.getDRegister() + l + cpu.getXRegister())
	hadress := uint32(cpu.getDRegister() + l + cpu.getXRegister() + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister())
	return pointer + 1, pointer

}

// (DIRECT),Y addressing mode otherwise
func (cpu CPU) admPDirectY() (uint8, uint8) {
	haddr, laddr := cpu.admPDirectYP()
	return cpu.memory.GetByte(haddr), cpu.memory.GetByte(laddr)
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
		return pointer + 1, pointer
	}

	l := uint16(LL)
	cpu.pFlag = cpu.getDRegister()&0x00FF+l+1 > 0xFF
	laddress := uint32(cpu.getDRegister() + l)
	hadress := uint32(cpu.getDRegister() + l + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := utils.JoinUint32(ll, hh, cpu.getDBRRegister()) + uint32(cpu.getYRegister())
	return pointer + 1, pointer
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
