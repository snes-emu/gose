package cpu

// getEFlag returns the emulation flag
func (cpu *CPU) getEFlag() int16 {
	return (cpu.P >> 8) & 1
}

// getNFlag returns the negative flag
func (cpu *CPU) getNFlag() int16 {
	return (cpu.P >> 7) & 1
}

// getVFlag returns the overflow flag
func (cpu *CPU) getVFlag() int16 {
	return (cpu.P >> 6) & 1
}

// getMFlag returns the accumulator and memory width flag
func (cpu *CPU) getMFlag() int16 {
	return (cpu.P >> 5) & 1
}

// getBFlag returns the break flag
func (cpu *CPU) getBFlag() int16 {
	return (cpu.P >> 4) & 1
}

// getXFlag returns the index register width flag
func (cpu *CPU) getXFlag() int16 {
	return (cpu.P >> 4) & 1
}

// getDFlag returns the decimal mode flag
func (cpu *CPU) getDFlag() int16 {
	return (cpu.P >> 3) & 1
}

// getIFlag returns the interrupt disable flag
func (cpu *CPU) getIFlag() int16 {
	return (cpu.P >> 2) & 1
}

// getZFlag returns the zero flag
func (cpu *CPU) getZFlag() int16 {
	return (cpu.P >> 1) & 1
}

// getCFlag returns the carry flag
func (cpu *CPU) getCFlag() int16 {
	return cpu.P & 1
}
