package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op6B() uint32 {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()
	counter := utils.ReadUint16(dataHi, dataLo) + 1
	kvalue := cpu.pullStack()
	dataHi, dataLo = utils.WriteUint16(counter)
	result := utils.ReadUint32(kvalue, dataHi, dataLo)
	return result
}
