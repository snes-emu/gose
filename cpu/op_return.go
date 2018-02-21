package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) rti() {
	cpu.plp()
	addressLo := cpu.pullStack()
	addressHi := cpu.pullStack()
	cpu.PC = utils.JoinUint16(addressHi, addressLo)
	if !cpu.eFlag {
		cpu.K = cpu.pullStack()
	}
	cpu.cycles += 7 - utils.BoolToUint16[cpu.eFlag]

}

func (cpu *CPU) op40() {
	cpu.rti()
}

func (cpu *CPU) rtl() {
	PCLo, PCHi, K := cpu.pullStackNew24()
	cpu.cycles += 6
	cpu.K = K
	cpu.PC = utils.JoinUint16(PCHi, PCLo) + 1
}

func (cpu *CPU) op6B() {
	cpu.rtl()
}

func (cpu *CPU) rts() {
	PCLo := cpu.pullStack()
	PCHi := cpu.pullStack()
	cpu.cycles += 6
	cpu.PC = utils.JoinUint16(PCHi, PCLo) + 1
}

func (cpu *CPU) op60() {
	cpu.rts()
}
