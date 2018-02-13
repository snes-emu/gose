package cpu

func (cpu *CPU) op18() {
	cpu.cFlag = false
	cpu.cycles += 2
	cpu.PC++
}
