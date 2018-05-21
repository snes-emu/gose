package core

import "github.com/snes-emu/gose/bit"

// getARegister returns the lower 8 bits of the accumulator
func (cpu *CPU) getARegister() uint8 {
	return bit.LowByte(cpu.C)
}

// setARegister sets the lower 8 bits of the accumulator
func (cpu *CPU) setARegister(a uint8) {
	cpu.C = bit.SetLowByte(cpu.C, a)
}

// getBRegister returns the upper 8 bits of the accumulator
func (cpu *CPU) getBRegister() uint8 {
	return bit.HighByte(cpu.C)
}

// setBRegister sets the upper 8 bits of the accumulator
func (cpu *CPU) setBRegister(b uint8) {
	cpu.C = bit.SetHighByte(cpu.C, b)
}

// getCRegister returns the 16 bits accumulator
func (cpu *CPU) getCRegister() uint16 {
	return cpu.C
}

// setCRegister sets the accumulator
func (cpu *CPU) setCRegister(c uint16) {
	cpu.C = c
}

// getDBRRegister returns the data bank register
func (cpu *CPU) getDBRRegister() uint8 {
	return cpu.DBR
}

// getDRegister returns the D register
func (cpu *CPU) getDRegister() uint16 {
	return cpu.D
}

// getDLRegister returns the lower 8 bits of the direct register
func (cpu *CPU) getDLRegister() uint8 {
	return bit.LowByte(cpu.D)
}

// setDLRegister sets the lower 8 bits of the direct register
func (cpu *CPU) setDLRegister(d uint8) {
	cpu.D = bit.SetLowByte(cpu.D, d)
}

// getDHRegister returns the upper 8 bits of the direct register
func (cpu *CPU) getDHRegister() uint8 {
	return bit.HighByte(cpu.D)
}

// setDHRegister sets the upper 8 bits of the direct register
func (cpu *CPU) setDHRegister(d uint8) {
	cpu.D = bit.SetHighByte(cpu.D, d)
}

// getKRegister returns the program bank register
func (cpu *CPU) getKRegister() uint8 {
	return cpu.K
}

// getPCRegister returns the program counter
func (cpu *CPU) getPCRegister() uint16 {
	return cpu.PC
}

// getPCLRegister returns the lower 8 bits of the program counter
func (cpu *CPU) getPCLRegister() uint8 {
	return bit.LowByte(cpu.PC)
}

// getPCHRegister returns the lower 8 bits of the program counter
func (cpu *CPU) getPCHRegister() uint8 {
	return bit.HighByte(cpu.PC)
}

// getSRegister returns the stack pointer
func (cpu *CPU) getSRegister() uint16 {
	return cpu.S
}

// setSRegister sets the S register
func (cpu *CPU) setSRegister(s uint16) {
	cpu.S = s
}

// getSLRegister returns the lower 8 bits of the stack pointer
func (cpu *CPU) getSLRegister() uint8 {
	return bit.LowByte(cpu.S)
}

// setSLRegister sets the lower 8 bits of the stack pointer
func (cpu *CPU) setSLRegister(s uint8) {
	cpu.S = bit.SetLowByte(cpu.S, s)
}

// getSHRegister returns the upper 8 bits of the stack pointer
func (cpu *CPU) getSHRegister() uint8 {
	return bit.HighByte(cpu.S)
}

// setSHRegister sets the upper 8 bits of the stack pointer
func (cpu *CPU) setSHRegister(s uint8) {
	cpu.S = bit.SetHighByte(cpu.S, s)
}

// getXRegister returns the X index register
func (cpu *CPU) getXRegister() uint16 {
	return cpu.X
}

// setXRegister sets the X register
func (cpu *CPU) setXRegister(x uint16) {
	cpu.X = x
}

// getXLRegister returns the lower 8 bits of the X index register
func (cpu *CPU) getXLRegister() uint8 {
	return bit.LowByte(cpu.X)
}

// setXLRegister sets the lower 8 bits of the X register
func (cpu *CPU) setXLRegister(x uint8) {
	cpu.X = uint16(x)
}

// getXHRegister returns the upper 8 bits of the X index register
func (cpu *CPU) getXHRegister() uint8 {
	return bit.HighByte(cpu.X)
}

// setXHRegister sets the upper 8 bits of the X register
func (cpu *CPU) setXHRegister(x uint8) {
	cpu.X = bit.SetHighByte(cpu.X, x)
}

// getYRegister returns the Y index register
func (cpu *CPU) getYRegister() uint16 {
	return cpu.Y
}

// setYRegister sets the Y register
func (cpu *CPU) setYRegister(y uint16) {
	cpu.Y = y
}

// getYLRegister returns the lower 8 bits of the Y indey register
func (cpu *CPU) getYLRegister() uint8 {
	return bit.LowByte(cpu.Y)
}

// setYLRegister sets the lower 8 bits of the Y register
func (cpu *CPU) setYLRegister(y uint8) {
	cpu.Y = uint16(y)
}

// getYHRegister returns the upper 8 bits of the Y indey register
func (cpu *CPU) getYHRegister() uint8 {
	return bit.HighByte(cpu.Y)
}

// setYHRegister sets the upper 8 bits of the Y register
func (cpu *CPU) setYHRegister(y uint8) {
	cpu.Y = bit.SetHighByte(cpu.Y, y)
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
