package cpu

import "github.com/snes-emu/gose/utils"

//op88 performs a decrement operation on the Y register, immediate mode
func (cpu *CPU) op88() {
	dataHi, dataLo := cpu.admImmediate()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(utils.ReadUint16(dataHi, dataLo)))
	}
	cpu.cycles += 2
}
