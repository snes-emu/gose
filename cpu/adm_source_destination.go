package cpu

// SOURCE,DESTINATION addressing mode
//TODO wrapping behavior is incorrect
func (cpu CPU) admSourceDestination(SS uint8, TT uint8) (uint32, uint32) {
	return uint32(SS)<<16 + uint32(cpu.getXRegister()), uint32(TT)<<16 + uint32(cpu.getYRegister())
}
