package cpu

//opE8 performs a increment operation on the X register, immediate mode
func (cpu *CPU) opE8() {
	if cpu.xFlag {
		cpu.setXLRegister(cpu.getXLRegister() + 1)
	} else {
		cpu.X++
	}
	cpu.cycles += 2
}
