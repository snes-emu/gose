package cpu

// RELATIVE8 addressing mode
func (cpu CPU) admRelative8() uint32 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	switch t := LL >= 0; t {
	case true:
		return uint32(cpu.getKRegister())<<16 + uint32(cpu.getPCRegister()+2+uint16(LL))
	default:
		return uint32(cpu.getKRegister())<<16 + uint32(cpu.getPCRegister()-254+uint16(LL))
	}
}

// RELATIVE16 addressing mode
func (cpu CPU) admRelative16() uint32 {
	LL := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	HH := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	return uint32(cpu.getKRegister())<<16 + uint32(cpu.getPCRegister()+3+uint16(LL)+uint16(HH)<<8)
}
