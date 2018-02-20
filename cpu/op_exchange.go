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
	cpu.eFlag, cpu.cFlag = cpu.cFlag, cpu.eFlag
	if cpu.eFlag {
		cpu.mFlag = true
		cpu.setXFlag(true)
		cpu.setSHRegister(0x01)
	}
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) opFB() {
	cpu.xce()
}
