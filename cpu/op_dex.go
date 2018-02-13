package cpu

//opCA performs a decrement operation on the X register
func (cpu *CPU) opCA() {
	if cpu.xFlag {
		cpu.setXLRegister(cpu.getXLRegister() - 1)
	} else {
		cpu.X--
	}
	cpu.cycles += 2
}
