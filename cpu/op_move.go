package cpu

func (cpu *CPU) mvn(SBank uint8, SAddress uint16, DBank uint8, DAddress uint16) {
	cpu.memory.SetByteBank(cpu.memory.GetByteBank(SBank, SAddress), DBank, DAddress)
	cpu.DBR = DBank
	cpu.C--
	if cpu.xFlag {
		cpu.setXLRegister(cpu.getXLRegister() + 1)
		cpu.setYLRegister(cpu.getYLRegister() + 1)
	} else {
		cpu.X++
		cpu.Y++
	}
	cpu.cycles += 7
	if cpu.getCRegister() == 0xFFFF {
		cpu.PC += 3
	}
}

func (cpu *CPU) op54() {
	SBank, SAddress, DBank, DAddress := cpu.admSourceDestination()
	cpu.mvn(SBank, SAddress, DBank, DAddress)
}

func (cpu *CPU) mvp(SBank uint8, SAddress uint16, DBank uint8, DAddress uint16) {
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

func (cpu *CPU) op44() {
	SBank, SAddress, DBank, DAddress := cpu.admSourceDestination()
	cpu.mvp(SBank, SAddress, DBank, DAddress)
}
