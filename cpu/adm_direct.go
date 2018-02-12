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
		address := utils.ReadUint32(0x00, cpu.getDHRegister(), LL)
		return address, 0x00
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return haddress, laddress
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
		address := utils.ReadUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
		return address, 0x00
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister() + 1)
	return haddress, laddress
}

// DIRECT,X addressing mode otherwise
func (cpu CPU) admDirectY() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
		address := utils.ReadUint32(0x00, cpu.getDHRegister(), LL+cpu.getYLRegister())
		return cpu.memory.GetByte(address), 0x00
	}

	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister() + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// (DIRECT) addressing mode when e is 1 and DL is $00
func (cpu CPU) admPDirect8() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	laddress := utils.ReadUint32(0x00, cpu.getDHRegister(), LL)
	haddress := utils.ReadUint32(0x00, cpu.getDHRegister(), LL+1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := utils.ReadUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT) addressing mode otherwise
func (cpu CPU) admPDirect() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// [DIRECT] addressing mode
func (cpu CPU) admBDirect() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := utils.ReadUint32(hh, mm, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT,X) addressing mode otherwise
func (cpu CPU) admPDirectX() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		laddress := utils.ReadUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
		haddress := utils.ReadUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister()+1)
		ll := cpu.memory.GetByte(laddress)
		hh := cpu.memory.GetByte(haddress)
		pointer := utils.ReadUint32(cpu.getDBRRegister(), hh, ll)
		return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
	}

	l := uint16(LL)
	laddress := uint32(cpu.getDRegister() + l + cpu.getXRegister())
	hadress := uint32(cpu.getDRegister() + l + cpu.getXRegister() + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := utils.ReadUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)

}

// (DIRECT),Y addressing mode otherwise
func (cpu CPU) admPDirectY() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)

	if cpu.eFlag && cpu.getDLRegister() == 0x00 {
		LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
		laddress := utils.ReadUint32(0x00, cpu.getDHRegister(), LL)
		haddress := utils.ReadUint32(0x00, cpu.getDHRegister(), LL+1)
		ll := cpu.memory.GetByte(laddress)
		hh := cpu.memory.GetByte(haddress)
		pointer := utils.ReadUint32(cpu.getDBRRegister(), hh, ll) + uint32(cpu.getYRegister())
		return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
	}

	l := uint16(LL)
	laddress := uint32(cpu.getDRegister() + l)
	hadress := uint32(cpu.getDRegister() + l + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := utils.ReadUint32(cpu.getDBRRegister(), hh, ll) + uint32(cpu.getYRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// [DIRECT],Y addressing mode
func (cpu CPU) admBDirectY() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := utils.ReadUint32(hh, mm, ll) + uint32(cpu.getYRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}
