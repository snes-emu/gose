package cpu

import "github.com/snes-emu/gose/memory"

type CPU struct {
	C       uint16 // Accumulator register
	DBR     uint8  // Data bank register
	D       uint16 // The direct register
	K       uint8  // The program bank register
	PC      uint16 // The program counter
	P       uint16 // Processor status register
	S       uint16 // The stack pointer register
	X       uint16 // The X index register
	Y       uint16 // The Y index register
	cycles  uint   // Number of cycles
	memory  memory.Memory
	opcodes []cpuOperation
}

type cpuOperation func()

var opcodes []cpuOperation

func makeCPU() CPU {
	cpu := CPU{}
	cpu.opcodes[0x61] = cpu.op61
	return cpu
}
