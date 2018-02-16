package cpu

import "github.com/snes-emu/gose/utils"

func (cpu *CPU) op90() {
	offset := cpu.admRelative8()
	t := !cpu.cFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) opB0() {
	offset := cpu.admRelative8()
	t := cpu.cFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) opF0() {
	offset := cpu.admRelative8()
	t := cpu.zFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) op30() {
	offset := cpu.admRelative8()
	t := cpu.nFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) opD0() {
	offset := cpu.admRelative8()
	t := !cpu.zFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) op10() {
	offset := cpu.admRelative8()
	t := !cpu.nFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) op80() {
	offset := cpu.admRelative8()
	cpu.cycles += 3 + utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	cpu.PC += offset + 2
}

func (cpu *CPU) op50() {
	offset := cpu.admRelative8()
	t := !cpu.vFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) op70() {
	offset := cpu.admRelative8()
	t := cpu.vFlag
	cpu.cycles += 2 + utils.BoolToUint16[t] + utils.BoolToUint16[t]*utils.BoolToUint16[cpu.eFlag]*utils.BoolToUint16[cpu.pFlag]
	if t {
		cpu.PC += offset + 2
	} else {
		cpu.PC += 2
	}

}

func (cpu *CPU) op82() {
	offset := cpu.admRelative16()
	cpu.cycles += 4
	cpu.PC += offset + 2

}
