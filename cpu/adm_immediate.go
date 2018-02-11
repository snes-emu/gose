package cpu

// IMMEDIATE addressing mode with 8-bit data
func (cpu CPU) admImmediate8(LL uint8) uint8 {
	return LL
}

// IMMEDIATE addressing mode with 16-bit data
func (cpu CPU) admImmediate16(LL uint8, HH uint8) (uint8, uint8) {
	return HH, LL
}
