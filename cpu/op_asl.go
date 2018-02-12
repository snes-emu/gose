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

	data := utils.ReadUint16(dataHi, dataLo)

	result := data << 1

	// Get the highbit before shifting
	cpu.cFlag = data&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.WriteUint16(data)

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
