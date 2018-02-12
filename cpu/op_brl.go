package cpu

func (cpu *CPU) op82() {
	offset := cpu.admRelative16()
	cpu.cycles += 4
	cpu.PC += offset + 2

}
