package cpu

//op88 performs a decrement operation on the Y register, immediate mode
func (cpu *CPU) op88() {
	if cpu.xFlag {
		cpu.setYLRegister(cpu.getYLRegister() - 1)
	} else {
		cpu.Y--
	}
	cpu.cycles += 2
}
