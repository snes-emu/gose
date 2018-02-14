package cpu

func (cpu *CPU) plp() {
	P := cpu.pullStack()
	cpu.cFlag = P&0x01 != 0
	cpu.zFlag = P&0x02 != 0
	cpu.iFlag = P&0x04 != 0
	cpu.dFlag = P&0x08 != 0
	cpu.mFlag = P&0x20 != 0
	cpu.vFlag = P&0x40 != 0
	cpu.nFlag = P&0x80 != 0
	if cpu.eFlag {
		cpu.bFlag = P&0x10 != 0
		cpu.xFlag = true
		cpu.mFlag = true
	} else {
		cpu.xFlag = P&0x10 != 0
	}
}

func (cpu *CPU) op28() {
	cpu.plp()
	cpu.cycles += 4
	cpu.PC++
}
