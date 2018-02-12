package cpu

import "github.com/snes-emu/gose/utils"

// lsr16acc performs a right shift on the 16 bit accumulator
func (cpu *CPU) lsr16acc() {
	result := cpu.getCRegister() >> 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// lsr8 performs a right shift on the lower 8 bit accumulator
func (cpu *CPU) lsr8acc() {
	result := cpu.getARegister() >> 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getARegister()&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// lsr16data performs a right shift on the 16 bit data
func (cpu *CPU) lsr16data(haddress, laddress uint32) {
	dataHi, dataLo := cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)

	data := utils.ReadUint16(dataHi, dataLo)

	result := data >> 1

	// Get the highbit before shifting
	cpu.cFlag = data&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.WriteUint16(data)

	cpu.memory.SetByte(resultHi, haddress)
	cpu.memory.SetByte(resultLo, laddress)
}

// lsr8data performs a right shift on the 8 bit data
func (cpu *CPU) lsr8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data >> 1

	// Get the highbit before shifting
	cpu.cFlag = data&0x80 != 0

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
	cpu.PC += 1
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
