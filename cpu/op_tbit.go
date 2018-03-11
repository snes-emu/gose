package cpu

import (
	"github.com/snes-emu/gose/utils"
)

// trb16 test the bits of the data with the bits of the accumulator then reset the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) trb16(data uint16) uint16 {
	cpu.zFlag = cpu.getCRegister()&data == 0
	return data &^ cpu.getCRegister()

}

// trb8 test the bits of the data with the bits of the accumulator then reset the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) trb8(data uint8) uint8 {
	cpu.zFlag = cpu.getARegister()&data == 0
	return data &^ cpu.getARegister()
}

// trb test the bits of the data with the bits of the accumulator then reset the bits of the data that are ones in the accumulator handling the 8/16 case
func (cpu *CPU) trb(laddr, haddr uint32) {
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.trb8(cpu.memory.GetByte(laddr)), laddr)
	} else {
		result := cpu.trb16(utils.JoinUint16(cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)))
		resultLo, resultHi := utils.SplitUint16(result)
		cpu.memory.SetByte(resultLo, laddr)
		cpu.memory.SetByte(resultHi, haddr)
	}
}

func (cpu *CPU) op14() {
	laddr, haddr := cpu.admDirectP()
	cpu.trb(laddr, haddr)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}
func (cpu *CPU) op1C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.trb(laddr, haddr)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

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
func (cpu *CPU) tsb(laddr, haddr uint32) {
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.tsb8(cpu.memory.GetByte(laddr)), laddr)
	} else {
		result := cpu.tsb16(utils.JoinUint16(cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)))
		resultLo, resultHi := utils.SplitUint16(result)
		cpu.memory.SetByte(resultLo, laddr)
		cpu.memory.SetByte(resultHi, haddr)
	}
}

func (cpu *CPU) op04() {
	laddr, haddr := cpu.admDirectP()
	cpu.tsb(laddr, haddr)
	cpu.cycles += 7 - 2*utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op0C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.tsb(laddr, haddr)
	cpu.cycles += 8 - 2*utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}
