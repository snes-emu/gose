package cpu

import (
	"github.com/snes-emu/gose/utils"
)

func (cpu *CPU) op00() {
	addressHi, addressLo := utils.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.bFlag = true
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFFE)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFFF)
		cpu.PC = utils.JoinUint16(addressHi, addressLo)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFE6)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFE7)
		cpu.PC = utils.JoinUint16(addressHi, addressLo)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.cycles += 8 - utils.BoolToUint16[cpu.eFlag]

}

func (cpu *CPU) op02() {
	addressHi, addressLo := utils.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFF4)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFF5)
		cpu.PC = utils.JoinUint16(addressHi, addressLo)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFE4)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFE5)
		cpu.PC = utils.JoinUint16(addressHi, addressLo)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.cycles += 8 - utils.BoolToUint16[cpu.eFlag]

}
