package cpu

func (cpu *CPU) clc() {
	cpu.cFlag = false
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op18() {
	cpu.clc()
}

func (cpu *CPU) cld() {
	cpu.dFlag = false
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) opD8() {
	cpu.cld()
}

func (cpu *CPU) cli() {
	cpu.iFlag = false
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op58() {
	cpu.cli()
}

func (cpu *CPU) clv() {
	cpu.vFlag = false
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) opB8() {
	cpu.clv()
}

func (cpu *CPU) sec() {
	cpu.cFlag = true
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op38() {
	cpu.sec()
}

func (cpu *CPU) sed() {
	cpu.dFlag = true
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) opF8() {
	cpu.sed()
}

func (cpu *CPU) sei() {
	cpu.iFlag = true
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op78() {
	cpu.sei()
}

func (cpu *CPU) rep() {
	_, dataLo := cpu.admImmediate8()
	cpu.cFlag = cpu.cFlag && dataLo&0x01 == 0
	cpu.zFlag = cpu.zFlag && dataLo&0x02 == 0
	cpu.iFlag = cpu.iFlag && dataLo&0x04 == 0
	cpu.dFlag = cpu.dFlag && dataLo&0x08 == 0
	if cpu.eFlag {
		cpu.bFlag = cpu.bFlag && dataLo&0x10 == 0
	} else {
		cpu.xFlag = cpu.xFlag && dataLo&0x10 == 0
	}
	cpu.mFlag = cpu.mFlag && dataLo&0x20 == 0
	cpu.vFlag = cpu.vFlag && dataLo&0x40 == 0
	cpu.nFlag = cpu.nFlag && dataLo&0x80 == 0
	cpu.cycles += 3
	cpu.PC += 2

}

func (cpu *CPU) opC2() {
	cpu.rep()
}

func (cpu *CPU) sep() {
	_, dataLo := cpu.admImmediate8()
	cpu.cFlag = cpu.cFlag || dataLo&0x01 != 0
	cpu.zFlag = cpu.zFlag || dataLo&0x02 != 0
	cpu.iFlag = cpu.iFlag || dataLo&0x04 != 0
	cpu.dFlag = cpu.dFlag || dataLo&0x08 != 0
	if cpu.eFlag {
		cpu.bFlag = cpu.bFlag || dataLo&0x10 != 0
	} else {
		cpu.xFlag = cpu.xFlag || dataLo&0x10 != 0
	}
	cpu.mFlag = cpu.mFlag || dataLo&0x20 != 0
	cpu.vFlag = cpu.vFlag || dataLo&0x40 != 0
	cpu.nFlag = cpu.nFlag || dataLo&0x80 != 0
	cpu.cycles += 3
	cpu.PC += 2
}

func (cpu *CPU) opE2() {
	cpu.sep()
}
