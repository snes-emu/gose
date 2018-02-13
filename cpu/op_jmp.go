package cpu

// jmp jumps to the address specified by the addressing mode
func (cpu *CPU) jmp(addr uint16) {
	cpu.PC = addr
}

// jmpLong jumps to the address specified by the long addressing
func (cpu *CPU) jmpLong(haddr uint8, laddr uint16) {
	cpu.K = haddr
	cpu.PC = laddr
}

func (cpu *CPU) op4C() {
	addr := cpu.admAbsoluteJ()
	cpu.jmp(addr)
	cpu.cycles += 3
	cpu.PC += 3
}

func (cpu *CPU) op5C() {
	haddr, laddr := cpu.admLongJ()
	cpu.jmpLong(haddr, laddr)
	cpu.cycles += 4
	cpu.PC += 4
}

func (cpu *CPU) op6C() {
	addr := cpu.admPAbsoluteJ()
	cpu.jmp(addr)
	cpu.cycles += 5
	cpu.PC += 3
}

func (cpu *CPU) op7C() {
	addr := cpu.admPAbsoluteXJ()
	cpu.jmp(addr)
	cpu.cycles += 6
	cpu.PC += 3
}

func (cpu *CPU) opDC() {
	haddr, laddr := cpu.admBAbsoluteJ()
	cpu.jmpLong(haddr, laddr)
	cpu.cycles += 6
	cpu.PC += 3
}
