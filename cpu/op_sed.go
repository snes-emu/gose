package cpu

func (cpu *CPU) opF8() {
	cpu.dFlag = true
	cpu.cycles += 2
	cpu.PC++
}
