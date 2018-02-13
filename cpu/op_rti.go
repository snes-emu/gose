package cpu

import (
	"github.com/snes-emu/gose/utils"
)

func (cpu *CPU) op40() {
	P := cpu.pullStack()
	addressLo := cpu.pullStack()
	addressHi := cpu.pullStack()
	cpu.cFlag = P&0x01 != 0
	cpu.zFlag = P&0x02 != 0
	cpu.iFlag = P&0x04 != 0
	cpu.dFlag = P&0x08 != 0
	cpu.mFlag = P&0x20 != 0
	cpu.vFlag = P&0x40 != 0
	cpu.nFlag = P&0x80 != 0
	cpu.PC = utils.ReadUint16(addressHi, addressLo)
	if cpu.eFlag {
		cpu.bFlag = P&0x10 != 0
	} else {
		cpu.xFlag = P&0x10 != 0
		cpu.K = cpu.pullStack()
	}
	cpu.cycles += 7 - utils.BoolToUint16[cpu.eFlag]

}
