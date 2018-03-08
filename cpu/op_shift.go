package cpu

import "github.com/snes-emu/gose/utils"

// asl16acc performs a left shift on the 16 bit accumulator
func (cpu *CPU) asl16acc() {
	result := cpu.getCRegister() << 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// asl8 performs a left shift on the lower 8 bit accumulator
func (cpu *CPU) asl8acc() {
	result := cpu.getARegister() << 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getARegister()&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// asl16data performs a left shift on the 16 bit data
func (cpu *CPU) asl16data(haddress, laddress uint32) {
	dataHi, dataLo := cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)

	data := utils.JoinUint16(dataLo, dataHi)

	result := data << 1

	// Get the highbit before shifting
	cpu.cFlag = data&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddress)
	cpu.memory.SetByte(resultLo, laddress)
}

// asl8data performs a left shift on the 8 bit data
func (cpu *CPU) asl8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data << 1

	// Get the highbit before shifting
	cpu.cFlag = data&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// asl performs a left shift taking care of 16/8bits cases
func (cpu *CPU) asl(haddress, laddress uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.asl8acc()
		} else {
			cpu.asl8data(laddress)
		}
	} else {
		if isAcc {
			cpu.asl16acc()
		} else {
			cpu.asl16data(haddress, laddress)
		}
	}
}

func (cpu *CPU) op06() {
	addrHi, addrLo := cpu.admDirectP()
	cpu.asl(addrHi, addrLo, false)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op0A() {
	cpu.asl(0, 0, true)
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op0E() {
	addrHi, addrLo := cpu.admAbsoluteP()
	cpu.asl(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op16() {
	addrHi, addrLo := cpu.admDirectXP()
	cpu.asl(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op1E() {
	addrHi, addrLo := cpu.admAbsoluteXP()
	cpu.asl(addrHi, addrLo, false)
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

// lsr16acc performs a right shift on the 16 bit accumulator
func (cpu *CPU) lsr16acc() {
	result := cpu.getCRegister() >> 1

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// lsr8 performs a right shift on the lower 8 bit accumulator
func (cpu *CPU) lsr8acc() {
	result := cpu.getARegister() >> 1

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getARegister()&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// lsr16data performs a right shift on the 16 bit data
func (cpu *CPU) lsr16data(haddress, laddress uint32) {
	dataHi, dataLo := cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)

	data := utils.JoinUint16(dataLo, dataHi)

	result := data >> 1

	// Get the lowbit before shifting
	cpu.cFlag = data&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddress)
	cpu.memory.SetByte(resultLo, laddress)
}

// lsr8data performs a right shift on the 8 bit data
func (cpu *CPU) lsr8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data >> 1

	// Get the lowbit before shifting
	cpu.cFlag = data&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// lsr performs a right shift taking care of 16/8bits cases
func (cpu *CPU) lsr(haddress, laddress uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.lsr8acc()
		} else {
			cpu.lsr8data(laddress)
		}
	} else {
		if isAcc {
			cpu.lsr16acc()
		} else {
			cpu.lsr16data(haddress, laddress)
		}
	}
}

func (cpu *CPU) op46() {
	addrHi, addrLo := cpu.admDirectP()
	cpu.lsr(addrHi, addrLo, false)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op4A() {
	cpu.lsr(0, 0, true)
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op4E() {
	addrHi, addrLo := cpu.admAbsoluteP()
	cpu.lsr(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op56() {
	addrHi, addrLo := cpu.admDirectXP()
	cpu.lsr(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op5E() {
	addrHi, addrLo := cpu.admAbsoluteXP()
	cpu.lsr(addrHi, addrLo, false)
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

// rol16acc performs a rotate left on the 16 bit accumulator
func (cpu *CPU) rol16acc() {
	result := cpu.getCRegister() << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// rol8 performs a rotate left on the lower 8 bit accumulator
func (cpu *CPU) rol8acc() {
	result := cpu.getARegister() << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = cpu.getARegister()&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// rol16data performs a rotate left on the 16 bit data
func (cpu *CPU) rol16data(haddress, laddress uint32) {
	dataHi, dataLo := cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)

	data := utils.JoinUint16(dataLo, dataHi)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = data&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddress)
	cpu.memory.SetByte(resultLo, laddress)
}

// rol8data performs a rotate left on the 8 bit data
func (cpu *CPU) rol8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = data&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// rol performs a rotate left taking care of 16/8bits cases
func (cpu *CPU) rol(haddress, laddress uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.rol8acc()
		} else {
			cpu.rol8data(laddress)
		}
	} else {
		if isAcc {
			cpu.rol16acc()
		} else {
			cpu.rol16data(haddress, laddress)
		}
	}
}

func (cpu *CPU) op26() {
	addrHi, addrLo := cpu.admDirectP()
	cpu.rol(addrHi, addrLo, false)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op2A() {
	cpu.rol(0, 0, true)
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op2E() {
	addrHi, addrLo := cpu.admAbsoluteP()
	cpu.rol(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op36() {
	addrHi, addrLo := cpu.admDirectXP()
	cpu.rol(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op3E() {
	addrHi, addrLo := cpu.admAbsoluteXP()
	cpu.rol(addrHi, addrLo, false)
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

// ror16acc performs a rotate right on the 16 bit accumulator
func (cpu *CPU) ror16acc() {
	result := cpu.getCRegister() >> 1

	if cpu.cFlag {
		result = result | 0x8000
	}

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// ror8 performs a rotate right on the lower 8 bit accumulator
func (cpu *CPU) ror8acc() {
	result := cpu.getARegister() << 1

	if cpu.cFlag {
		result = result | 0x80
	}

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getARegister()&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// ror16data performs a rotate right on the 16 bit data
func (cpu *CPU) ror16data(haddress, laddress uint32) {
	dataHi, dataLo := cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)

	data := utils.JoinUint16(dataLo, dataHi)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x8000
	}

	// Get the lowbit before shifting
	cpu.cFlag = data&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddress)
	cpu.memory.SetByte(resultLo, laddress)
}

// ror8data performs a rotate right on the 8 bit data
func (cpu *CPU) ror8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x80
	}

	// Get the lowbit before shifting
	cpu.cFlag = data&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// ror performs a rotate right taking care of 16/8bits cases
func (cpu *CPU) ror(haddress, laddress uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.ror8acc()
		} else {
			cpu.ror8data(laddress)
		}
	} else {
		if isAcc {
			cpu.ror16acc()
		} else {
			cpu.ror16data(haddress, laddress)
		}
	}
}

func (cpu *CPU) op66() {
	addrHi, addrLo := cpu.admDirectP()
	cpu.ror(addrHi, addrLo, false)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op6A() {
	cpu.ror(0, 0, true)
	cpu.cycles += 2
	cpu.PC++
}

func (cpu *CPU) op6E() {
	addrHi, addrLo := cpu.admAbsoluteP()
	cpu.ror(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op76() {
	addrHi, addrLo := cpu.admDirectXP()
	cpu.ror(addrHi, addrLo, false)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op7E() {
	addrHi, addrLo := cpu.admAbsoluteXP()
	cpu.ror(addrHi, addrLo, false)
	cpu.cycles += 9 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}
