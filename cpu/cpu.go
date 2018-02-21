package cpu

import (
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
	opcodes []cpuOperation
}

type cpuOperation func()

var opcodes []cpuOperation

func New() *CPU {
	cpu := &CPU{}
	cpu.memory = memory.New()
	// cpu.opcodes[0x61] = cpu.op61
	return cpu
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
