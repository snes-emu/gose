package cpu

func (cpu *CPU) op3B() {
	cpu.C = cpu.S
	// Last bit value
	cpu.nFlag = cpu.S&0x8000 != 0
	cpu.zFlag = cpu.S == 0
	cpu.cycles += 2
}
