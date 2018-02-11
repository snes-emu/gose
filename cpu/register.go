package cpu

// lowerBits returns the lower bits of a uint16 number
func lowerBits(x uint16) uint8 {
	return uint8(x & 0xff)
}

// upperBits returns the lower bits of a uint16 number
func upperBits(x uint16) uint8 {
	return uint8(x >> 8)
}

// getARegister returns the lower 8 bits of the accumulator
func (cpu CPU) getARegister() uint8 {
	return lowerBits(cpu.C)
}

// getBRegister returns the upper 8 bits of the accumulator
func (cpu CPU) getBRegister() uint8 {
	return upperBits(cpu.C)
}

// getCRegister returns the 16 bits accumulator
func (cpu CPU) getCRegister() uint16 {
	return cpu.C
}

// getDBRRegister returns the data bank register
func (cpu CPU) getDBRRegister() uint8 {
	return cpu.DBR
}

// getDRegister returns the D register
func (cpu CPU) getDRegister() uint16 {
	return cpu.D
}

// getDLRegister returns the lower 8 bits of the direct register
func (cpu CPU) getDLRegister() uint8 {
	return lowerBits(cpu.D)
}

// getDHRegister returns the upper 8 bits of the direct register
func (cpu CPU) getDHRegister() uint8 {
	return upperBits(cpu.D)
}

// getKRegister returns the program bank register
func (cpu CPU) getKRegister() uint8 {
	return cpu.K
}

// getPCRegister returns the program counter
func (cpu CPU) getPCRegister() uint16 {
	return cpu.PC
}

// getPCLRegister returns the lower 8 bits of the program counter
func (cpu CPU) getPCLRegister() uint8 {
	return lowerBits(cpu.PC)
}

// getPCHRegister returns the lower 8 bits of the program counter
func (cpu CPU) getPCHRegister() uint8 {
	return upperBits(cpu.PC)
}

// getSRegister returns the stack pointer
func (cpu CPU) getSRegister() uint16 {
	return cpu.S
}

// getSLRegister returns the lower 8 bits of the stack pointer
func (cpu CPU) getSLRegister() uint8 {
	return lowerBits(cpu.S)
}

// getSHRegister returns the upper 8 bits of the stack pointer
func (cpu CPU) getSHRegister() uint8 {
	return upperBits(cpu.S)
}

// getXRegister returns the X index register
func (cpu CPU) getXRegister() uint16 {
	return cpu.X
}

// getXLRegister returns the lower 8 bits of the X index register
func (cpu CPU) getXLRegister() uint8 {
	return lowerBits(cpu.X)
}

// getXHRegister returns the upper 8 bits of the X index register
func (cpu CPU) getXHRegister() uint8 {
	return upperBits(cpu.X)
}

// getYRegister returns the Y index register
func (cpu CPU) getYRegister() uint16 {
	return cpu.Y
}

// getYLRegister returns the lower 8 bits of the Y index register
func (cpu CPU) getYLRegister() uint8 {
	return lowerBits(cpu.Y)
}

// getYHRegister returns the upper 8 bits of the Y index register
func (cpu CPU) getYHRegister() uint8 {
	return upperBits(cpu.Y)
}
