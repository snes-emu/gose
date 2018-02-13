package cpu

func (cpu *CPU) op58() {
	cpu.iFlag = false
	cpu.cycles += 2
	cpu.PC++
}
