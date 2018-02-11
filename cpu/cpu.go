package cpu

import "github.com/snes-emu/gose/memory"

// CPU represents the cpu 65C816
type CPU struct {
	C       uint16 // Accumulator register
	DBR     uint8  // Data bank register
	D       uint16 // The direct register
	K       uint8  // The program bank register
	PC      uint16 // The program counter
	eFlag   bool   // The emulation flag
	nFlag   bool   // The negative flag
	vFlag   bool   // The overflow flag
	mFlag   bool   // The accumulator and memory width flag
	bFlag   bool   // The break flag
	xFlag   bool   // The index register width flag
	dFlag   bool   // The decimal mode flag
	iFlag   bool   // The interrupt disable flag
	zFlag   bool   // The zero flag
	cFlag   bool   // The carry flag
	S       uint16 // The stack pointer register
	X       uint16 // The X index register
	Y       uint16 // The Y index register
	cycles  uint16 // Number of cycles
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
