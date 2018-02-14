package cpu

func (cpu *CPU) op4B() {
	cpu.pushStack(cpu.getKRegister())
	cpu.cycles += 3
	cpu.PC++
}
