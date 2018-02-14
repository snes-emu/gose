package cpu

func (cpu *CPU) op0B() {
	cpu.pushStack(cpu.getDHRegister())
	cpu.pushStack(cpu.getDLRegister())
	cpu.cycles += 4
	cpu.PC++
}
