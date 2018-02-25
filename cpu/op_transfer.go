package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) tcd() {
	cpu.D = cpu.C
	// Last bit value
	cpu.nFlag = cpu.D&0x8000 != 0
	cpu.zFlag = cpu.D == 0
	cpu.cycles += 2
}

func (cpu *CPU) op5B() {
	cpu.tcd()
}

func (cpu *CPU) tcs() {
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	if cpu.eFlag {
		_, dataLo := utils.SplitUint16(cpu.C)
		cpu.S = utils.JoinUint16(0x01, dataLo)
	} else {
		cpu.S = cpu.C
	}
	cpu.cycles += 2
}

func (cpu *CPU) op1B() {
	cpu.tcs()
}

func (cpu *CPU) tdc() {
	cpu.C = cpu.D
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	cpu.cycles += 2
}

func (cpu *CPU) op7B() {
	cpu.tdc()
}

func (cpu *CPU) tsc() {
	cpu.C = cpu.S
	// Last bit value
	cpu.nFlag = cpu.S&0x8000 != 0
	cpu.zFlag = cpu.S == 0
	cpu.cycles += 2
}

func (cpu *CPU) op3B() {
	cpu.tsc()
}

func (cpu *CPU) tax() {
	if cpu.xFlag {
		result := cpu.getARegister()
		cpu.setXLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.X = cpu.C
		cpu.nFlag = cpu.X&0x8000 != 0
		cpu.xFlag = cpu.X == 0
	}
}

func (cpu *CPU) opAA() {
	cpu.tax()
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) tay() {
	if cpu.xFlag {
		result := cpu.getARegister()
		cpu.setYLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.Y = cpu.C
		cpu.nFlag = cpu.Y&0x8000 != 0
		cpu.xFlag = cpu.Y == 0
	}
}

func (cpu *CPU) opA8() {
	cpu.tay()
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) tsx() {
	if cpu.xFlag {
		result := cpu.getSLRegister()
		cpu.setXLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.X = cpu.S
		cpu.nFlag = cpu.X&0x8000 != 0
		cpu.xFlag = cpu.X == 0
	}
}

func (cpu *CPU) opBA() {
	cpu.tsx()
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) txa() {
	if cpu.mFlag {
		result := cpu.getXLRegister()
		cpu.setARegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.C = cpu.X
		cpu.nFlag = cpu.C&0x8000 != 0
		cpu.xFlag = cpu.C == 0
	}
}

func (cpu *CPU) op8A() {
	cpu.txa()
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) txs() {
	if cpu.eFlag {
		result := cpu.getXLRegister()
		cpu.setSLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.S = cpu.X
		cpu.nFlag = cpu.S&0x8000 != 0
		cpu.xFlag = cpu.S == 0
	}
}

func (cpu *CPU) op9A() {
	cpu.txs()
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) txy() {
	if cpu.eFlag {
		result := cpu.getXLRegister()
		cpu.setYLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.Y = cpu.X
		cpu.nFlag = cpu.Y&0x8000 != 0
		cpu.xFlag = cpu.Y == 0
	}
}

func (cpu *CPU) op9B() {
	cpu.txy()
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) tya() {
	if cpu.eFlag {
		result := cpu.getYLRegister()
		cpu.setARegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.C = cpu.Y
		cpu.nFlag = cpu.C&0x8000 != 0
		cpu.xFlag = cpu.C == 0
	}
}

func (cpu *CPU) op98() {
	cpu.tya()
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) tyx() {
	if cpu.eFlag {
		result := cpu.getYLRegister()
		cpu.setXLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.xFlag = result == 0
	} else {
		cpu.X = cpu.Y
		cpu.nFlag = cpu.X&0x8000 != 0
		cpu.xFlag = cpu.X == 0
	}
}

func (cpu *CPU) opBB() {
	cpu.tyx()
	cpu.cycles += 2
	cpu.PC++
}
