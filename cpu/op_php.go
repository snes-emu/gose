package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) php() {
	P := utils.BoolToUint8[cpu.cFlag]*0x01 +
		utils.BoolToUint8[cpu.zFlag]*0x02 +
		utils.BoolToUint8[cpu.iFlag]*0x04 +
		utils.BoolToUint8[cpu.dFlag]*0x08 +
		utils.BoolToUint8[cpu.mFlag]*0x20 +
		utils.BoolToUint8[cpu.vFlag]*0x40 +
		utils.BoolToUint8[cpu.nFlag]*0x80
	if cpu.eFlag {
		P += utils.BoolToUint8[cpu.bFlag] * 0x10
	} else {
		P += utils.BoolToUint8[cpu.xFlag] * 0x10
	}
	cpu.pushStack(P)
}

func (cpu *CPU) op08() {
	cpu.php()
	cpu.cycles += 3
	cpu.PC++
}
