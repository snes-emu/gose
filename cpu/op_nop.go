package cpu

func (cpu *CPU) opEA() {
	cpu.cycles += 2
	cpu.PC++
}
