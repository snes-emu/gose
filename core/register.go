package core

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

// setARegister sets the lower 8 bits of the accumulator
func (cpu *CPU) setARegister(a uint8) {
	cpu.C = (cpu.C & 0xff00) | uint16(a)
}

// getBRegister returns the upper 8 bits of the accumulator
func (cpu CPU) getBRegister() uint8 {
	return upperBits(cpu.C)
}

// setBRegister sets the upper 8 bits of the accumulator
func (cpu *CPU) setBRegister(b uint8) {
	cpu.C = (cpu.C & 0x00ff) | uint16(b)<<8
}

// getCRegister returns the 16 bits accumulator
func (cpu CPU) getCRegister() uint16 {
	return cpu.C
}

// setCRegister sets the accumulator
func (cpu *CPU) setCRegister(c uint16) {
	cpu.C = c
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

// setDLRegister sets the lower 8 bits of the direct register
func (cpu *CPU) setDLRegister(d uint8) {
	cpu.D = (cpu.D & 0xff00) | uint16(d)
}

// getDHRegister returns the upper 8 bits of the direct register
func (cpu CPU) getDHRegister() uint8 {
	return upperBits(cpu.D)
}

// setDLRegister sets the lower 8 bits of the direct register
func (cpu *CPU) setDHRegister(d uint8) {
	cpu.D = (cpu.D & 0x00ff) | uint16(d)<<8
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

// setSRegister sets the S register
func (cpu *CPU) setSRegister(s uint16) {
	cpu.S = s
}

// getSLRegister returns the lower 8 bits of the stack pointer
func (cpu CPU) getSLRegister() uint8 {
	return lowerBits(cpu.S)
}

// setSLRegister sets the lower 8 bits of the stack pointer
func (cpu *CPU) setSLRegister(s uint8) {
	cpu.S = (cpu.S & 0xff00) | uint16(s)
}

// getSHRegister returns the upper 8 bits of the stack pointer
func (cpu CPU) getSHRegister() uint8 {
	return upperBits(cpu.S)
}

// setSHRegister sets the upper 8 bits of the stack pointer
func (cpu *CPU) setSHRegister(s uint8) {
	cpu.S = (cpu.S & 0x00ff) | uint16(s)<<8
}

// getXRegister returns the X index register
func (cpu CPU) getXRegister() uint16 {
	return cpu.X
}

// setXRegister sets the X register
func (cpu *CPU) setXRegister(x uint16) {
	cpu.X = x
}

// getXLRegister returns the lower 8 bits of the X index register
func (cpu CPU) getXLRegister() uint8 {
	return lowerBits(cpu.X)
}

// setXLRegister sets the lower 8 bits of the X register
func (cpu *CPU) setXLRegister(x uint8) {
	cpu.X = uint16(x)
}

// getXHRegister returns the upper 8 bits of the X index register
func (cpu CPU) getXHRegister() uint8 {
	return upperBits(cpu.X)
}

// setXHRegister sets the upper 8 bits of the X register
func (cpu *CPU) setXHRegister(x uint8) {
	cpu.X = (cpu.X & 0x00ff) | uint16(x)
}

// getYRegister returns the Y index register
func (cpu CPU) getYRegister() uint16 {
	return cpu.Y
}

// setYRegister sets the Y register
func (cpu *CPU) setYRegister(y uint16) {
	cpu.Y = y
}

// getYLRegister returns the lower 8 bits of the Y indey register
func (cpu CPU) getYLRegister() uint8 {
	return lowerBits(cpu.Y)
}

// setYLRegister sets the lower 8 bits of the Y register
func (cpu *CPU) setYLRegister(y uint8) {
	cpu.Y = uint16(y)
}

// getYHRegister returns the upper 8 bits of the Y indey register
func (cpu CPU) getYHRegister() uint8 {
	return upperBits(cpu.Y)
}

// setYHRegister sets the upper 8 bits of the Y register
func (cpu *CPU) setYHRegister(y uint8) {
	cpu.Y = (cpu.Y & 0x00ff) | uint16(y)
}

// setXFlag sets the x flag and take care of the reset of X and Y higher bits
func (cpu *CPU) setXFlag(x bool) {
	cpu.xFlag = x
	if x {
		cpu.setXHRegister(0x00)
		cpu.setYHRegister(0x00)
	}
}

func (cpu *CPU) setEFlag(e bool) {
	cpu.eFlag = e
	// Reset m flag, x flag and SH register for emulation mode
	if cpu.eFlag {
		cpu.mFlag = true
		cpu.setXFlag(true)
		cpu.setSHRegister(0x01)
	}
}
