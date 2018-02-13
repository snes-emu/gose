package cpu

import "github.com/snes-emu/gose/utils"

//opCA performs a decrement operation on the X register, immediate mode
func (cpu *CPU) opCA() {
	dataHi, dataLo := cpu.admImmediate()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(utils.ReadUint16(dataHi, dataLo)))
	}
	cpu.cycles += 2
}
