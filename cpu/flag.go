package cpu

// getEFlag returns the emulation flag
func (cpu CPU) getEFlag() uint16 {
	return (cpu.P >> 8) & 1
}

// getNFlag returns the negative flag
func (cpu CPU) getNFlag() uint16 {
	return (cpu.P >> 7) & 1
}

// getVFlag returns the overflow flag
func (cpu CPU) getVFlag() uint16 {
	return (cpu.P >> 6) & 1
}

// getMFlag returns the accumulator and memory width flag
func (cpu CPU) getMFlag() uint16 {
	return (cpu.P >> 5) & 1
}

// getBFlag returns the break flag
func (cpu CPU) getBFlag() uint16 {
	return (cpu.P >> 4) & 1
}

// getXFlag returns the index register width flag
func (cpu CPU) getXFlag() uint16 {
	return (cpu.P >> 4) & 1
}

// getDFlag returns the decimal mode flag
func (cpu CPU) getDFlag() uint16 {
	return (cpu.P >> 3) & 1
}

// getIFlag returns the interrupt disable flag
func (cpu CPU) getIFlag() uint16 {
	return (cpu.P >> 2) & 1
}

// getZFlag returns the zero flag
func (cpu CPU) getZFlag() uint16 {
	return (cpu.P >> 1) & 1
}

// getCFlag returns the carry flag
func (cpu CPU) getCFlag() uint16 {
	return cpu.P & 1
}
