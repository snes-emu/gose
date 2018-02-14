package cpu

import (
	"github.com/snes-emu/gose/utils"
)

func (cpu *CPU) abort() {
	addressHi, addressLo := utils.WriteUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFF8)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFF9)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFE8)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFE9)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) nmi() {
	addressHi, addressLo := utils.WriteUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFFA)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFFB)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFEA)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFEB)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) reset() {
	addressHi, addressLo := utils.WriteUint16(cpu.getPCRegister())
	cpu.eFlag = true
	cpu.pushStack(addressHi)
	cpu.pushStack(addressLo)
	cpu.php()
	cpu.K = 0x00
	addressLo = cpu.memory.GetByteBank(0x00, 0xFFFC)
	addressHi = cpu.memory.GetByteBank(0x00, 0xFFFD)
	cpu.PC = utils.ReadUint16(addressHi, addressLo)
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) irq() {
	addressHi, addressLo := utils.WriteUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFFE)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFFF)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, 0xFFEE)
		addressHi := cpu.memory.GetByteBank(0x00, 0xFFEF)
		cpu.PC = utils.ReadUint16(addressHi, addressLo)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}
