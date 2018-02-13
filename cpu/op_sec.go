package cpu

func (cpu *CPU) op38() {
	cpu.cFlag = true
	cpu.cycles += 2
	cpu.PC++
}
