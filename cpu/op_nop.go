package cpu

func (cpu *CPU) opEA() {
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op42() {
	cpu.cycles += 2
	cpu.PC += 2
}
