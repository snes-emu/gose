package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op60() uint32 {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()
	counter := utils.ReadUint16(dataHi, dataLo) + 1
	kvalue := cpu.getKRegister()
	dataHi, dataLo = utils.WriteUint16(counter)
	result := utils.ReadUint32(kvalue, dataHi, dataLo)
	return result
}
