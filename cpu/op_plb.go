package cpu

func (cpu *CPU) opAB() {
	cpu.DBR = cpu.pullStack()
	cpu.nFlag = cpu.getDBRRegister()&0x80 != 0
	cpu.zFlag = cpu.getDBRRegister() == 0
	cpu.cycles += 4
	cpu.PC++
}
