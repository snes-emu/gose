package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) dec16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data - 1
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

func (cpu *CPU) dec8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data - 1
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

//op3A performs a decrement operation on the accumulator
func (cpu *CPU) op3A() {
	dataHi, dataLo := cpu.admAccumulator()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 2
}

//opC6 performs a decrement operation on the D register
func (cpu *CPU) opC6() {
	dataHi, dataLo := cpu.admDirect()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

//opCE performs a decrement operation through the absolute access mode
func (cpu *CPU) opCE() {
	dataHi, dataLo := cpu.admAbsolute()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
}

//opD6 performs a decrement operation on the X register, direct mode
func (cpu *CPU) opD6() {
	dataHi, dataLo := cpu.admDirectX()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

//opDE performs a decrement operation on the X register, absolute mode
func (cpu *CPU) opDE() {
	dataHi, dataLo := cpu.admAbsoluteX()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
}

//opCA performs a decrement operation on the X register
func (cpu *CPU) opCA() {
	if cpu.xFlag {
		result := cpu.getXLRegister() - 1
		cpu.setXLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.X--
		// Last bit value
		cpu.nFlag = cpu.X&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.X == 0
	}
	cpu.cycles += 2
}

//op88 performs a decrement operation on the Y register, immediate mode
func (cpu *CPU) op88() {
	if cpu.xFlag {
		cpu.setYLRegister(cpu.getYLRegister() - 1)
	} else {
		cpu.Y--
	}
	cpu.cycles += 2
}

func (cpu *CPU) inc16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data + 1
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

func (cpu *CPU) inc8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data + 1
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

//op1A performs a increment operation on the accumulator
func (cpu *CPU) op1A() {
	dataHi, dataLo := cpu.admAccumulator()
	if cpu.mFlag {
		cpu.setARegister(cpu.inc8(dataLo))
	} else {
		cpu.setCRegister(cpu.inc16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 2
}

//opE6 performs a increment operation on the D register
func (cpu *CPU) opE6() {
	dataHi, dataLo := cpu.admDirect()
	if cpu.mFlag {
		cpu.setARegister(cpu.inc8(dataLo))
	} else {
		cpu.setCRegister(cpu.inc16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

//opEE performs a increment operation through the absolute access mode
func (cpu *CPU) opEE() {
	dataHi, dataLo := cpu.admAbsolute()
	if cpu.mFlag {
		cpu.setARegister(cpu.inc8(dataLo))
	} else {
		cpu.setCRegister(cpu.inc16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
}

//opF6 performs a increment operation on the X register, direct mode
func (cpu *CPU) opF6() {
	dataHi, dataLo := cpu.admDirectX()
	if cpu.mFlag {
		cpu.setARegister(cpu.inc8(dataLo))
	} else {
		cpu.setCRegister(cpu.inc16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}

//opF6 performs a increment operation on the X register, absolute mode
func (cpu *CPU) opFE() {
	dataHi, dataLo := cpu.admAbsoluteX()
	if cpu.mFlag {
		cpu.setARegister(cpu.inc8(dataLo))
	} else {
		cpu.setCRegister(cpu.inc16(utils.JoinUint16(dataHi, dataLo)))
	}
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
}

//opE8 performs a increment operation on the X register, immediate mode
func (cpu *CPU) opE8() {
	if cpu.xFlag {
		cpu.setXLRegister(cpu.getXLRegister() + 1)
	} else {
		cpu.X++
	}
	cpu.cycles += 2
}

//opC8 performs a increment operation on the Y register, immediate mode
func (cpu *CPU) opC8() {
	if cpu.xFlag {
		cpu.setYLRegister(cpu.getYLRegister() + 1)
	} else {
		cpu.Y++
	}
	cpu.cycles += 2
}
