package cpu

import "github.com/snes-emu/gose/utils"

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

	data := utils.ReadUint16(dataHi, dataLo)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x8000
	}

	// Get the lowbit before shifting
	cpu.cFlag = data&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.WriteUint16(data)

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
