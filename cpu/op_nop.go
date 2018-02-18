package cpu

func (cpu *CPU) nop() {
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) opEA() {
	cpu.nop()
}

func (cpu *CPU) wdm() {
	cpu.cycles += 2
	cpu.PC += 2
}

func (cpu *CPU) op42() {
	cpu.wdm()
}
