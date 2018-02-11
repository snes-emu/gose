package cpu

// IMMEDIATE addressing mode with 16-bit/8-bit data depending on the m flag
func (cpu CPU) admImmediate() (uint8, uint8) {

	if cpu.mFlag {
		return cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1), 0x00
	}

	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return HH, LL
}
