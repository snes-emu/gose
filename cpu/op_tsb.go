package cpu

import (
	"github.com/snes-emu/gose/utils"
)

// tsb16 test the bits of the data with the bits of the accumulator then set the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) tsb16(data uint16) uint16 {
	cpu.zFlag = cpu.getCRegister()&data == 0
	return data | cpu.getCRegister()

}

// tsb8 test the bits of the data with the bits of the accumulator then set the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) tsb8(data uint8) uint8 {
	cpu.zFlag = cpu.getARegister()&data == 0
	return data | cpu.getARegister()
}

// tsb test the bits of the data with the bits of the accumulator then set the bits of the data that are ones in the accumulator handling the 8/16 case
func (cpu *CPU) tsb(addressHi, addressLo uint32) {
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.tsb8(cpu.memory.GetByte(addressLo)), addressLo)
	} else {
		result := cpu.tsb16(utils.JoinUint16(cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)))
		resultHi, resultLo := utils.SplitUint16(result)
		cpu.memory.SetByte(resultLo, addressLo)
		cpu.memory.SetByte(resultHi, addressHi)
	}
}

func (cpu *CPU) op04() {
	addressHi, addressLo := cpu.admDirectP()
	cpu.tsb(addressHi, addressLo)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}
func (cpu *CPU) op0C() {
	addressHi, addressLo := cpu.admAbsoluteP()
	cpu.tsb(addressHi, addressLo)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
}
