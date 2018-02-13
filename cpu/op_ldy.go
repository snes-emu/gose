package cpu

import "github.com/snes-emu/gose/utils"

// ldy16 load data into the y register
func (cpu *CPU) ldy16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data != 0

	cpu.setYRegister(data)
}

// ldy8 load data into the lower bits of the y register
func (cpu *CPU) ldy8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data != 0

	cpu.setYLRegister(data)
}

// ldy load data into the y register taking care of 16bit/8bit cases
func (cpu *CPU) ldy(dataHi, dataLo uint8) {
	if cpu.xFlag {
		cpu.ldy8(dataLo)
	} else {
		cpu.ldy16(utils.ReadUint16(dataHi, dataLo))
	}
}
