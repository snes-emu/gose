package cpu

func (cpu *CPU) opB8() {
	cpu.vFlag = false
	cpu.cycles += 2
	cpu.PC++
}
