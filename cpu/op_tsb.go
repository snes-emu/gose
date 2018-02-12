package cpu

import (
	"github.com/snes-emu/gose/utils"
)

// adc16 performs an add with carry 16bit operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) tsb16(data uint16) uint16 {
	cpu.zFlag = cpu.getCRegister()&data == 0
	return data | cpu.getCRegister()

}

// adc8 performs an add with carry 8bit operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) tsb8(data uint8) uint8 {
	cpu.zFlag = cpu.getARegister()&data == 0
	return data | cpu.getARegister()
}

// sbc performs a substract with carry operation handling the 8/16 bit cases
func (cpu *CPU) tsb(addressHi, addressLo uint32) {
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.tsb8(cpu.memory.GetByte(addressLo)), addressLo)
	} else {
		result := cpu.tsb16(utils.ReadUint16(cpu.memory.GetByte(addressHi), cpu.memory.GetByte(addressLo)))
		resultHi, resultLo := utils.WriteUint16(result)
		cpu.memory.SetByte(resultLo, addressLo)
		cpu.memory.SetByte(resultHi, addressHi)
	}
}
