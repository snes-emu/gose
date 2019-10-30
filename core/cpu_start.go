package core

import (
	"encoding/json"
	"fmt"

	"github.com/snes-emu/gose/log"

	"github.com/snes-emu/gose/config"
)

type instruction struct {
	name    string
	opcodes []uint8
}

type cpuState struct {
	C           uint16
	DBR         uint8
	D           uint16
	K           uint8
	PC          uint16
	S           uint16
	X           uint16
	Y           uint16
	Flags       string
	Instruction string
}

func init() {
	for _, instruction := range instructions {
		for _, opcode := range instruction.opcodes {
			opcodeToInstruction[opcode] = instruction
		}
	}
}

func (cpu *CPU) execOpcode() {
	K := cpu.getKRegister()
	PC := cpu.getPCRegister()
	opcode := cpu.memory.GetByteBank(K, PC)
	cpu.logState(K, PC, opcode)
	cpu.opcodes[opcode]()
}

var (
	instructions = []instruction{
		instruction{
			name:    "ADC",
			opcodes: []uint8{0x61, 0x63, 0x65, 0x67, 0x69, 0x6D, 0x6F, 0x71, 0x72, 0x73, 0x75, 0x77, 0x79, 0x7D, 0x7F},
		},
		instruction{
			name:    "SBC",
			opcodes: []uint8{0xE1, 0xE3, 0xE5, 0xE7, 0xE9, 0xED, 0xEF, 0xF1, 0xF2, 0xF3, 0xF5, 0xF7, 0xF9, 0xFD, 0xFF},
		},
		instruction{
			name:    "CMP",
			opcodes: []uint8{0xC1, 0xC3, 0xC5, 0xC7, 0xC9, 0xCD, 0xCF, 0xD1, 0xD2, 0xD3, 0xD5, 0xD7, 0xD9, 0xDD, 0xDF},
		},
		instruction{
			name:    "CPX",
			opcodes: []uint8{0xE0, 0xE4, 0xEC},
		},
		instruction{
			name:    "CPY",
			opcodes: []uint8{0xC0, 0xC4, 0xCC},
		},
		instruction{
			name:    "DEC",
			opcodes: []uint8{0x3A, 0xC6, 0xCE, 0xD6, 0xDE},
		},
		instruction{
			name:    "DEX",
			opcodes: []uint8{0xCA},
		},
		instruction{
			name:    "DEY",
			opcodes: []uint8{0x88},
		},
		instruction{
			name:    "INC",
			opcodes: []uint8{0x1A, 0xE6, 0xEE, 0xF6, 0xFE},
		},
		instruction{
			name:    "INX",
			opcodes: []uint8{0xE8},
		},
		instruction{
			name:    "INY",
			opcodes: []uint8{0xC8},
		},
		instruction{
			name:    "AND",
			opcodes: []uint8{0x21, 0x23, 0x25, 0x27, 0x29, 0x2D, 0x2F, 0x31, 0x32, 0x33, 0x35, 0x37, 0x39, 0x3D, 0x3F},
		},
		instruction{
			name:    "EOR",
			opcodes: []uint8{0x41, 0x43, 0x45, 0x47, 0x49, 0x4D, 0x4F, 0x51, 0x52, 0x53, 0x55, 0x57, 0x59, 0x5D, 0x5F},
		},
		instruction{
			name:    "ORA",
			opcodes: []uint8{0x01, 0x03, 0x05, 0x07, 0x09, 0x0D, 0x0F, 0x11, 0x12, 0x13, 0x15, 0x17, 0x19, 0x1D, 0x1F},
		},
		instruction{
			name:    "BIT",
			opcodes: []uint8{0x24, 0x2C, 0x34, 0x3C, 0x89},
		},
		instruction{
			name:    "TRB",
			opcodes: []uint8{0x14, 0x1C},
		},
		instruction{
			name:    "TSB",
			opcodes: []uint8{0x04, 0x0C},
		},
		instruction{
			name:    "ASL",
			opcodes: []uint8{0x06, 0x0A, 0x0E, 0x16, 0x1E},
		},
		instruction{
			name:    "LSR",
			opcodes: []uint8{0x46, 0x4A, 0x4E, 0x56, 0x5E},
		},
		instruction{
			name:    "ROL",
			opcodes: []uint8{0x26, 0x2A, 0x2E, 0x36, 0x3E},
		},
		instruction{
			name:    "ROR",
			opcodes: []uint8{0x66, 0x6A, 0x6E, 0x76, 0x7E},
		},
		instruction{
			name:    "BCC",
			opcodes: []uint8{0x90},
		},
		instruction{
			name:    "BCS",
			opcodes: []uint8{0xB0},
		},
		instruction{
			name:    "BEQ",
			opcodes: []uint8{0xF0},
		},
		instruction{
			name:    "BMI",
			opcodes: []uint8{0x30},
		},
		instruction{
			name:    "BNE",
			opcodes: []uint8{0xD0},
		},
		instruction{
			name:    "BPL",
			opcodes: []uint8{0x10},
		},
		instruction{
			name:    "BRA",
			opcodes: []uint8{0x80},
		},
		instruction{
			name:    "BVC",
			opcodes: []uint8{0x50},
		},
		instruction{
			name:    "BVS",
			opcodes: []uint8{0x70},
		},
		instruction{
			name:    "BRL",
			opcodes: []uint8{0x82},
		},
		instruction{
			name:    "JMP",
			opcodes: []uint8{0x4C, 0x5C, 0x6C, 0x7C, 0xDC},
		},
		instruction{
			name:    "JSL",
			opcodes: []uint8{0x22},
		},
		instruction{
			name:    "JSR",
			opcodes: []uint8{0x20, 0xFC},
		},
		instruction{
			name:    "RTL",
			opcodes: []uint8{0x6B},
		},
		instruction{
			name:    "RTS",
			opcodes: []uint8{0x60},
		},
		instruction{
			name:    "BRK",
			opcodes: []uint8{0x00},
		},
		instruction{
			name:    "COP",
			opcodes: []uint8{0x02},
		},
		instruction{
			name:    "RTI",
			opcodes: []uint8{0x40},
		},
		instruction{
			name:    "CLC",
			opcodes: []uint8{0x18},
		},
		instruction{
			name:    "CLD",
			opcodes: []uint8{0xD8},
		},
		instruction{
			name:    "CLI",
			opcodes: []uint8{0x58},
		},
		instruction{
			name:    "CLV",
			opcodes: []uint8{0xB8},
		},
		instruction{
			name:    "SEC",
			opcodes: []uint8{0x38},
		},
		instruction{
			name:    "SED",
			opcodes: []uint8{0xF8},
		},
		instruction{
			name:    "SEI",
			opcodes: []uint8{0x78},
		},
		instruction{
			name:    "REP",
			opcodes: []uint8{0xC2},
		},
		instruction{
			name:    "SEP",
			opcodes: []uint8{0xE2},
		},
		instruction{
			name:    "LDA",
			opcodes: []uint8{0xA1, 0xA3, 0xA5, 0xA7, 0xA9, 0xAD, 0xAF, 0xB1, 0xB2, 0xB3, 0xB5, 0xB7, 0xB9, 0xBD, 0xBF},
		},
		instruction{
			name:    "LDX",
			opcodes: []uint8{0xA2, 0xA6, 0xAE, 0xB6, 0xBE},
		},
		instruction{
			name:    "LDY",
			opcodes: []uint8{0xA0, 0xA4, 0xAC, 0xB4, 0xBC},
		},
		instruction{
			name:    "STA",
			opcodes: []uint8{0x81, 0x83, 0x85, 0x87, 0x8D, 0x8F, 0x91, 0x92, 0x93, 0x95, 0x97, 0x99, 0x9D, 0x9F},
		},
		instruction{
			name:    "STX",
			opcodes: []uint8{0x86, 0x8E, 0x96},
		},
		instruction{
			name:    "STY",
			opcodes: []uint8{0x84, 0x8C, 0x94},
		},
		instruction{
			name:    "STZ",
			opcodes: []uint8{0x64, 0x74, 0x9C, 0x9E},
		},
		instruction{
			name:    "MVN",
			opcodes: []uint8{0x54},
		},
		instruction{
			name:    "MVP",
			opcodes: []uint8{0x44},
		},
		instruction{
			name:    "NOP",
			opcodes: []uint8{0xEA},
		},
		instruction{
			name:    "WDM",
			opcodes: []uint8{0x42},
		},
		instruction{
			name:    "PEA",
			opcodes: []uint8{0xF4},
		},
		instruction{
			name:    "PEI",
			opcodes: []uint8{0xD4},
		},
		instruction{
			name:    "PER",
			opcodes: []uint8{0x62},
		},
		instruction{
			name:    "PHA",
			opcodes: []uint8{0x48},
		},
		instruction{
			name:    "PHX",
			opcodes: []uint8{0xDA},
		},
		instruction{
			name:    "PHY",
			opcodes: []uint8{0x5A},
		},
		instruction{
			name:    "PLA",
			opcodes: []uint8{0x68},
		},
		instruction{
			name:    "PLX",
			opcodes: []uint8{0xFA},
		},
		instruction{
			name:    "PLY",
			opcodes: []uint8{0x7A},
		},
		instruction{
			name:    "PHB",
			opcodes: []uint8{0x8B},
		},
		instruction{
			name:    "PHD",
			opcodes: []uint8{0x0B},
		},
		instruction{
			name:    "PHK",
			opcodes: []uint8{0x4B},
		},
		instruction{
			name:    "PHP",
			opcodes: []uint8{0x08},
		},
		instruction{
			name:    "PLB",
			opcodes: []uint8{0xAB},
		},
		instruction{
			name:    "PLD",
			opcodes: []uint8{0x2B},
		},
		instruction{
			name:    "PLP",
			opcodes: []uint8{0x28},
		},
		instruction{
			name:    "STP",
			opcodes: []uint8{0xDB},
		},
		instruction{
			name:    "WAI",
			opcodes: []uint8{0xCB},
		},
		instruction{
			name:    "TAX",
			opcodes: []uint8{0xAA},
		},
		instruction{
			name:    "TAY",
			opcodes: []uint8{0xA8},
		},
		instruction{
			name:    "TSX",
			opcodes: []uint8{0xBA},
		},
		instruction{
			name:    "TXA",
			opcodes: []uint8{0x8A},
		},
		instruction{
			name:    "TXS",
			opcodes: []uint8{0x9A},
		},
		instruction{
			name:    "TXY",
			opcodes: []uint8{0x9B},
		},
		instruction{
			name:    "TYA",
			opcodes: []uint8{0x98},
		},
		instruction{
			name:    "TYX",
			opcodes: []uint8{0xBB},
		},
		instruction{
			name:    "TCD",
			opcodes: []uint8{0x5B},
		},
		instruction{
			name:    "TCS",
			opcodes: []uint8{0x1B},
		},
		instruction{
			name:    "TDC",
			opcodes: []uint8{0x7B},
		},
		instruction{
			name:    "TSC",
			opcodes: []uint8{0x3B},
		},
		instruction{
			name:    "XBA",
			opcodes: []uint8{0xEB},
		},
		instruction{
			name:    "XCE",
			opcodes: []uint8{0xFB},
		},
	}

	opcodeToInstruction = make([]instruction, 0x100)
)

func (cpu *CPU) prettyFlags() string {
	PString := ""
	if cpu.eFlag {
		PString += "E"
	} else {
		PString += "e"
	}
	if cpu.nFlag {
		PString += "N"
	} else {
		PString += "n"
	}
	if cpu.vFlag {
		PString += "V"
	} else {
		PString += "v"
	}
	if cpu.mFlag {
		PString += "M"
	} else {
		PString += "m"
	}
	if cpu.xFlag {
		PString += "X"
	} else {
		PString += "x"
	}
	if cpu.dFlag {
		PString += "D"
	} else {
		PString += "d"
	}
	if cpu.iFlag {
		PString += "I"
	} else {
		PString += "i"
	}
	if cpu.zFlag {
		PString += "Z"
	} else {
		PString += "z"
	}
	if cpu.cFlag {
		PString += "C"
	} else {
		PString += "c"
	}
	return PString
}

func (cpu *CPU) MarshalJSON() ([]byte, error) {
	K := cpu.getKRegister()
	PC := cpu.getPCRegister()

	opcode := cpu.memory.GetByteBank(K, PC)

	return json.Marshal(cpuState{
		K:           K,
		PC:          PC,
		Instruction: opcodeToInstruction[opcode].name,
		C:           cpu.getCRegister(),
		DBR:         cpu.getDBRRegister(),
		D:           cpu.getDRegister(),
	})
}

func (cpu *CPU) logState(K uint8, PC uint16, opcode uint8) {
	if !config.DebugServer() {
		return
	}
	instruction := opcodeToInstruction[opcode]

	log.Debug(fmt.Sprintf("$%02X:%04X %02X %s A:%04X X:%04X Y:%04X D:%04X DB:%02X S:%04X P:%s", K, PC, opcode, instruction.name, cpu.getCRegister(), cpu.getXRegister(), cpu.getYRegister(), cpu.getDRegister(), cpu.getDBRRegister(), cpu.getSRegister(), cpu.prettyFlags()))
}
