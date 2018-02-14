package cpu

func (cpu *CPU) opFB() {
	temp := cpu.eFlag
	cpu.eFlag = cpu.cFlag
	cpu.cFlag = temp
	cpu.cycles += 2
	cpu.PC++
}
