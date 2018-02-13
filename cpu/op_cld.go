package cpu

func (cpu *CPU) opD8() {
	cpu.dFlag = false
	cpu.cycles += 2
	cpu.PC++
}
