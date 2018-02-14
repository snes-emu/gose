package cpu

import "github.com/snes-emu/gose/utils"

// ply16 pull the Y register from the stack
func (cpu *CPU) ply16() {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()

	result := utils.ReadUint16(dataHi, dataLo)

	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result != 0

	cpu.setYRegister(result)
}

// ply8 pull the lower bits of the Y register from the stack
func (cpu *CPU) ply8() {
	result := cpu.pullStack()

	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result != 0

	cpu.setYLRegister(result)
}

func (cpu *CPU) ply() {
	if cpu.xFlag {
		cpu.ply8()
	} else {
		cpu.ply16()
	}
}
