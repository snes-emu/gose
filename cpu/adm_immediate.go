package cpu

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
	return HH, LL
}
