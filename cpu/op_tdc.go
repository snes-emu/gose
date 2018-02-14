package cpu

func (cpu *CPU) op7B() {
	cpu.C = cpu.D
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	cpu.cycles += 2
}
