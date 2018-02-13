package cpu

//opC8 performs a increment operation on the Y register, immediate mode
func (cpu *CPU) opC8() {
	if cpu.xFlag {
		cpu.setYLRegister(cpu.getYLRegister() + 1)
	} else {
		cpu.Y++
	}
	cpu.cycles += 2
}
