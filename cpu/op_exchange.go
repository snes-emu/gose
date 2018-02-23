package cpu

func (cpu *CPU) xba() {
	temp := cpu.getBRegister()
	cpu.setBRegister(cpu.getARegister())
	cpu.setARegister(temp)
	cpu.nFlag = temp&0x80 != 0
	cpu.zFlag = temp == 0
	cpu.cycles += 3
	cpu.PC++
}

func (cpu *CPU) opEB() {
	cpu.xba()
}

func (cpu *CPU) xce() {
	temp := cpu.eFlag
	cpu.setEFlag(cpu.cFlag)
	cpu.cFlag = temp
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) opFB() {
	cpu.xce()
}
