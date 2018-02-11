package cpu

// ACCUMULATOR addressing mode
func (cpu CPU) admAccumulator() (uint8, uint8) {
	return cpu.getBRegister(), cpu.getARegister()
}
