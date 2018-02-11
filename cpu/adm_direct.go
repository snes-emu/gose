package cpu

func (cpu CPU) admDirect8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL)
	return cpu.memory.GetByte(address)
}

func (cpu CPU) admDirect(LL uint8) (uint8, uint8) {
	address := uint32(cpu.getDRegister() + uint16(LL))
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectX8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
	return cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectX(LL uint8) (uint8, uint8) {
	address := uint32(cpu.getDRegister() + uint16(LL) + cpu.getXRegister())
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectY8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getYLRegister())
	return cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectY(LL uint8) (uint8, uint8) {
	laddress := cpu.getDRegister() + uint16(LL) + cpu.getYRegister()
	var uint32 address
	*(*uint16)(unsafe.Pointer(&address)) = laddress
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admPDirect8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL)
	ll := cpu.memory.GetByte(address)
	hh := cpu.memory.GetByte(address + 1)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

func (cpu CPU) admPDirect(LL uint8) (uint8, uint8) {
	address := uint32(cpu.getDRegister() + uint16(LL))
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admBDirect(LL uint8) (uint8, uint8) {
	address := cpu.getDRegister() + uint16(LL)
	ll := cpu.memory.GetByte(uint32(address))
	mm := cpu.memory.GetByte(uint32(address + 1))
	hh := cpu.memory.GetByte(uint32(address + 2))
	pointer := readUint32(hh, mm, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

func (cpu CPU) admPDirectX8(LL uint8) (uint8, uint8) {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
	ll := cpu.memory.GetByte(address)
	hh := cpu.memory.GetByte(address + 1)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

func (cpu CPU) admPDirectX(LL uint8) (uint8, uint8) {
	address := uint32(cpu.getDRegister() + uint16(LL) + cpu.getXRegister())
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}
