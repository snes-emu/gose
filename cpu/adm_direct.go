package cpu

// DIRECT addressing mode only when old instruction, e is 1 and DL is $00
func (cpu CPU) admDirect8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL)
	return cpu.memory.GetByte(address)
}

// DIRECT addressing mode otherwise
func (cpu CPU) admDirect(LL uint8) (uint8, uint8) {
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// DIRECT,X addressing mode when e is 1 and DL is $00
func (cpu CPU) admDirectX8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
	return cpu.memory.GetByte(address)
}

// DIRECT,X addressing mode otherwise
func (cpu CPU) admDirectX(LL uint8) (uint8, uint8) {
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getXRegister() + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// DIRECT,Y addressing mode when e is 1 and DL is $00
func (cpu CPU) admDirectY8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getYLRegister())
	return cpu.memory.GetByte(address)
}

// DIRECT,X addressing mode otherwise
func (cpu CPU) admDirectY(LL uint8) (uint8, uint8) {
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister())
	haddress := uint32(cpu.getDRegister() + ll + cpu.getYRegister() + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// (DIRECT) addressing mode when e is 1 and DL is $00
func (cpu CPU) admPDirect8(LL uint8) (uint8, uint8) {
	laddress := readUint32(0x00, cpu.getDHRegister(), LL)
	haddress := readUint32(0x00, cpu.getDHRegister(), LL+1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT) addressing mode otherwise
func (cpu CPU) admPDirect(LL uint8) (uint8, uint8) {
	ll := uint16(LL)
	laddress := uint32(cpu.getDRegister() + ll)
	haddress := uint32(cpu.getDRegister() + ll + 1)
	return cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)
}

// [DIRECT] addressing mode
func (cpu CPU) admBDirect(LL uint8) (uint8, uint8) {
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := readUint32(hh, mm, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT,X) addressing mode when e is 1 and DL is $00
func (cpu CPU) admPDirectX8(LL uint8) (uint8, uint8) {
	laddress := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
	haddress := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister()+1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT,X) addressing mode otherwise
func (cpu CPU) admPDirectX(LL uint8) (uint8, uint8) {
	l := uint16(LL)
	laddress := uint32(cpu.getDRegister() + l + cpu.getXRegister())
	hadress := uint32(cpu.getDRegister() + l + cpu.getXRegister() + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT),Y addressing mode when e is 1 and DL is $00
func (cpu CPU) admPDirectY8(LL uint8) (uint8, uint8) {
	laddress := readUint32(0x00, cpu.getDHRegister(), LL)
	haddress := readUint32(0x00, cpu.getDHRegister(), LL+1)
	ll := cpu.memory.GetByte(laddress)
	hh := cpu.memory.GetByte(haddress)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll) + uint32(cpu.getYRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// (DIRECT),Y addressing mode otherwise
func (cpu CPU) admPDirectY(LL uint8) (uint8, uint8) {
	l := uint16(LL)
	laddress := uint32(cpu.getDRegister() + l)
	hadress := uint32(cpu.getDRegister() + l + 1)
	hh := cpu.memory.GetByte(hadress)
	ll := cpu.memory.GetByte(laddress)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll) + uint32(cpu.getYRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

// [DIRECT],Y addressing mode
func (cpu CPU) admBDirectY(LL uint8) (uint8, uint8) {
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := readUint32(hh, mm, ll) + uint32(cpu.getYRegister())
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}
