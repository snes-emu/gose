package cpu

import (
	"github.com/snes-emu/gose/utils"
)

func (cpu *CPU) op00() {
	P := utils.BoolToUint8[cpu.cFlag]*0x01 +
		utils.BoolToUint8[cpu.zFlag]*0x02 +
		utils.BoolToUint8[cpu.iFlag]*0x04 +
		utils.BoolToUint8[cpu.dFlag]*0x08 +
		utils.BoolToUint8[cpu.mFlag]*0x20 +
		utils.BoolToUint8[cpu.vFlag]*0x40 +
		utils.BoolToUint8[cpu.nFlag]*0x80
	addressHi, addressLo := utils.WriteUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.bFlag = true
		P += utils.BoolToUint8[cpu.bFlag] * 0x10
		cpu.pushStack(P)
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFFE)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFFF)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		P += utils.BoolToUint8[cpu.xFlag] * 0x10
		cpu.pushStack(P)
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFE6)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFE7)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.cycles += 8 - utils.BoolToUint16[cpu.eFlag]

}
