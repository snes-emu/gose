package cpu

// adc performs an add with carry operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) adc func() {
	if cpu.getDFlag() == 0 {
		// Decial mode off -> binary arithmetic used
	} else {
		// Decimal mode on -> BCD arithmetic used
	}
}

opcodes[0x61] = (cpu *CPU) func() {
	
	w := 1
	if cpu.getDLRegister() == 0 {
		w := 0
	}

	cpu.cycles += 7 - cpu.getMFlag() + w
}