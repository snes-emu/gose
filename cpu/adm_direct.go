package cpu

import "github.com/snes-emu/gose/utils"

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
