package cpu

type CPU struct {
	AB     uint16 // Accumulator register
	DBR    uint8  // Data bank register
	D      uint16 // The direct register
	K      uint8  // The program bank register
	PC     uint16 // The program counter
	P      uint16 // Processor status register
	S      uint16 // The stack pointer register
	X      uint16 // The X index register
	Y      uint16 // The Y index register
	cycles int    // Number of cycles
}

type cpuOperation func()

var opcodes []cpuOperation
