package core

import (
	"fmt"
	"os"
)

const HBLANKSTART = 1096
const HBLANKEND = 3

func (cpu *CPU) HandleIRQ() {
	switch cpu.ioMemory.hvIRQ {
	// 0: Disabled
	case 0:
		// 1: H=H and V=any
	case 1:
		cpu.irq()
		// 2: H=0 and V=V, 3: H=H and V=V
	case 2, 3:
		if cpu.ppu.VCounter() == cpu.ioMemory.virqPos {
			cpu.irq()
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown value for cpu.ioMemory.hvIRQ: %v", cpu.ioMemory.hvIRQ)
	}
}

func (cpu *CPU) enterVblank() {

	K := cpu.getKRegister()
	PC := cpu.getPCRegister()
	opcode := cpu.memory.GetByteBank(K, PC)
	cpu.logState(K, PC, opcode)
	cpu.ioMemory.vBlankNMIFlag = true
	if cpu.ioMemory.vBlankNMIEnable {
		cpu.nmi()
	}
}

func (cpu *CPU) leavVblank() {
	cpu.ioMemory.vBlankNMIFlag = false
}
