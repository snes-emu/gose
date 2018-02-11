package cpu

func (cpu CPU) admLongJ(LL uint8, MM uint8, HH uint8) uint32 {
	return readUint32(HH, MM, LL)
}

func (cpu CPU) admLong(LL uint8, MM uint8, HH uint8) (uint8, uint8) {
	address := readUint32(HH, MM, LL)
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admLongX(LL uint8, MM uint8, HH uint8) (uint8, uint8) {
	address := readUint32(HH, MM, LL) + uint32(cpu.getXRegister())
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}
