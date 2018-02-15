package cpu

func (cpu *CPU) op44() {
	SBank, SAddress, DBank, DAddress := cpu.admSourceDestination()
	cpu.memory.SetByteBank(cpu.memory.GetByteBank(SBank, SAddress), DBank, DAddress)
	cpu.DBR = DBank
	cpu.C--
	if cpu.xFlag {
		cpu.setXLRegister(cpu.getXLRegister() - 1)
		cpu.setYLRegister(cpu.getYLRegister() - 1)
	} else {
		cpu.X--
		cpu.Y--
	}
	cpu.cycles += 7
	if cpu.getCRegister() == 0xFFFF {
		cpu.PC += 3
	}
}
