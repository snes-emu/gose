package cpu

// adc performs an add with carry operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) adc16 func(data, carry uint16) {
	if cpu.getDFlag() == 0 {
		// Decial mode off -> binary arithmetic used
		result := cpu.C + data + carry
	} else {
		// Decimal mode on -> BCD arithmetic used
		result = (cpu.C & 0x000f) + (data & 0x000f) + carry + (cpu.C & 0x00f0) + (data & 0x00f0) + (cpu.C & 0x0f00) + (data & 0x0f00) + (cpu.C & 0xf000) + (data & 0xf000) 
	}
	cpu.d = result >> 14
	if data > 65535 - cpu.C {
		cpu.v = 1
		cpu.c = 1
	} else {
		cpu.v = 0
		cpu.c = 0
	}
	if result == 0 {
		cpu.z = 1
	} else {
		cpu.z = 0
	}
}

opcodes[0x61] = (cpu *CPU) func() {
	
	w := 1
	if cpu.getDLRegister() == 0 {
		w := 0
	}

	cpu.cycles += 7 - cpu.getMFlag() + w
}