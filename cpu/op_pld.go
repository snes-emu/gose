package cpu

func (cpu *CPU) op2B() {
	cpu.setDHRegister(cpu.pullStack())
	cpu.setDLRegister(cpu.pullStack())
	cpu.nFlag = cpu.getDRegister()&0x80 != 0
	cpu.zFlag = cpu.getDRegister() == 0
	cpu.cycles += 4
	cpu.PC++
}
