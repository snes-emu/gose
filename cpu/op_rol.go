package cpu

import "github.com/snes-emu/gose/utils"

// rol16acc performs a rotate left on the 16 bit accumulator
func (cpu *CPU) rol16acc() {
	result := cpu.getCRegister() << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// rol8 performs a rotate left on the lower 8 bit accumulator
func (cpu *CPU) rol8acc() {
	result := cpu.getARegister() << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = cpu.getARegister()&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// rol16data performs a rotate left on the 16 bit data
func (cpu *CPU) rol16data(haddress, laddress uint32) {
	dataHi, dataLo := cpu.memory.GetByte(haddress), cpu.memory.GetByte(laddress)

	data := utils.ReadUint16(dataHi, dataLo)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = data&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultHi, resultLo := utils.WriteUint16(data)

	cpu.memory.SetByte(resultHi, haddress)
	cpu.memory.SetByte(resultLo, laddress)
}

// rol8data performs a rotate left on the 8 bit data
func (cpu *CPU) rol8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = data&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// rol performs a rotate left taking care of 16/8bits cases
func (cpu *CPU) rol(haddress, laddress uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.rol8acc()
		} else {
			cpu.rol8data(laddress)
		}
	} else {
		if isAcc {
			cpu.rol16acc()
		} else {
			cpu.rol16data(haddress, laddress)
		}
	}
}
