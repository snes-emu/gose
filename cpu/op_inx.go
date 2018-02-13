package cpu

import "github.com/snes-emu/gose/utils"

//opE8 performs a increment operation on the X register, immediate mode
func (cpu *CPU) opE8() {
	dataHi, dataLo := cpu.admImmediate()
	if cpu.mFlag {
		cpu.setARegister(cpu.inc8(dataLo))
	} else {
		cpu.setCRegister(cpu.inc16(utils.ReadUint16(dataHi, dataLo)))
	}
	cpu.cycles += 2
}
