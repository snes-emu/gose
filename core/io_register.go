package core

import (
	"github.com/snes-emu/gose/io"
	"github.com/snes-emu/gose/utils"
)

func (cpu *CPU) initIORegisters() {
	cpu.registerIORegisters()
}

func (cpu *CPU) registerIORegisters() {
	cpu.ioRegisters[0x016] = io.NewRegister(cpu.joya, cpu.joywr)
	cpu.ioRegisters[0x017] = io.NewRegister(cpu.joyb, nil)
	cpu.ioRegisters[0x200] = io.NewRegister(cpu.joyb, nil)
	cpu.ioRegisters[0x201] = io.NewRegister(nil, cpu.wrio)
	cpu.ioRegisters[0x202] = io.NewRegister(nil, cpu.wrmpya)
	cpu.ioRegisters[0x203] = io.NewRegister(nil, cpu.wrmpyb)
	cpu.ioRegisters[0x204] = io.NewRegister(nil, cpu.wrdivl)
	cpu.ioRegisters[0x205] = io.NewRegister(nil, cpu.wrdivh)
	cpu.ioRegisters[0x206] = io.NewRegister(nil, cpu.wrdivb)
	cpu.ioRegisters[0x207] = io.NewRegister(nil, cpu.htimel)
	cpu.ioRegisters[0x208] = io.NewRegister(nil, cpu.h0xtime)
	cpu.ioRegisters[0x209] = io.NewRegister(nil, cpu.vtimel)
	cpu.ioRegisters[0x20a] = io.NewRegister(nil, cpu.v0xtime)
	cpu.ioRegisters[0x20d] = io.NewRegister(nil, cpu.memsel)
	cpu.ioRegisters[0x210] = io.NewRegister(cpu.rdnmi, nil)
	cpu.ioRegisters[0x211] = io.NewRegister(cpu.timeup, nil)
	cpu.ioRegisters[0x212] = io.NewRegister(cpu.hvbjoy, nil)
	cpu.ioRegisters[0x213] = io.NewRegister(cpu.rdio, nil)
	cpu.ioRegisters[0x214] = io.NewRegister(cpu.rddivl, nil)
	cpu.ioRegisters[0x215] = io.NewRegister(cpu.rddivh, nil)
	cpu.ioRegisters[0x216] = io.NewRegister(cpu.rdmpyl, nil)
	cpu.ioRegisters[0x217] = io.NewRegister(cpu.rdmpyh, nil)
	cpu.ioRegisters[0x218] = io.NewRegister(cpu.joy1l, nil)
	cpu.ioRegisters[0x21A] = io.NewRegister(cpu.joy2l, nil)
	cpu.ioRegisters[0x21B] = io.NewRegister(cpu.joy2h, nil)
	cpu.ioRegisters[0x21C] = io.NewRegister(cpu.joy3l, nil)
	cpu.ioRegisters[0x21D] = io.NewRegister(cpu.joy3h, nil)
	cpu.ioRegisters[0x21E] = io.NewRegister(cpu.joy4l, nil)
	cpu.ioRegisters[0x21F] = io.NewRegister(cpu.joy4h, nil)
}

// 0x4016/Write - JOYWR - Joypad Output (W)
func (cpu *CPU) joywr(data uint8) {
	// TODO
}

// 0x4016/Read  - JOYA  - Joypad Input Register A (R)
func (cpu *CPU) joya() uint8 {
	// TODO
	return 0
}

// 0x4017/Read  - JOYB  - Joypad Input Register B (R)
func (cpu *CPU) joyb() uint8 {
	// TODO
	return 0
}

// 0x4200 - NMITIMEN- Interrupt Enable and Joypad Request (W)
func (cpu *CPU) nmitimen(data uint8) {
	// TODO
}

// 0x4201 - WRIO    - Joypad Programmable I/O Port (Open-Collector Output) (W)
func (cpu *CPU) wrio(data uint8) {
	// TODO
}

// 0x4202 - WRMPYA  - Set unsigned 8bit Multiplicand (W)
func (cpu *CPU) wrmpya(data uint8) {
	cpu.ioMemory[0x202] = data
}

// 0x4203 - WRMPYB  - Set unsigned 8bit Multiplier and Start Multiplication (W)
func (cpu *CPU) wrmpyb(data uint8) {
	mult := uint(cpu.ioMemory[0x202]) * uint(data)
	ll, hh := utils.SplitUint16(uint16(mult))
	cpu.ioMemory[0x216] = ll
	cpu.ioMemory[0x217] = hh

	// also mutates rddivl and rddivh
	cpu.ioMemory[0x214] = data
	cpu.ioMemory[0x215] = 0x00
}

// 0x4204 - WRDIVL  - Set unsigned 16bit Dividend (lower 8bit) (W)
func (cpu *CPU) wrdivl(data uint8) {
	cpu.ioMemory[0x204] = data
}

// 0x4205 - WRDIVH - Set unsigned 16bit Dividend (upper 8bit) (W)
func (cpu *CPU) wrdivh(data uint8) {
	cpu.ioMemory[0x205] = data
}

// 0x4206 - WRDIVB  - Set unsigned 8bit Divisor and Start Division (W)
func (cpu *CPU) wrdivb(data uint8) {
	divisor := uint16(data)
	dividend := utils.JoinUint16(cpu.ioMemory[0x204], cpu.ioMemory[0x205])

	quotient := uint16(0xffff)
	remainder := dividend

	if data != 0 {
		quotient = uint16(dividend / divisor)
		remainder = dividend % divisor
	}

	llq, hhq := utils.SplitUint16(quotient)
	cpu.ioMemory[0x214] = llq
	cpu.ioMemory[0x215] = hhq

	llr, hhr := utils.SplitUint16(remainder)
	cpu.ioMemory[0x216] = llr
	cpu.ioMemory[0x217] = hhr
}

// 0x4207 - HTIMEL  - H-Count Timer Setting (lower 8bits) (W)
func (cpu *CPU) htimel(data uint8) {
	// TODO
}

// 0x4208 - H0xTIME  - H-Count Timer Setting (upper 1bit) (W)
func (cpu *CPU) h0xtime(data uint8) {
	// TODO
}

// 0x4209 - VTIMEL  - V-Count Timer Setting (lower 8bits) (W)
func (cpu *CPU) vtimel(data uint8) {
	// TODO
}

// 0x420A - V0xTIME  - V-Count Timer Setting (upper 1bit) (W)
func (cpu *CPU) v0xtime(data uint8) {
	// TODO
}

// 0x420D - MEMSEL  - Memory-2 Waitstate Control (W)
func (cpu *CPU) memsel(data uint8) {
	// TODO
}

// 0x4210 - RDNMI   - V-Blank NMI Flag and CPU Version Number (Read/Ack) (R)
func (cpu *CPU) rdnmi() uint8 {
	// TODO
	return 0
}

// 0x4211 - TIMEUP  - H/V-Timer IRQ Flag (Read/Ack)  (R)
func (cpu *CPU) timeup() uint8 {
	// TODO
	return 0
}

// 0x4212 - HVBJOY  - H/V-Blank flag and Joypad Busy flag (R) (R)
func (cpu *CPU) hvbjoy() uint8 {
	// TODO
	return 0
}

// 0x4213 - RDIO    - Joypad Programmable I/O Port (Input)  (R)
func (cpu *CPU) rdio() uint8 {
	// TODO
	return 0
}

// 0x4214 - RDDIVL  - Unsigned Division Result (Quotient) (lower 8bit)  (R)
func (cpu *CPU) rddivl() uint8 {
	// TODO
	return 0
}

// 0x4215 - RDDIVH  - Unsigned Division Result (Quotient) (upper 8bit) (R)
func (cpu *CPU) rddivh() uint8 {
	// TODO
	return 0
}

// 0x4216 - RDMPYL  - Unsigned Division Remainder / Multiply Product (lower 8bit) (R)
func (cpu *CPU) rdmpyl() uint8 {
	// TODO
	return 0
}

// 0x4217 - RDMPYH  - Unsigned Division Remainder / Multiply Product (upper 8bit) (R)
func (cpu *CPU) rdmpyh() uint8 {
	// TODO
	return 0
}

// 0x4218 - JOY1L   - Joypad 1 (gameport 1, pin 4) (lower 8bit) (R)
func (cpu *CPU) joy1l() uint8 {
	// TODO
	return 0
}

// 0x4219 - JOY1H   - Joypad 1 (gameport 1, pin 4) (upper 8bit) (R)

func (cpu *CPU) joy1h() uint8 {
	// TODO
	return 0
}

// 0x421A - JOY2L   - Joypad 2 (gameport 2, pin 4) (lower 8bit) (R)
func (cpu *CPU) joy2l() uint8 {
	// TODO
	return 0
}

// 0x421B - JOY2H   - Joypad 2 (gameport 2, pin 4) (upper 8bit) (R)
func (cpu *CPU) joy2h() uint8 {
	// TODO
	return 0
}

// 0x421C - JOY3L   - Joypad 3 (gameport 1, pin 5) (lower 8bit) (R)
func (cpu *CPU) joy3l() uint8 {
	// TODO
	return 0
}

// 0x421D - JOY3H   - Joypad 3 (gameport 1, pin 5) (upper 8bit) (R)
func (cpu *CPU) joy3h() uint8 {
	// TODO
	return 0
}

// 0x421E - JOY4L   - Joypad 4 (gameport 2, pin 5) (lower 8bit) (R)
func (cpu *CPU) joy4l() uint8 {
	// TODO
	return 0
}

// 0x421F - JOY4H   - Joypad 4 (gameport 2, pin 5) (upper 8bit) (R)
func (cpu *CPU) joy4h() uint8 {
	// TODO
	return 0
}
