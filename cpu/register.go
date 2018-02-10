package cpu

// getARegister returns the lower 8 bits of the accumulator
func (cpu CPU) getARegister() uint8 {}

// getBRegister returns the upper 8 bits of the accumulator
func (cpu CPU) getBRegister() uint8 {}

// getCRegister returns the 16 bits accumulator
func (cpu CPU) getCRegister() uint16 {}

// getDBRRegister returns the data bank register
func (cpu CPU) getDBRRegister() uint16 {}

// getDRegister returns the D register
func (cpu CPU) getDRegister() uint16 {}

// getDLRegister returns the lower 8 bits of the direct register
func (cpu CPU) getDLRegister() uint8 {}

// getDHRegister returns the upper 8 bits of the direct register
func (cpu CPU) getDHRegister() uint8 {}

// getKRegister returns the program bank register
func (cpu CPU) getKRegister() uint8 {}

// getPCRegister returns the program counter
func (cpu CPU) getPCRegister() uint16 {}

// getPCLRegister returns the lower 8 bits of the program counter
func (cpu CPU) getPCLRegister() uint8 {}

// getPCHRegister returns the lower 8 bits of the program counter
func (cpu CPU) getPCHRegister() uint8 {}

// getPRegister returns the processor status register
func (cpu CPU) getPRegister() uint8 {}

// getSRegister returns the stack pointer
func (cpu CPU) getSRegister() uint16 {}

// getSLRegister returns the lower 8 bits of the stack pointer
func (cpu CPU) getSLRegister() uint8 {}

// getSHRegister returns the upper 8 bits of the stack pointer
func (cpu CPU) getSHRegister() uint8 {}

// getXRegister returns the X index register
func (cpu CPU) getXRegister() uint16 {}

// getXLRegister returns the lower 8 bits of the X index register
func (cpu CPU) getXLRegister() uint8 {}

// getXHRegister returns the upper 8 bits of the X index register
func (cpu CPU) getXHRegister() uint8 {}

// getYRegister returns the Y index register
func (cpu CPU) getYRegister() uint16 {}

// getYLRegister returns the lower 8 bits of the Y index register
func (cpu CPU) getYLRegister() uint8 {}

// getYHRegister returns the upper 8 bits of the Y index register
func (cpu CPU) getYHRegister() uint8 {}
