package cpu

import (
	"github.com/snes-emu/gose/utils"
)

// lda16 load data into the accumulator
func (cpu *CPU) lda16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data == 0
	cpu.setCRegister(data)
}

// lda8 load data into the lower bits of the accumulator
func (cpu *CPU) lda8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data == 0

	cpu.setARegister(data)
}

// lda load data into the accumulator
func (cpu *CPU) lda(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.lda8(dataLo)
	} else {
		cpu.lda16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA1() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA3() {
	dataLo, dataHi := cpu.admStackS()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opA5() {
	dataLo, dataHi := cpu.admDirect()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA7() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA9() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opAD() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) opAF() {
	dataLo, dataHi := cpu.admLong()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) opB1() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) opB2() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB3() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opB5() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB7() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB9() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opBD() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opBF() {
	dataLo, dataHi := cpu.admLongX()
	cpu.lda(dataLo, dataHi)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

// ldx16 load data into the x register
func (cpu *CPU) ldx16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data == 0

	cpu.setXRegister(data)
}

// ldx8 load data into the lower bits of the x register
func (cpu *CPU) ldx8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data == 0

	cpu.setXLRegister(data)
}

// ldx load data into the x register taking care of 16bit/8bit cases
func (cpu *CPU) ldx(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.ldx8(dataLo)
	} else {
		cpu.ldx16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA2() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.ldx(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opA6() {
	dataLo, dataHi := cpu.admDirect()
	cpu.ldx(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opAE() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.ldx(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) opB6() {
	dataLo, dataHi := cpu.admDirectY()
	cpu.ldx(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opBE() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.ldx(dataLo, dataHi)
	cpu.cycles += 6 - 2*utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

// ldy16 load data into the y register
func (cpu *CPU) ldy16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data == 0

	cpu.setYRegister(data)
}

// ldy8 load data into the lower bits of the y register
func (cpu *CPU) ldy8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data == 0

	cpu.setYLRegister(data)
}

// ldy load data into the y register taking care of 16bit/8bit cases
func (cpu *CPU) ldy(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.ldy8(dataLo)
	} else {
		cpu.ldy16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA0() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.ldy(dataLo, dataHi)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opA4() {
	dataLo, dataHi := cpu.admDirect()
	cpu.ldy(dataLo, dataHi)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opAC() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.ldy(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) opB4() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.ldy(dataLo, dataHi)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opBC() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.ldy(dataLo, dataHi)
	cpu.cycles += 6 - 2*utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

// sta16 stores the accumulator in the memory
func (cpu *CPU) sta16(laddr, haddr uint32) {

	dataLo, dataHi := utils.SplitUint16(cpu.getCRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sta8 stores the lower part of the accumulator in the memory
func (cpu *CPU) sta8(addr uint32) {

	cpu.memory.SetByte(cpu.getARegister(), addr)
}

// sta stores the accumulator in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sta(laddr, haddr uint32) {
	if cpu.mFlag {
		cpu.sta8(laddr)
	} else {
		cpu.sta16(laddr, haddr)
	}
}

func (cpu *CPU) op81() {
	laddr, haddr := cpu.admPDirectXP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op83() {
	laddr, haddr := cpu.admStackSP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op85() {
	laddr, haddr := cpu.admDirectP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op87() {
	laddr, haddr := cpu.admBDirectP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8D() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op8F() {
	laddr, haddr := cpu.admLongP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) op91() {
	laddr, haddr := cpu.admPDirectYP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op92() {
	laddr, haddr := cpu.admPDirectP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op93() {
	laddr, haddr := cpu.admPStackSYP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op95() {
	laddr, haddr := cpu.admDirectXP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op97() {
	laddr, haddr := cpu.admBDirectYP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op99() {
	laddr, haddr := cpu.admAbsoluteYP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9D() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9F() {
	laddr, haddr := cpu.admLongXP()
	cpu.sta(laddr, haddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

// stx16 stores the x register in the memory
func (cpu *CPU) stx16(laddr, haddr uint32) {

	dataLo, dataHi := utils.SplitUint16(cpu.getXRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stx8 stores the lower part of the x register in the memory
func (cpu *CPU) stx8(addr uint32) {

	cpu.memory.SetByte(cpu.getXLRegister(), addr)
}

// stx stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stx(laddr, haddr uint32) {
	if cpu.xFlag {
		cpu.stx8(laddr)
	} else {
		cpu.stx16(laddr, haddr)
	}
}

func (cpu *CPU) op86() {
	laddr, haddr := cpu.admDirectP()
	cpu.stx(laddr, haddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8E() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.stx(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) op96() {
	laddr, haddr := cpu.admDirectYP()
	cpu.stx(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

// sty16 stores the x register in the memory
func (cpu *CPU) sty16(laddr, haddr uint32) {

	dataLo, dataHi := utils.SplitUint16(cpu.getYRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sty8 stores the lower part of the x register in the memory
func (cpu *CPU) sty8(addr uint32) {

	cpu.memory.SetByte(cpu.getYLRegister(), addr)
}

// sty stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sty(laddr, haddr uint32) {
	if cpu.xFlag {
		cpu.sty8(laddr)
	} else {
		cpu.sty16(laddr, haddr)
	}
}

func (cpu *CPU) op84() {
	laddr, haddr := cpu.admDirectP()
	cpu.sty(laddr, haddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.sty(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) op94() {
	laddr, haddr := cpu.admDirectXP()
	cpu.sty(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

// stz16 stores 0 in the memory
func (cpu *CPU) stz16(laddr, haddr uint32) {

	dataLo, dataHi := utils.SplitUint16(0x0000)

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stz8 stores 0 in the memory
func (cpu *CPU) stz8(addr uint32) {

	cpu.memory.SetByte(0x00, addr)
}

// stz stores 0 in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stz(laddr, haddr uint32) {
	if cpu.xFlag {
		cpu.stz8(laddr)
	} else {
		cpu.stz16(laddr, haddr)
	}
}

func (cpu *CPU) op64() {
	laddr, haddr := cpu.admDirectP()
	cpu.stz(laddr, haddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op74() {
	laddr, haddr := cpu.admDirectXP()
	cpu.stz(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op9C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.stz(laddr, haddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9E() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.stz(laddr, haddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}
