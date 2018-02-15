package cpu

// SOURCE,DESTINATION addressing mode
func (cpu CPU) admSourceDestination() (uint8, uint16, uint8, uint16) {
	SBank := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+1)
	DBank := cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()+2)
	var SAddress, DAddress uint16
	if cpu.xFlag {
		SAddress = uint16(cpu.getXLRegister())
		DAddress = uint16(cpu.getYLRegister())
	} else {
		SAddress = cpu.getXRegister()
		DAddress = cpu.getYRegister()
	}
	return SBank, SAddress, DBank, DAddress
}
