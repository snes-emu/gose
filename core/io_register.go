package core

import (
	"github.com/snes-emu/gose/bit"
	"github.com/snes-emu/gose/io"
)

type ioMemory struct {
	bytes           [0x380]uint8 // Raw bytes
	hirqPos         uint16       // Horizontal IRQ position
	virqPos         uint16       // Vertical IRQ position
	irqFlag         bool         // IRQ Flag used in TIMEUP
	vBlankNMIEnable bool         // VBlank NMI Enable  (0=Disable, 1=Enable) (Initially disabled on reset)
	hvIRQ           uint8        // H/V IRQ (0=Disable, 1=At H=H + V=Any, 2=At V=V + H=0, 3=At H=H + V=V)
	joypadEnable    bool         // Joypad Enable    (0=Disable, 1=Enable Automatic Reading of Joypad) TODO
	vBlankNMIFlag   bool         // (0=None, 1=Interrupt Request) (set on Begin of Vblank)
}

func (cpu *CPU) initIORegisters(rf *io.RegisterFactory) {
	cpu.ioMemory = &ioMemory{bytes: [0x380]uint8{}}
	for i := 0; i < 0x380; i++ {
		cpu.ioRegisters[i] = rf.NewRegister(nil, nil)
	}
	cpu.registerIORegisters(rf)
	cpu.initDma(rf)
}

func (cpu *CPU) registerIORegisters(rf *io.RegisterFactory) {

	cpu.ioRegisters[0x016] = rf.NewRegister(cpu.joya, cpu.joywr, "JOY")
	cpu.ioRegisters[0x017] = rf.NewRegister(cpu.joyb, nil, "JOYB")
	cpu.ioRegisters[0x200] = rf.NewRegister(nil, cpu.nmitimen, "NMITIMEN")
	cpu.ioRegisters[0x201] = rf.NewRegister(nil, cpu.wrio, "WRIO")
	cpu.ioRegisters[0x202] = rf.NewRegister(nil, cpu.wrmpya, "WRMPYA")
	cpu.ioRegisters[0x203] = rf.NewRegister(nil, cpu.wrmpyb, "WRMPYB")
	cpu.ioRegisters[0x204] = rf.NewRegister(nil, cpu.wrdivl, "WRDIVL")
	cpu.ioRegisters[0x205] = rf.NewRegister(nil, cpu.wrdivh, "WRDIVH")
	cpu.ioRegisters[0x206] = rf.NewRegister(nil, cpu.wrdivb, "WRDIVB")
	cpu.ioRegisters[0x207] = rf.NewRegister(nil, cpu.htimel, "HTIMEL")
	cpu.ioRegisters[0x208] = rf.NewRegister(nil, cpu.h0xtime, "H0XTIME")
	cpu.ioRegisters[0x209] = rf.NewRegister(nil, cpu.vtimel, "VTIMEL")
	cpu.ioRegisters[0x20a] = rf.NewRegister(nil, cpu.v0xtime, "V0XTIME")
	cpu.ioRegisters[0x20d] = rf.NewRegister(nil, cpu.memsel, "MEMSEL")
	cpu.ioRegisters[0x210] = rf.NewRegister(cpu.rdnmi, nil, "RDNMI")
	cpu.ioRegisters[0x211] = rf.NewRegister(cpu.timeup, nil, "TIMEUP")
	cpu.ioRegisters[0x212] = rf.NewRegister(cpu.hvbjoy, nil, "HVBJOY")
	cpu.ioRegisters[0x213] = rf.NewRegister(cpu.rdio, nil, "RDIO")
	cpu.ioRegisters[0x214] = rf.NewRegister(cpu.rddivl, nil, "RDDIVL")
	cpu.ioRegisters[0x215] = rf.NewRegister(cpu.rddivh, nil, "RDDIVH")
	cpu.ioRegisters[0x216] = rf.NewRegister(cpu.rdmpyl, nil, "RDMPYL")
	cpu.ioRegisters[0x217] = rf.NewRegister(cpu.rdmpyh, nil, "RDMPYH")
	cpu.ioRegisters[0x218] = rf.NewRegister(cpu.joy1l, nil, "JOY1L")
	cpu.ioRegisters[0x219] = rf.NewRegister(cpu.joy1h, nil, "JOY1H")
	cpu.ioRegisters[0x21A] = rf.NewRegister(cpu.joy2l, nil, "JOY2L")
	cpu.ioRegisters[0x21B] = rf.NewRegister(cpu.joy2h, nil, "JOY2H")
	cpu.ioRegisters[0x21C] = rf.NewRegister(cpu.joy3l, nil, "JOY3L")
	cpu.ioRegisters[0x21D] = rf.NewRegister(cpu.joy3h, nil, "JOY3H")
	cpu.ioRegisters[0x21E] = rf.NewRegister(cpu.joy4l, nil, "JOY4L")
	cpu.ioRegisters[0x21F] = rf.NewRegister(cpu.joy4h, nil, "JOY4H")
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
	cpu.ioMemory.vBlankNMIEnable = data&0x80 != 0
	cpu.ioMemory.hvIRQ = (data & 0x30) >> 4
	cpu.ioMemory.joypadEnable = data&0x01 != 0
}

// 0x4201 - WRIO    - Joypad Programmable I/O Port (Open-Collector Output) (W)
func (cpu *CPU) wrio(data uint8) {
	// TODO
}

// 0x4202 - WRMPYA  - Set unsigned 8bit Multiplicand (W)
func (cpu *CPU) wrmpya(data uint8) {
	cpu.ioMemory.bytes[0x202] = data
}

// 0x4203 - WRMPYB  - Set unsigned 8bit Multiplier and Start Multiplication (W)
func (cpu *CPU) wrmpyb(data uint8) {
	mult := uint(cpu.ioMemory.bytes[0x202]) * uint(data)
	ll, hh := bit.SplitUint16(uint16(mult))
	cpu.ioMemory.bytes[0x216] = ll
	cpu.ioMemory.bytes[0x217] = hh

	// also mutates rddivl and rddivh
	cpu.ioMemory.bytes[0x214] = data
	cpu.ioMemory.bytes[0x215] = 0x00
}

// 0x4204 - WRDIVL  - Set unsigned 16bit Dividend (lower 8bit) (W)
func (cpu *CPU) wrdivl(data uint8) {
	cpu.ioMemory.bytes[0x204] = data
}

// 0x4205 - WRDIVH - Set unsigned 16bit Dividend (upper 8bit) (W)
func (cpu *CPU) wrdivh(data uint8) {
	cpu.ioMemory.bytes[0x205] = data
}

// 0x4206 - WRDIVB  - Set unsigned 8bit Divisor and Start Division (W)
func (cpu *CPU) wrdivb(data uint8) {
	divisor := uint16(data)
	dividend := bit.JoinUint16(cpu.ioMemory.bytes[0x204], cpu.ioMemory.bytes[0x205])

	quotient := uint16(0xffff)
	remainder := dividend

	if data != 0 {
		quotient = uint16(dividend / divisor)
		remainder = dividend % divisor
	}

	llq, hhq := bit.SplitUint16(quotient)
	cpu.ioMemory.bytes[0x214] = llq
	cpu.ioMemory.bytes[0x215] = hhq

	llr, hhr := bit.SplitUint16(remainder)
	cpu.ioMemory.bytes[0x216] = llr
	cpu.ioMemory.bytes[0x217] = hhr
}

// 0x4207 - HTIMEL  - H-Count Timer Setting (lower 8bits) (W)
func (cpu *CPU) htimel(data uint8) {
	cpu.ioMemory.hirqPos = (cpu.ioMemory.hirqPos & 0xff00) | uint16(data)
}

// 0x4208 - H0xTIME  - H-Count Timer Setting (upper 1bit) (W)
func (cpu *CPU) h0xtime(data uint8) {
	cpu.ioMemory.hirqPos = (cpu.ioMemory.hirqPos & 0xff00) | ((uint16(data) << 8) & 0x100)
}

// 0x4209 - VTIMEL  - V-Count Timer Setting (lower 8bits) (W)
func (cpu *CPU) vtimel(data uint8) {
	cpu.ioMemory.virqPos = (cpu.ioMemory.virqPos & 0xff00) | uint16(data)
}

// 0x420A - V0xTIME  - V-Count Timer Setting (upper 1bit) (W)
func (cpu *CPU) v0xtime(data uint8) {
	cpu.ioMemory.virqPos = (cpu.ioMemory.virqPos & 0xff00) | ((uint16(data) << 8) & 0x100)
}

// 0x420D - MEMSEL  - Memory-2 Waitstate Control (W)
func (cpu *CPU) memsel(data uint8) {
	// TODO
}

// 0x4210 - RDNMI   - V-Blank NMI Flag and CPU Version Number (Read/Ack) (R)
func (cpu *CPU) rdnmi() uint8 {
	// TODO: maybe the version is not correct there
	version := uint8(2)
	res := (bit.BoolToUint8(cpu.ioMemory.vBlankNMIFlag)<<7 | version)
	cpu.ioMemory.vBlankNMIFlag = false
	return res
}

// 0x4211 - TIMEUP  - H/V-Timer IRQ Flag (Read/Ack)  (R)
func (cpu *CPU) timeup() uint8 {
	var res uint8

	if cpu.ioMemory.irqFlag {
		res = 0x80
		cpu.ioMemory.irqFlag = false
	}

	return res
}

// 0x4212 - HVBJOY  - H/V-Blank flag and Joypad Busy flag (R) (R)
func (cpu *CPU) hvbjoy() uint8 {
	// TODO: bit 0 of res should be used there !!! see documentation for further information
	var res uint8

	// HBlank
	hc := cpu.ppu.HCounter()

	if hc < HBLANKEND || hc > HBLANKSTART {
		res |= 0x40
	}

	// VBlank
	if cpu.ppu.VCounter() > cpu.ppu.VDisplay() {
		res |= 0x80
	}

	return res
}

// 0x4213 - RDIO    - Joypad Programmable I/O Port (Input)  (R)
func (cpu *CPU) rdio() uint8 {
	// TODO
	return 0
}

// 0x4214 - RDDIVL  - Unsigned Division Result (Quotient) (lower 8bit)  (R)
func (cpu *CPU) rddivl() uint8 {
	return cpu.ioMemory.bytes[0x214]
}

// 0x4215 - RDDIVH  - Unsigned Division Result (Quotient) (upper 8bit) (R)
func (cpu *CPU) rddivh() uint8 {
	return cpu.ioMemory.bytes[0x215]
}

// 0x4216 - RDMPYL  - Unsigned Division Remainder / Multiply Product (lower 8bit) (R)
func (cpu *CPU) rdmpyl() uint8 {
	return cpu.ioMemory.bytes[0x216]
}

// 0x4217 - RDMPYH  - Unsigned Division Remainder / Multiply Product (upper 8bit) (R)
func (cpu *CPU) rdmpyh() uint8 {
	return cpu.ioMemory.bytes[0x217]
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
