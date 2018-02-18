package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) branch(cond bool, offset uint16) {
	cpu.cycles += 2 + utils.BoolToUint16[cond] + utils.BoolToUint16[cond]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += offset*utils.BoolToUint16[cond] + 2
}

func (cpu *CPU) bcc(offset uint16) {
	cpu.branch(!cpu.cFlag, offset)
}

func (cpu *CPU) op90() {
	cpu.bcc(cpu.admRelative8())
}

func (cpu *CPU) bcs(offset uint16) {
	cpu.branch(cpu.cFlag, offset)
}

func (cpu *CPU) opB0() {
	cpu.bcs(cpu.admRelative8())
}

func (cpu *CPU) beq(offset uint16) {
	cpu.branch(cpu.zFlag, offset)
}

func (cpu *CPU) opF0() {
	cpu.beq(cpu.admRelative8())
}

func (cpu *CPU) bmi(offset uint16) {
	cpu.branch(cpu.nFlag, offset)
}

func (cpu *CPU) op30() {
	cpu.bmi(cpu.admRelative8())
}

func (cpu *CPU) bne(offset uint16) {
	cpu.branch(!cpu.zFlag, offset)
}

func (cpu *CPU) opD0() {
	cpu.bne(cpu.admRelative8())
}

func (cpu *CPU) bpl(offset uint16) {
	cpu.branch(!cpu.nFlag, offset)
}

func (cpu *CPU) op10() {
	cpu.bpl(cpu.admRelative8())
}

func (cpu *CPU) bra(offset uint16) {
	cpu.cycles += 3 + utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += offset + 2
}

func (cpu *CPU) op80() {
	cpu.bra(cpu.admRelative8())
}

func (cpu *CPU) bvc(offset uint16) {
	cpu.branch(!cpu.vFlag, offset)
}

func (cpu *CPU) op50() {
	cpu.bvc(cpu.admRelative8())
}

func (cpu *CPU) bvs(offset uint16) {
	cpu.branch(cpu.vFlag, offset)
}

func (cpu *CPU) op70() {
	cpu.bvs(cpu.admRelative8())
}

func (cpu *CPU) brl(offset uint16) {
	cpu.cycles += 4
	cpu.PC += offset + 2
}

func (cpu *CPU) op82() {
	cpu.brl(cpu.admRelative16())
}
