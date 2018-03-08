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
func (cpu *CPU) lda(dataHi, dataLo uint8) {
	if cpu.mFlag {
		cpu.lda8(dataLo)
	} else {
		cpu.lda16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA1() {
	dataHi, dataLo := cpu.admPDirectX()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA3() {
	dataHi, dataLo := cpu.admStackS()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opA5() {
	dataHi, dataLo := cpu.admDirect()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA7() {
	dataHi, dataLo := cpu.admBDirect()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opA9() {
	dataHi, dataLo := cpu.admImmediateM()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.mFlag]
}

func (cpu *CPU) opAD() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) opAF() {
	dataHi, dataLo := cpu.admLong()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) opB1() {
	dataHi, dataLo := cpu.admPDirectY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 2
}

func (cpu *CPU) opB2() {
	dataHi, dataLo := cpu.admPDirect()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB3() {
	dataHi, dataLo := cpu.admPStackSY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) opB5() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB7() {
	dataHi, dataLo := cpu.admBDirectY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opB9() {
	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opBD() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.lda(dataHi, dataLo)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

func (cpu *CPU) opBF() {
	dataHi, dataLo := cpu.admLongX()
	cpu.lda(dataHi, dataLo)
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
func (cpu *CPU) ldx(dataHi, dataLo uint8) {
	if cpu.xFlag {
		cpu.ldx8(dataLo)
	} else {
		cpu.ldx16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA2() {
	dataHi, dataLo := cpu.admImmediateX()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opA6() {
	dataHi, dataLo := cpu.admDirect()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opAE() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) opB6() {
	dataHi, dataLo := cpu.admDirectY()
	cpu.ldx(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opBE() {
	dataHi, dataLo := cpu.admAbsoluteY()
	cpu.ldx(dataHi, dataLo)
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
func (cpu *CPU) ldy(dataHi, dataLo uint8) {
	if cpu.xFlag {
		cpu.ldy8(dataLo)
	} else {
		cpu.ldy16(utils.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA0() {
	dataHi, dataLo := cpu.admImmediateX()
	cpu.ldy(dataHi, dataLo)
	cpu.cycles += 3 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3 - utils.BoolToUint16[cpu.xFlag]
}

func (cpu *CPU) opA4() {
	dataHi, dataLo := cpu.admDirect()
	cpu.ldy(dataHi, dataLo)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opAC() {
	dataHi, dataLo := cpu.admAbsolute()
	cpu.ldy(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) opB4() {
	dataHi, dataLo := cpu.admDirectX()
	cpu.ldy(dataHi, dataLo)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) opBC() {
	dataHi, dataLo := cpu.admAbsoluteX()
	cpu.ldy(dataHi, dataLo)
	cpu.cycles += 6 - 2*utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.xFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += 3
}

// sta16 stores the accumulator in the memory
func (cpu *CPU) sta16(haddr, laddr uint32) {

	dataLo, dataHi := utils.SplitUint16(cpu.getCRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sta8 stores the lower part of the accumulator in the memory
func (cpu *CPU) sta8(addr uint32) {

	cpu.memory.SetByte(cpu.getARegister(), addr)
}

// sta stores the accumulator in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sta(haddr, laddr uint32) {
	if cpu.mFlag {
		cpu.sta8(laddr)
	} else {
		cpu.sta16(haddr, laddr)
	}
}

func (cpu *CPU) op81() {
	haddr, laddr := cpu.admPDirectXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op83() {
	haddr, laddr := cpu.admStackSP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op85() {
	haddr, laddr := cpu.admDirectP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op87() {
	haddr, laddr := cpu.admBDirectP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8D() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op8F() {
	haddr, laddr := cpu.admLongP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

func (cpu *CPU) op91() {
	haddr, laddr := cpu.admPDirectYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op92() {
	haddr, laddr := cpu.admPDirectP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op93() {
	haddr, laddr := cpu.admPStackSYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 8 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 2
}

func (cpu *CPU) op95() {
	haddr, laddr := cpu.admDirectXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op97() {
	haddr, laddr := cpu.admBDirectYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op99() {
	haddr, laddr := cpu.admAbsoluteYP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9D() {
	haddr, laddr := cpu.admAbsoluteXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9F() {
	haddr, laddr := cpu.admLongXP()
	cpu.sta(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 4
}

// stx16 stores the x register in the memory
func (cpu *CPU) stx16(haddr, laddr uint32) {

	dataLo, dataHi := utils.SplitUint16(cpu.getXRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stx8 stores the lower part of the x register in the memory
func (cpu *CPU) stx8(addr uint32) {

	cpu.memory.SetByte(cpu.getXLRegister(), addr)
}

// stx stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stx(haddr, laddr uint32) {
	if cpu.xFlag {
		cpu.stx8(laddr)
	} else {
		cpu.stx16(haddr, laddr)
	}
}

func (cpu *CPU) op86() {
	haddr, laddr := cpu.admDirectP()
	cpu.stx(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8E() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.stx(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) op96() {
	haddr, laddr := cpu.admDirectYP()
	cpu.stx(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

// sty16 stores the x register in the memory
func (cpu *CPU) sty16(haddr, laddr uint32) {

	dataLo, dataHi := utils.SplitUint16(cpu.getYRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sty8 stores the lower part of the x register in the memory
func (cpu *CPU) sty8(addr uint32) {

	cpu.memory.SetByte(cpu.getYLRegister(), addr)
}

// sty stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sty(haddr, laddr uint32) {
	if cpu.xFlag {
		cpu.sty8(laddr)
	} else {
		cpu.sty16(haddr, laddr)
	}
}

func (cpu *CPU) op84() {
	haddr, laddr := cpu.admDirectP()
	cpu.sty(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op8C() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.sty(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag]
	cpu.PC += 3
}

func (cpu *CPU) op94() {
	haddr, laddr := cpu.admDirectXP()
	cpu.sty(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.xFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

// stz16 stores 0 in the memory
func (cpu *CPU) stz16(haddr, laddr uint32) {

	dataLo, dataHi := utils.SplitUint16(0x0000)

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stz8 stores 0 in the memory
func (cpu *CPU) stz8(addr uint32) {

	cpu.memory.SetByte(0x00, addr)
}

// stz stores 0 in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stz(haddr, laddr uint32) {
	if cpu.xFlag {
		cpu.stz8(laddr)
	} else {
		cpu.stz16(haddr, laddr)
	}
}

func (cpu *CPU) op64() {
	haddr, laddr := cpu.admDirectP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 4 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op74() {
	haddr, laddr := cpu.admDirectXP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
	cpu.PC += 2
}

func (cpu *CPU) op9C() {
	haddr, laddr := cpu.admAbsoluteP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 5 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}

func (cpu *CPU) op9E() {
	haddr, laddr := cpu.admAbsoluteXP()
	cpu.stz(haddr, laddr)
	cpu.cycles += 6 - utils.BoolToUint16[cpu.mFlag]
	cpu.PC += 3
}
