package cpu

import (
	"github.com/snes-emu/gose/utils"
)

//p16 pushes the next 16-bit value into the stack
func (cpu *CPU) p16(dataLo, dataHi uint8) {
	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
}

//p8 pushes the next 8-bit value into the stack
func (cpu *CPU) p8(data uint8) {
	cpu.pushStack(data)
}

// PEA instruction
func (cpu *CPU) opF4() {
	dataLo, dataHi := cpu.admImmediate16()
	cpu.pushStackNew16(dataLo, dataHi)
	cpu.cycles += 5
	cpu.PC += 3
}

// PEI instruction
func (cpu *CPU) opD4() {
	dataLo, dataHi := cpu.admDirectNew()
	cpu.pushStackNew16(dataLo, dataHi)
	cpu.cycles += 6 + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

// PER instuction
func (cpu *CPU) op62() {
	dataLo, dataHi := cpu.admImmediate16()
	cpu.pushStackNew16(dataLo, dataHi)
	cpu.cycles += 6
	cpu.PC += 3
}

// pha16 push the accumulator onto the stack
func (cpu *CPU) pha16() {
	dataLo, dataHi := utils.SplitUint16(cpu.getCRegister())
	cpu.p16(dataLo, dataHi)
}

// pha8 push the lower bit of the accumulator onto the stack
func (cpu *CPU) pha8() {
	cpu.p8(cpu.getARegister())
}

func (cpu *CPU) pha() {
	if cpu.mFlag {
		cpu.pha8()
	} else {
		cpu.pha16()
	}
}

func (cpu *CPU) op48() {
	cpu.pha()
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC++
}

// PHB instruction
func (cpu *CPU) op8B() {
	cpu.pushStackNew8(cpu.getDBRRegister())
	cpu.cycles += 3
	cpu.PC++
}

// PHD instruction
func (cpu *CPU) op0B() {
	cpu.pushStackNew16(utils.SplitUint16(cpu.getDRegister()))
	cpu.cycles += 4
	cpu.PC++
}

// PHK instruction
func (cpu *CPU) op4B() {
	cpu.pushStackNew8(cpu.getKRegister())
	cpu.cycles += 3
	cpu.PC++
}

func (cpu *CPU) php() {
	P := utils.BoolToUint8[cpu.cFlag]*0x01 +
		utils.BoolToUint8[cpu.zFlag]*0x02 +
		utils.BoolToUint8[cpu.iFlag]*0x04 +
		utils.BoolToUint8[cpu.dFlag]*0x08 +
		utils.BoolToUint8[cpu.mFlag]*0x20 +
		utils.BoolToUint8[cpu.vFlag]*0x40 +
		utils.BoolToUint8[cpu.nFlag]*0x80
	if cpu.eFlag {
		P += utils.BoolToUint8[cpu.bFlag] * 0x10
	} else {
		P += utils.BoolToUint8[cpu.xFlag] * 0x10
	}
	cpu.pushStack(P)
}

func (cpu *CPU) op08() {
	cpu.php()
	cpu.cycles += 3
	cpu.PC++
}

// phx16 push the X register onto the stack
func (cpu *CPU) phx16() {
	dataLo, dataHi := utils.SplitUint16(cpu.getXRegister())
	cpu.p16(dataLo, dataHi)
}

// phx8 push the lower bit of the X register onto the stack
func (cpu *CPU) phx8() {
	cpu.p8(cpu.getXLRegister())
}

func (cpu *CPU) phx() {
	if cpu.xFlag {
		cpu.phx8()
	} else {
		cpu.phx16()
	}
}

func (cpu *CPU) opDA() {
	cpu.phx()
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC++
}

// phy16 push the Y register onto the stack
func (cpu *CPU) phy16() {
	dataLo, dataHi := utils.SplitUint16(cpu.getYRegister())
	cpu.p16(dataLo, dataHi)
}

// phy8 push the lower bit of the Y register onto the stack
func (cpu *CPU) phy8() {
	cpu.p8(cpu.getYLRegister())
}

func (cpu *CPU) phy() {
	if cpu.xFlag {
		cpu.phy8()
	} else {
		cpu.phy16()
	}
}

func (cpu *CPU) op5A() {
	cpu.phy()
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC++
}

// pla16 pull the accumulator from the stack
func (cpu *CPU) pla16() {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()

	result := utils.JoinUint16(dataLo, dataHi)

	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// pla8 pull the lower bits of the accumulator from the stack
func (cpu *CPU) pla8() {
	result := cpu.pullStack()

	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

func (cpu *CPU) pla() {
	if cpu.mFlag {
		cpu.pla8()
	} else {
		cpu.pla16()
	}
}

func (cpu *CPU) op68() {
	cpu.pla()
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC++
}

// PLB instruction
func (cpu *CPU) opAB() {
	cpu.DBR = cpu.pullStackNew8()
	cpu.nFlag = cpu.getDBRRegister()&0x80 != 0
	cpu.zFlag = cpu.getDBRRegister() == 0
	cpu.cycles += 4
	cpu.PC++
}

// PLD instruction
func (cpu *CPU) op2B() {
	cpu.D = utils.JoinUint16(cpu.pullStackNew16())
	cpu.nFlag = cpu.getDRegister()&0x80 != 0
	cpu.zFlag = cpu.getDRegister() == 0
	cpu.cycles += 4
	cpu.PC++
}

func (cpu *CPU) plp() {
	P := cpu.pullStack()
	cpu.cFlag = P&0x01 != 0
	cpu.zFlag = P&0x02 != 0
	cpu.iFlag = P&0x04 != 0
	cpu.dFlag = P&0x08 != 0
	cpu.mFlag = P&0x20 != 0
	cpu.vFlag = P&0x40 != 0
	cpu.nFlag = P&0x80 != 0
	if cpu.eFlag {
		cpu.bFlag = P&0x10 != 0
		cpu.setXFlag(true)
		cpu.mFlag = true
	} else {
		cpu.setXFlag(P&0x10 != 0)
	}
}

func (cpu *CPU) op28() {
	cpu.plp()
	cpu.cycles += 4
	cpu.PC++
}

// plx16 pull the X register from the stack
func (cpu *CPU) plx16() {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()

	result := utils.JoinUint16(dataLo, dataHi)

	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setXRegister(result)
}

// plx8 pull the lower bits of the X register from the stack
func (cpu *CPU) plx8() {
	result := cpu.pullStack()

	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setXLRegister(result)
}

func (cpu *CPU) plx() {
	if cpu.xFlag {
		cpu.plx8()
	} else {
		cpu.plx16()
	}
}

func (cpu *CPU) opFA() {
	cpu.plx()
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC++
}

// ply16 pull the Y register from the stack
func (cpu *CPU) ply16() {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()

	result := utils.JoinUint16(dataLo, dataHi)

	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setYRegister(result)
}

// ply8 pull the lower bits of the Y register from the stack
func (cpu *CPU) ply8() {
	result := cpu.pullStack()

	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setYLRegister(result)
}

func (cpu *CPU) ply() {
	if cpu.xFlag {
		cpu.ply8()
	} else {
		cpu.ply16()
	}
}

func (cpu *CPU) op7A() {
	cpu.ply()
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC++
}
