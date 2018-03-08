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
		cpu.setCRegister(cpu.dec16(utils.JoinUint16(dataLo, dataHi)))
	}
	cpu.cycles += 2
	cpu.PC++
}

//opC6 performs a decrement operation on memory through direct addressing mode
func (cpu *CPU) opC6() {
	addressHi, addressLo := cpu.admDirectP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.dec16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

//opCE performs a decrement operation on memory through the absolute addressing mode
func (cpu *CPU) opCE() {
	addressHi, addressLo := cpu.admAbsoluteP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.dec16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

//opD6 performs a decrement operation on memory through direct,X addressing mode
func (cpu *CPU) opD6() {
	addressHi, addressLo := cpu.admDirectXP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.dec16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

//opDE performs a decrement operation on memory through absolute,X addressing mode
func (cpu *CPU) opDE() {
	addressHi, addressLo := cpu.admAbsoluteXP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.dec16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
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
	cpu.PC++
}

//op88 performs a decrement operation on the Y register, immediate mode
func (cpu *CPU) op88() {
	if cpu.xFlag {
		result := cpu.getYLRegister() - 1
		cpu.setYLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.Y--
		// Last bit value
		cpu.nFlag = cpu.Y&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.Y == 0
	}
	cpu.cycles += 2
	cpu.PC++
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
		cpu.setCRegister(cpu.inc16(utils.JoinUint16(dataLo, dataHi)))
	}
	cpu.cycles += 2
	cpu.PC++
}

//opE6 performs a increment operation on memory through direct addressing mode
func (cpu *CPU) opE6() {
	addressHi, addressLo := cpu.admDirectP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.inc16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

//opEE performs a increment operation through the absolute access mode
func (cpu *CPU) opEE() {
	addressHi, addressLo := cpu.admAbsoluteP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.inc16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

//opF6 performs a increment operation on memory through direct,X addressing mode
func (cpu *CPU) opF6() {
	addressHi, addressLo := cpu.admDirectXP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.inc16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

//opF6 performs a increment operation on memory through absolute,X addressing mode
func (cpu *CPU) opFE() {
	addressHi, addressLo := cpu.admAbsoluteXP()
	dataHi, dataLo := cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), addressLo)
	} else {
		resultLo, resultHi := utils.SplitUint16(cpu.inc16(utils.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, addressHi)
		cpu.memory.SetByte(resultLo, addressLo)
	}
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

//opE8 performs a increment operation on the X register, immediate mode
func (cpu *CPU) opE8() {
	if cpu.xFlag {
		result := cpu.getXLRegister() + 1
		cpu.setXLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.X++
		// Last bit value
		cpu.nFlag = cpu.X&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.X == 0
	}
	cpu.cycles += 2
	cpu.PC++
}

//opC8 performs a increment operation on the Y register, immediate mode
func (cpu *CPU) opC8() {
	if cpu.xFlag {
		result := cpu.getYLRegister() + 1
		cpu.setXLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.Y++
		// Last bit value
		cpu.nFlag = cpu.Y&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.Y == 0
		cpu.cycles += 2
	}
	cpu.cycles += 2
	cpu.PC++
}
