package cpu

import (
	"github.com/snes-emu/gose/utils"
)

const (
	brkNativeVector    = 0xFFE6
	copNativeVector    = 0xFFE4
	brkEmulationVector = 0xFFFE
	copEmulationVector = 0xFFF4
)

func (cpu *CPU) brk() {
	laddr, haddr := utils.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.bFlag = true
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, brkEmulationVector)
		haddr := cpu.memory.GetByteBank(0x00, brkEmulationVector+1)
		cpu.PC = utils.JoinUint16(laddr, haddr)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, brkNativeVector)
		haddr := cpu.memory.GetByteBank(0x00, brkNativeVector+1)
		cpu.PC = utils.JoinUint16(laddr, haddr)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.cycles += 8 - utils.BoolToUint16[cpu.eFlag]
}

func (cpu *CPU) op00() {
	cpu.brk()
}

func (cpu *CPU) cop() {
	laddr, haddr := utils.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, copEmulationVector)
		haddr := cpu.memory.GetByteBank(0x00, copEmulationVector+1)
		cpu.PC = utils.JoinUint16(laddr, haddr)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, copNativeVector)
		haddr := cpu.memory.GetByteBank(0x00, copNativeVector+1)
		cpu.PC = utils.JoinUint16(laddr, haddr)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.cycles += 8 - utils.BoolToUint16[cpu.eFlag]

}

func (cpu *CPU) op02() {
	cpu.cop()
}
