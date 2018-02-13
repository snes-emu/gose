package cpu

func (cpu *CPU) op78() {
	cpu.iFlag = true
	cpu.cycles += 2
	cpu.PC++
}
