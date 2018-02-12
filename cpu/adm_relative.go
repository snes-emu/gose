package cpu

// RELATIVE8 addressing mode
func (cpu CPU) admRelative8() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	switch t := LL >= 0; t {
	case true:
		return uint16(LL)
	default:
		return uint16(LL) - 256
	}
}

// RELATIVE16 addressing mode
func (cpu CPU) admRelative16() uint16 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return uint16(LL) + uint16(HH)<<8
}
