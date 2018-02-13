package cpu

import "github.com/snes-emu/gose/utils"

// ldx16 load data into the x register
func (cpu *CPU) ldx16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data != 0

	cpu.setXRegister(data)
}

// ldx8 load data into the lower bits of the x register
func (cpu *CPU) ldx8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data != 0

	cpu.setXLRegister(data)
}

// ldx load data into the x register taking care of 16bit/8bit cases
func (cpu *CPU) ldx(dataHi, dataLo uint8) {
	if cpu.xFlag {
		cpu.ldx8(dataLo)
	} else {
		cpu.ldx16(utils.ReadUint16(dataHi, dataLo))
	}
}
