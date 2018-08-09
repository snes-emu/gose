// +build debug

package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type instruction struct {
	name string
}

var (
	instructions = map[string]instruction{
		"ADC": instruction{
			name: "ADC",
		},
		"SBC": instruction{
			name: "SBC",
		},
		"CMP": instruction{
			name: "CMP",
		},
		"CPX": instruction{
			name: "CPX",
		},
		"CPY": instruction{
			name: "CPY",
		},
		"DEC": instruction{
			name: "DEC",
		},
		"DEX": instruction{
			name: "DEX",
		},
		"DEY": instruction{
			name: "DEY",
		},
		"INC": instruction{
			name: "INC",
		},
		"INX": instruction{
			name: "INX",
		},
		"INY": instruction{
			name: "INY",
		},
		"AND": instruction{
			name: "AND",
		},
		"EOR": instruction{
			name: "EOR",
		},
		"ORA": instruction{
			name: "ORA",
		},
		"BIT": instruction{
			name: "BIT",
		},
		"TRB": instruction{
			name: "TRB",
		},
		"TSB": instruction{
			name: "TSB",
		},
		"ASL": instruction{
			name: "ASL",
		},
		"LSR": instruction{
			name: "LSR",
		},
		"ROL": instruction{
			name: "ROL",
		},
		"ROR": instruction{
			name: "ROR",
		},
		"BCC": instruction{
			name: "BCC",
		},
		"BCS": instruction{
			name: "BCS",
		},
		"BEQ": instruction{
			name: "BEQ",
		},
		"BMI": instruction{
			name: "BMI",
		},
		"BNE": instruction{
			name: "BNE",
		},
		"BPL": instruction{
			name: "BPL",
		},
		"BRA": instruction{
			name: "BRA",
		},
		"BVC": instruction{
			name: "BVC",
		},
		"BVS": instruction{
			name: "BVS",
		},
		"BRL": instruction{
			name: "BRL",
		},
		"JMP": instruction{
			name: "JMP",
		},
		"JSL": instruction{
			name: "JSL",
		},
		"JSR": instruction{
			name: "JSR",
		},
		"RTL": instruction{
			name: "RTL",
		},
		"RTS": instruction{
			name: "RTS",
		},
		"BRK": instruction{
			name: "BRK",
		},
		"COP": instruction{
			name: "COP",
		},
		"RTI": instruction{
			name: "RTI",
		},
		"CLC": instruction{
			name: "CLC",
		},
		"CLD": instruction{
			name: "CLD",
		},
		"CLI": instruction{
			name: "CLI",
		},
		"CLV": instruction{
			name: "CLV",
		},
		"SEC": instruction{
			name: "SEC",
		},
		"SED": instruction{
			name: "SED",
		},
		"SEI": instruction{
			name: "SEI",
		},
		"REP": instruction{
			name: "REP",
		},
		"SEP": instruction{
			name: "SEP",
		},
		"LDA": instruction{
			name: "LDA",
		},
		"LDX": instruction{
			name: "LDX",
		},
		"LDY": instruction{
			name: "LDY",
		},
		"STA": instruction{
			name: "STA",
		},
		"STX": instruction{
			name: "STX",
		},
		"STY": instruction{
			name: "STY",
		},
		"STZ": instruction{
			name: "STZ",
		},
		"MVN": instruction{
			name: "MVN",
		},
		"MVP": instruction{
			name: "MVP",
		},
		"NOP": instruction{
			name: "NOP",
		},
		"WDM": instruction{
			name: "WDM",
		},
		"PEA": instruction{
			name: "PEA",
		},
		"PEI": instruction{
			name: "PEI",
		},
		"PER": instruction{
			name: "PER",
		},
		"PHA": instruction{
			name: "PHA",
		},
		"PHX": instruction{
			name: "PHX",
		},
		"PHY": instruction{
			name: "PHY",
		},
		"PLA": instruction{
			name: "PLA",
		},
		"PLX": instruction{
			name: "PLX",
		},
		"PLY": instruction{
			name: "PLY",
		},
		"PHB": instruction{
			name: "PHB",
		},
		"PHD": instruction{
			name: "PHD",
		},
		"PHK": instruction{
			name: "PHK",
		},
		"PHP": instruction{
			name: "PHP",
		},
		"PLB": instruction{
			name: "PLB",
		},
		"PLD": instruction{
			name: "PLD",
		},
		"PLP": instruction{
			name: "PLP",
		},
		"STP": instruction{
			name: "STP",
		},
		"WAI": instruction{
			name: "WAI",
		},
		"TAX": instruction{
			name: "TAX",
		},
		"TAY": instruction{
			name: "TAY",
		},
		"TSX": instruction{
			name: "TSX",
		},
		"TXA": instruction{
			name: "TXA",
		},
		"TXS": instruction{
			name: "TXS",
		},
		"TXY": instruction{
			name: "TXY",
		},
		"TYA": instruction{
			name: "TYA",
		},
		"TYX": instruction{
			name: "TYX",
		},
		"TCD": instruction{
			name: "TCD",
		},
		"TCS": instruction{
			name: "TCS",
		},
		"TDC": instruction{
			name: "TDC",
		},
		"TSC": instruction{
			name: "TSC",
		},
		"XBA": instruction{
			name: "XBA",
		},
		"XCE": instruction{
			name: "XCE",
		},
	}

	opcodeToInstruction = []instruction{
		instructions["BRK"],
		instructions["ORA"],
		instructions["COP"],
		instructions["ORA"],
		instructions["TSB"],
		instructions["ORA"],
		instructions["ASL"],
		instructions["ORA"],
		instructions["PHP"],
		instructions["ORA"],
		instructions["ASL"],
		instructions["PHD"],
		instructions["TSB"],
		instructions["ORA"],
		instructions["ASL"],
		instructions["ORA"],
		instructions["BPL"],
		instructions["ORA"],
		instructions["ORA"],
		instructions["ORA"],
		instructions["TRB"],
		instructions["ORA"],
		instructions["ASL"],
		instructions["ORA"],
		instructions["CLC"],
		instructions["ORA"],
		instructions["INC"],
		instructions["TCS"],
		instructions["TRB"],
		instructions["ORA"],
		instructions["ASL"],
		instructions["ORA"],
		instructions["JSR"],
		instructions["AND"],
		instructions["JSL"],
		instructions["AND"],
		instructions["BIT"],
		instructions["AND"],
		instructions["ROL"],
		instructions["AND"],
		instructions["PLP"],
		instructions["AND"],
		instructions["ROL"],
		instructions["PLD"],
		instructions["BIT"],
		instructions["AND"],
		instructions["ROL"],
		instructions["AND"],
		instructions["BMI"],
		instructions["AND"],
		instructions["AND"],
		instructions["AND"],
		instructions["BIT"],
		instructions["AND"],
		instructions["ROL"],
		instructions["AND"],
		instructions["SEC"],
		instructions["AND"],
		instructions["DEC"],
		instructions["TSC"],
		instructions["BIT"],
		instructions["AND"],
		instructions["ROL"],
		instructions["AND"],
		instructions["RTI"],
		instructions["EOR"],
		instructions["WDM"],
		instructions["EOR"],
		instructions["MVP"],
		instructions["EOR"],
		instructions["LSR"],
		instructions["EOR"],
		instructions["PHA"],
		instructions["EOR"],
		instructions["LSR"],
		instructions["PHK"],
		instructions["JMP"],
		instructions["EOR"],
		instructions["LSR"],
		instructions["EOR"],
		instructions["BVC"],
		instructions["EOR"],
		instructions["EOR"],
		instructions["EOR"],
		instructions["MVN"],
		instructions["EOR"],
		instructions["LSR"],
		instructions["EOR"],
		instructions["CLI"],
		instructions["EOR"],
		instructions["PHY"],
		instructions["TCD"],
		instructions["JMP"],
		instructions["EOR"],
		instructions["LSR"],
		instructions["EOR"],
		instructions["RTS"],
		instructions["ADC"],
		instructions["PER"],
		instructions["ADC"],
		instructions["STZ"],
		instructions["ADC"],
		instructions["ROR"],
		instructions["ADC"],
		instructions["PLA"],
		instructions["ADC"],
		instructions["ROR"],
		instructions["RTL"],
		instructions["JMP"],
		instructions["ADC"],
		instructions["ROR"],
		instructions["ADC"],
		instructions["BVS"],
		instructions["ADC"],
		instructions["ADC"],
		instructions["ADC"],
		instructions["STZ"],
		instructions["ADC"],
		instructions["ROR"],
		instructions["ADC"],
		instructions["SEI"],
		instructions["ADC"],
		instructions["PLY"],
		instructions["TDC"],
		instructions["JMP"],
		instructions["ADC"],
		instructions["ROR"],
		instructions["ADC"],
		instructions["BRA"],
		instructions["STA"],
		instructions["BRL"],
		instructions["STA"],
		instructions["STY"],
		instructions["STA"],
		instructions["STX"],
		instructions["STA"],
		instructions["DEY"],
		instructions["BIT"],
		instructions["TXA"],
		instructions["PHB"],
		instructions["STY"],
		instructions["STA"],
		instructions["STX"],
		instructions["STA"],
		instructions["BCC"],
		instructions["STA"],
		instructions["STA"],
		instructions["STA"],
		instructions["STY"],
		instructions["STA"],
		instructions["STX"],
		instructions["STA"],
		instructions["TYA"],
		instructions["STA"],
		instructions["TXS"],
		instructions["TXY"],
		instructions["STZ"],
		instructions["STA"],
		instructions["STZ"],
		instructions["STA"],
		instructions["LDY"],
		instructions["LDA"],
		instructions["LDX"],
		instructions["LDA"],
		instructions["LDY"],
		instructions["LDA"],
		instructions["LDX"],
		instructions["LDA"],
		instructions["TAY"],
		instructions["LDA"],
		instructions["TAX"],
		instructions["PLB"],
		instructions["LDY"],
		instructions["LDA"],
		instructions["LDX"],
		instructions["LDA"],
		instructions["BCS"],
		instructions["LDA"],
		instructions["LDA"],
		instructions["LDA"],
		instructions["LDY"],
		instructions["LDA"],
		instructions["LDX"],
		instructions["LDA"],
		instructions["CLV"],
		instructions["LDA"],
		instructions["TSX"],
		instructions["TYX"],
		instructions["LDY"],
		instructions["LDA"],
		instructions["LDX"],
		instructions["LDA"],
		instructions["CPY"],
		instructions["CMP"],
		instructions["REP"],
		instructions["CMP"],
		instructions["CPY"],
		instructions["CMP"],
		instructions["DEC"],
		instructions["CMP"],
		instructions["INY"],
		instructions["CMP"],
		instructions["DEX"],
		instructions["WAI"],
		instructions["CPY"],
		instructions["CMP"],
		instructions["DEC"],
		instructions["CMP"],
		instructions["BNE"],
		instructions["CMP"],
		instructions["CMP"],
		instructions["CMP"],
		instructions["PEI"],
		instructions["CMP"],
		instructions["DEC"],
		instructions["CMP"],
		instructions["CLD"],
		instructions["CMP"],
		instructions["PHX"],
		instructions["STP"],
		instructions["JMP"],
		instructions["CMP"],
		instructions["DEC"],
		instructions["CMP"],
		instructions["CPX"],
		instructions["SBC"],
		instructions["SEP"],
		instructions["SBC"],
		instructions["CPX"],
		instructions["SBC"],
		instructions["INC"],
		instructions["SBC"],
		instructions["INX"],
		instructions["SBC"],
		instructions["NOP"],
		instructions["XBA"],
		instructions["CPX"],
		instructions["SBC"],
		instructions["INC"],
		instructions["SBC"],
		instructions["BEQ"],
		instructions["SBC"],
		instructions["SBC"],
		instructions["SBC"],
		instructions["PEA"],
		instructions["SBC"],
		instructions["INC"],
		instructions["SBC"],
		instructions["SED"],
		instructions["SBC"],
		instructions["PLX"],
		instructions["XCE"],
		instructions["JSR"],
		instructions["SBC"],
		instructions["INC"],
		instructions["SBC"],
	}
)

func (cpu *CPU) StartDebug() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigs:
			fmt.Printf("Emulator exited")
			os.Exit(0)
		default:
			K := cpu.getKRegister()
			PC := cpu.getPCRegister()
			opcode := cpu.memory.GetByteBank(K, PC)
			cpu.logState(K, PC, opcode)
			cpu.opcodes[opcode]()
		}
	}
}

func (cpu *CPU) logState(K uint8, PC uint16, opcode uint8) {
	instruction := opcodeToInstruction[opcode]
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
	fmt.Printf("$%02X:%04X %02X %s A:%04X X:%04X Y:%04X D:%04X DB:%02X S:%04X P:%s\n", K, PC, opcode, instruction.name, cpu.getCRegister(), cpu.getXRegister(), cpu.getYRegister(), cpu.getDRegister(), cpu.getDBRRegister(), cpu.getSRegister(), PString)
}
