package cpu

import (
	"github.com/snes-emu/gose/utils"
)

const (
	abortNativeVector    = 0xFFE8
	nmiNativeVector      = 0xFFEA
	resetNativeVector    = 0xFFEC
	irqNativeVector      = 0xFFEE
	abortEmulationVector = 0xFFF8
	nmiEmulationVector   = 0xFFFA
	resetEmulationVector = 0xFFFC
	irqEmulationVector   = 0xFFFE
)

func (cpu *CPU) abort() {
	addressLo, addressHi := utils.SplitUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, abortEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, abortEmulationVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, abortNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, abortNativeVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) nmi() {
	addressLo, addressHi := utils.SplitUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, nmiEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, nmiEmulationVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, nmiNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, nmiNativeVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) reset() {
	cpu.setEFlag(true)
	cpu.D = 0x0000
	cpu.DBR = 0x00
	cpu.K = 0x00
	addressLo := cpu.memory.GetByteBank(0x00, resetEmulationVector)
	addressHi := cpu.memory.GetByteBank(0x00, resetEmulationVector+1)
	cpu.PC = utils.JoinUint16(addressLo, addressHi)
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) irq() {
	addressLo, addressHi := utils.SplitUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, irqEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, irqEmulationVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, irqNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, irqNativeVector+1)
		cpu.PC = utils.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}
