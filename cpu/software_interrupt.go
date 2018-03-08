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
	addressLo, addressHi := utils.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.bFlag = true
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, brkEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, brkEmulationVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, brkNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, brkNativeVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.cycles += 8 - utils.BoolToUint16[cpu.eFlag]
}

func (cpu *CPU) op00() {
	cpu.brk()
}

func (cpu *CPU) cop() {
	addressLo, addressHi := utils.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, copEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, copEmulationVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, copNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, copNativeVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.cycles += 8 - utils.BoolToUint16[cpu.eFlag]

}

func (cpu *CPU) op02() {
	cpu.cop()
}
