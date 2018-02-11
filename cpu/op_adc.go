package cpu

import "github.com/snes-emu/gose/utils"

// adc performs an add with carry operation the formula is: accumulator = accumulator + data + carry

func (cpu *CPU) adc16(data, carry uint16) {
	var result uint16
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		result = (cpu.C & 0x000f) + (data & 0x000f) + carry + (cpu.C & 0x00f0) + (data & 0x00f0) + (cpu.C & 0x0f00) + (data & 0x0f00) + (cpu.C & 0xf000) + (data & 0xf000)
	} else {
		// Decial mode off -> binary arithmetic used
		result = cpu.C + data + carry
	}
	if data > 65535-cpu.C {
		cpu.vFlag = true
		cpu.cFlag = true
	} else {
		cpu.vFlag = false
		cpu.cFlag = false
	}
	if result == 0 {
		cpu.zFlag = true
	} else {
		cpu.zFlag = false
	}
}

func (cpu *CPU) op61() {
	// TODO

	cpu.cycles += 7 - utils.BoolToUint16[cpu.mFlag] + utils.BoolToUint16[cpu.getDLRegister() == 0]
}
