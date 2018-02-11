package cpu

// IMMEDIATE addressing mode with 8-bit data
func (cpu CPU) admImmediate8() uint8 {
	return cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
}

// IMMEDIATE addressing mode with 16-bit data
func (cpu CPU) admImmediate16() (uint8, uint8) {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return HH, LL
}
