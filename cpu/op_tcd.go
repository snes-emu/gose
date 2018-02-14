package cpu

func (cpu *CPU) op5B() {
	cpu.D = cpu.C
	// Last bit value
	cpu.nFlag = cpu.D&0x8000 != 0
	cpu.zFlag = cpu.D == 0
	cpu.cycles += 2
}
