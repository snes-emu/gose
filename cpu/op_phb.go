package cpu

func (cpu *CPU) op8B() {
	cpu.pushStack(cpu.getDBRRegister())
	cpu.cycles += 3
	cpu.PC++
}
