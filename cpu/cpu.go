package cpu

import (
	"fmt"

	"github.com/snes-emu/gose/memory"
	"github.com/snes-emu/gose/utils"
)

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
	pFlag   bool   // page boundary crossed virtual flag
	S       uint16 // The stack pointer register
	X       uint16 // The X index register
	Y       uint16 // The Y index register
	cycles  uint16 // Number of cycles
	memory  *memory.Memory
	opcodes [256]cpuOperation
}

type cpuOperation func()

var opcodes []cpuOperation

func New(memory *memory.Memory) *CPU {
	cpu := &CPU{memory: memory}
	cpu.opcodes[0x0] = cpu.op00
	cpu.opcodes[0x1] = cpu.op01
	cpu.opcodes[0x2] = cpu.op02
	cpu.opcodes[0x3] = cpu.op03
	cpu.opcodes[0x4] = cpu.op04
	cpu.opcodes[0x5] = cpu.op05
	cpu.opcodes[0x6] = cpu.op06
	cpu.opcodes[0x7] = cpu.op07
	cpu.opcodes[0x8] = cpu.op08
	cpu.opcodes[0x9] = cpu.op09
	cpu.opcodes[0xA] = cpu.op0A
	cpu.opcodes[0xB] = cpu.op0B
	cpu.opcodes[0xC] = cpu.op0C
	cpu.opcodes[0xD] = cpu.op0D
	cpu.opcodes[0xE] = cpu.op0E
	cpu.opcodes[0xF] = cpu.op0F
	cpu.opcodes[0x10] = cpu.op10
	cpu.opcodes[0x11] = cpu.op11
	cpu.opcodes[0x12] = cpu.op12
	cpu.opcodes[0x13] = cpu.op13
	cpu.opcodes[0x14] = cpu.op14
	cpu.opcodes[0x15] = cpu.op15
	cpu.opcodes[0x16] = cpu.op16
	cpu.opcodes[0x17] = cpu.op17
	cpu.opcodes[0x18] = cpu.op18
	cpu.opcodes[0x19] = cpu.op19
	cpu.opcodes[0x1A] = cpu.op1A
	cpu.opcodes[0x1B] = cpu.op1B
	cpu.opcodes[0x1C] = cpu.op1C
	cpu.opcodes[0x1D] = cpu.op1D
	cpu.opcodes[0x1E] = cpu.op1E
	cpu.opcodes[0x1F] = cpu.op1F
	cpu.opcodes[0x20] = cpu.op20
	cpu.opcodes[0x21] = cpu.op21
	cpu.opcodes[0x22] = cpu.op22
	cpu.opcodes[0x23] = cpu.op23
	cpu.opcodes[0x24] = cpu.op24
	cpu.opcodes[0x25] = cpu.op25
	cpu.opcodes[0x26] = cpu.op26
	cpu.opcodes[0x27] = cpu.op27
	cpu.opcodes[0x28] = cpu.op28
	cpu.opcodes[0x29] = cpu.op29
	cpu.opcodes[0x2A] = cpu.op2A
	cpu.opcodes[0x2B] = cpu.op2B
	cpu.opcodes[0x2C] = cpu.op2C
	cpu.opcodes[0x2D] = cpu.op2D
	cpu.opcodes[0x2E] = cpu.op2E
	cpu.opcodes[0x2F] = cpu.op2F
	cpu.opcodes[0x30] = cpu.op30
	cpu.opcodes[0x31] = cpu.op31
	cpu.opcodes[0x32] = cpu.op32
	cpu.opcodes[0x33] = cpu.op33
	cpu.opcodes[0x34] = cpu.op34
	cpu.opcodes[0x35] = cpu.op35
	cpu.opcodes[0x36] = cpu.op36
	cpu.opcodes[0x37] = cpu.op37
	cpu.opcodes[0x38] = cpu.op38
	cpu.opcodes[0x39] = cpu.op39
	cpu.opcodes[0x3A] = cpu.op3A
	cpu.opcodes[0x3B] = cpu.op3B
	cpu.opcodes[0x3C] = cpu.op3C
	cpu.opcodes[0x3D] = cpu.op3D
	cpu.opcodes[0x3E] = cpu.op3E
	cpu.opcodes[0x3F] = cpu.op3F
	cpu.opcodes[0x40] = cpu.op40
	cpu.opcodes[0x41] = cpu.op41
	cpu.opcodes[0x42] = cpu.op42
	cpu.opcodes[0x43] = cpu.op43
	cpu.opcodes[0x44] = cpu.op44
	cpu.opcodes[0x45] = cpu.op45
	cpu.opcodes[0x46] = cpu.op46
	cpu.opcodes[0x47] = cpu.op47
	cpu.opcodes[0x48] = cpu.op48
	cpu.opcodes[0x49] = cpu.op49
	cpu.opcodes[0x4A] = cpu.op4A
	cpu.opcodes[0x4B] = cpu.op4B
	cpu.opcodes[0x4C] = cpu.op4C
	cpu.opcodes[0x4D] = cpu.op4D
	cpu.opcodes[0x4E] = cpu.op4E
	cpu.opcodes[0x4F] = cpu.op4F
	cpu.opcodes[0x50] = cpu.op50
	cpu.opcodes[0x51] = cpu.op51
	cpu.opcodes[0x52] = cpu.op52
	cpu.opcodes[0x53] = cpu.op53
	cpu.opcodes[0x54] = cpu.op54
	cpu.opcodes[0x55] = cpu.op55
	cpu.opcodes[0x56] = cpu.op56
	cpu.opcodes[0x57] = cpu.op57
	cpu.opcodes[0x58] = cpu.op58
	cpu.opcodes[0x59] = cpu.op59
	cpu.opcodes[0x5A] = cpu.op5A
	cpu.opcodes[0x5B] = cpu.op5B
	cpu.opcodes[0x5C] = cpu.op5C
	cpu.opcodes[0x5D] = cpu.op5D
	cpu.opcodes[0x5E] = cpu.op5E
	cpu.opcodes[0x5F] = cpu.op5F
	cpu.opcodes[0x60] = cpu.op60
	cpu.opcodes[0x61] = cpu.op61
	cpu.opcodes[0x62] = cpu.op62
	cpu.opcodes[0x63] = cpu.op63
	cpu.opcodes[0x64] = cpu.op64
	cpu.opcodes[0x65] = cpu.op65
	cpu.opcodes[0x66] = cpu.op66
	cpu.opcodes[0x67] = cpu.op67
	cpu.opcodes[0x68] = cpu.op68
	cpu.opcodes[0x69] = cpu.op69
	cpu.opcodes[0x6A] = cpu.op6A
	cpu.opcodes[0x6B] = cpu.op6B
	cpu.opcodes[0x6C] = cpu.op6C
	cpu.opcodes[0x6D] = cpu.op6D
	cpu.opcodes[0x6E] = cpu.op6E
	//cpu.opcodes[0x6F] = cpu.op6F
	cpu.opcodes[0x70] = cpu.op70
	cpu.opcodes[0x71] = cpu.op71
	cpu.opcodes[0x72] = cpu.op72
	cpu.opcodes[0x73] = cpu.op73
	cpu.opcodes[0x74] = cpu.op74
	cpu.opcodes[0x75] = cpu.op75
	cpu.opcodes[0x76] = cpu.op76
	cpu.opcodes[0x77] = cpu.op77
	cpu.opcodes[0x78] = cpu.op78
	cpu.opcodes[0x79] = cpu.op79
	cpu.opcodes[0x7A] = cpu.op7A
	cpu.opcodes[0x7B] = cpu.op7B
	cpu.opcodes[0x7C] = cpu.op7C
	cpu.opcodes[0x7D] = cpu.op7D
	cpu.opcodes[0x7E] = cpu.op7E
	cpu.opcodes[0x7F] = cpu.op7F
	cpu.opcodes[0x80] = cpu.op80
	cpu.opcodes[0x81] = cpu.op81
	cpu.opcodes[0x82] = cpu.op82
	cpu.opcodes[0x83] = cpu.op83
	cpu.opcodes[0x84] = cpu.op84
	cpu.opcodes[0x85] = cpu.op85
	cpu.opcodes[0x86] = cpu.op86
	cpu.opcodes[0x87] = cpu.op87
	cpu.opcodes[0x88] = cpu.op88
	cpu.opcodes[0x89] = cpu.op89
	//cpu.opcodes[0x8A] = cpu.op8A
	cpu.opcodes[0x8B] = cpu.op8B
	cpu.opcodes[0x8C] = cpu.op8C
	cpu.opcodes[0x8D] = cpu.op8D
	cpu.opcodes[0x8E] = cpu.op8E
	cpu.opcodes[0x8F] = cpu.op8F
	cpu.opcodes[0x90] = cpu.op90
	cpu.opcodes[0x91] = cpu.op91
	cpu.opcodes[0x92] = cpu.op92
	cpu.opcodes[0x93] = cpu.op93
	cpu.opcodes[0x94] = cpu.op94
	cpu.opcodes[0x95] = cpu.op95
	cpu.opcodes[0x96] = cpu.op96
	cpu.opcodes[0x97] = cpu.op97
	//cpu.opcodes[0x98] = cpu.op98
	cpu.opcodes[0x99] = cpu.op99
	//cpu.opcodes[0x9A] = cpu.op9A
	//cpu.opcodes[0x9B] = cpu.op9B
	cpu.opcodes[0x9C] = cpu.op9C
	cpu.opcodes[0x9D] = cpu.op9D
	cpu.opcodes[0x9E] = cpu.op9E
	cpu.opcodes[0x9F] = cpu.op9F
	cpu.opcodes[0xA0] = cpu.opA0
	cpu.opcodes[0xA1] = cpu.opA1
	cpu.opcodes[0xA2] = cpu.opA2
	cpu.opcodes[0xA3] = cpu.opA3
	cpu.opcodes[0xA4] = cpu.opA4
	cpu.opcodes[0xA5] = cpu.opA5
	cpu.opcodes[0xA6] = cpu.opA6
	cpu.opcodes[0xA7] = cpu.opA7
	//cpu.opcodes[0xA8] = cpu.opA8
	cpu.opcodes[0xA9] = cpu.opA9
	//cpu.opcodes[0xAA] = cpu.opAA
	cpu.opcodes[0xAB] = cpu.opAB
	cpu.opcodes[0xAC] = cpu.opAC
	cpu.opcodes[0xAD] = cpu.opAD
	cpu.opcodes[0xAE] = cpu.opAE
	cpu.opcodes[0xAF] = cpu.opAF
	cpu.opcodes[0xB0] = cpu.opB0
	cpu.opcodes[0xB1] = cpu.opB1
	cpu.opcodes[0xB2] = cpu.opB2
	cpu.opcodes[0xB3] = cpu.opB3
	cpu.opcodes[0xB4] = cpu.opB4
	cpu.opcodes[0xB5] = cpu.opB5
	cpu.opcodes[0xB6] = cpu.opB6
	cpu.opcodes[0xB7] = cpu.opB7
	cpu.opcodes[0xB8] = cpu.opB8
	cpu.opcodes[0xB9] = cpu.opB9
	//cpu.opcodes[0xBA] = cpu.opBA
	//cpu.opcodes[0xBB] = cpu.opBB
	cpu.opcodes[0xBC] = cpu.opBC
	cpu.opcodes[0xBD] = cpu.opBD
	cpu.opcodes[0xBE] = cpu.opBE
	cpu.opcodes[0xBF] = cpu.opBF
	cpu.opcodes[0xC0] = cpu.opC0
	cpu.opcodes[0xC1] = cpu.opC1
	cpu.opcodes[0xC2] = cpu.opC2
	cpu.opcodes[0xC3] = cpu.opC3
	cpu.opcodes[0xC4] = cpu.opC4
	cpu.opcodes[0xC5] = cpu.opC5
	cpu.opcodes[0xC6] = cpu.opC6
	cpu.opcodes[0xC7] = cpu.opC7
	cpu.opcodes[0xC8] = cpu.opC8
	cpu.opcodes[0xC9] = cpu.opC9
	cpu.opcodes[0xCA] = cpu.opCA
	//cpu.opcodes[0xCB] = cpu.opCB
	cpu.opcodes[0xCC] = cpu.opCC
	cpu.opcodes[0xCD] = cpu.opCD
	cpu.opcodes[0xCE] = cpu.opCE
	cpu.opcodes[0xCF] = cpu.opCF
	cpu.opcodes[0xD0] = cpu.opD0
	cpu.opcodes[0xD1] = cpu.opD1
	cpu.opcodes[0xD2] = cpu.opD2
	cpu.opcodes[0xD3] = cpu.opD3
	cpu.opcodes[0xD4] = cpu.opD4
	cpu.opcodes[0xD5] = cpu.opD5
	cpu.opcodes[0xD6] = cpu.opD6
	cpu.opcodes[0xD7] = cpu.opD7
	cpu.opcodes[0xD8] = cpu.opD8
	cpu.opcodes[0xD9] = cpu.opD9
	cpu.opcodes[0xDA] = cpu.opDA
	//cpu.opcodes[0xDB] = cpu.opDB
	cpu.opcodes[0xDC] = cpu.opDC
	cpu.opcodes[0xDD] = cpu.opDD
	cpu.opcodes[0xDE] = cpu.opDE
	cpu.opcodes[0xDF] = cpu.opDF
	cpu.opcodes[0xE0] = cpu.opE0
	cpu.opcodes[0xE1] = cpu.opE1
	cpu.opcodes[0xE2] = cpu.opE2
	cpu.opcodes[0xE3] = cpu.opE3
	cpu.opcodes[0xE4] = cpu.opE4
	cpu.opcodes[0xE5] = cpu.opE5
	cpu.opcodes[0xE6] = cpu.opE6
	cpu.opcodes[0xE7] = cpu.opE7
	cpu.opcodes[0xE8] = cpu.opE8
	cpu.opcodes[0xE9] = cpu.opE9
	cpu.opcodes[0xEA] = cpu.opEA
	cpu.opcodes[0xEB] = cpu.opEB
	cpu.opcodes[0xEC] = cpu.opEC
	cpu.opcodes[0xED] = cpu.opED
	cpu.opcodes[0xEE] = cpu.opEE
	//cpu.opcodes[0xEF] = cpu.opEF
	cpu.opcodes[0xF0] = cpu.opF0
	cpu.opcodes[0xF1] = cpu.opF1
	cpu.opcodes[0xF2] = cpu.opF2
	cpu.opcodes[0xF3] = cpu.opF3
	cpu.opcodes[0xF4] = cpu.opF4
	cpu.opcodes[0xF5] = cpu.opF5
	cpu.opcodes[0xF6] = cpu.opF6
	cpu.opcodes[0xF7] = cpu.opF7
	cpu.opcodes[0xF8] = cpu.opF8
	cpu.opcodes[0xF9] = cpu.opF9
	cpu.opcodes[0xFA] = cpu.opFA
	cpu.opcodes[0xFB] = cpu.opFB
	cpu.opcodes[0xFC] = cpu.opFC
	cpu.opcodes[0xFD] = cpu.opFD
	cpu.opcodes[0xFE] = cpu.opFE
	cpu.opcodes[0xFF] = cpu.opFF

	return cpu
}

func (cpu *CPU) Init() {
	cpu.reset()
}

func (cpu *CPU) Execute(cycles uint16) {
	for cpu.cycles < cycles {
		opcode := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister())
		fmt.Printf("%x\n", opcode)
		cpu.opcodes[opcode]()
	}
}

func (cpu *CPU) pushStack(data uint8) {
	if cpu.eFlag {
		cpu.memory.SetByteBank(data, 0x00, utils.JoinUint16(0x01, cpu.getSLRegister()))
		cpu.setSLRegister(cpu.getSLRegister() - 1)
	} else {
		cpu.memory.SetByteBank(data, 0x00, cpu.getSRegister())
		cpu.S--
	}
}

func (cpu *CPU) pullStack() uint8 {
	var data uint8
	if cpu.eFlag {
		data = cpu.memory.GetByteBank(0x00, utils.JoinUint16(0x01, cpu.getSLRegister()+1))
		cpu.setSLRegister(cpu.getSLRegister() + 1)
		return data
	}
	data = cpu.memory.GetByteBank(0x00, cpu.getSRegister()+1)
	cpu.S++
	return data
}

func (cpu *CPU) pushStackNew8(data uint8) {
	cpu.pushStackNew(data)
	if cpu.eFlag {
		cpu.setSHRegister(0x01)
	}
}

func (cpu *CPU) pushStackNew16(dataHi, dataLo uint8) {
	cpu.pushStackNew(dataHi)
	cpu.pushStackNew(dataLo)
	if cpu.eFlag {
		cpu.setSHRegister(0x01)
	}
}

func (cpu *CPU) pushStackNew24(dataHi, dataMid, dataLo uint8) {
	cpu.pushStackNew(dataHi)
	cpu.pushStackNew(dataMid)
	cpu.pushStackNew(dataLo)
	if cpu.eFlag {
		cpu.setSHRegister(0x01)
	}
}

func (cpu *CPU) pullStackNew8() (data uint8) {
	data = cpu.pullStackNew()
	if cpu.eFlag {
		cpu.setSHRegister(0x01)
	}
	return
}

func (cpu *CPU) pullStackNew16() (dataHi, dataLo uint8) {
	dataLo = cpu.pullStackNew()
	dataHi = cpu.pullStackNew()
	if cpu.eFlag {
		cpu.setSHRegister(0x01)
	}
	return
}

func (cpu *CPU) pullStackNew24() (dataHi, dataMid, dataLo uint8) {
	dataLo = cpu.pullStackNew()
	dataMid = cpu.pullStackNew()
	dataHi = cpu.pullStackNew()
	if cpu.eFlag {
		cpu.setSHRegister(0x01)
	}
	return
}

func (cpu *CPU) pushStackNew(data uint8) {
	cpu.memory.SetByteBank(data, 0x00, cpu.getSRegister())
	cpu.S--
}

func (cpu *CPU) pullStackNew() uint8 {
	var data uint8
	data = cpu.memory.GetByteBank(0x00, cpu.getSRegister()+1)
	cpu.S++
	return data
}
