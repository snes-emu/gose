package core

import (
	"os"

	"github.com/snes-emu/gose/log"

	"github.com/snes-emu/gose/bit"
)

const (
	abortNativeVector    = 0xFFE8
	nmiNativeVector      = 0xFFEA
	resetNativeVector    = 0xFFEC
	irqNativeVector      = 0xFFEE
	abortEmulationVector = 0xFFF8
	nmiEmulationVector   = 0xFFFA
	resetEmulationVector = 0xFFFC
	irqEmulationVector   = 0xFFFE
)

func (cpu *CPU) abort() {
	addressLo, addressHi := bit.SplitUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, abortEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, abortEmulationVector+1)
		cpu.PC = bit.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, abortNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, abortNativeVector+1)
		cpu.PC = bit.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) nmi() {
	addressLo, addressHi := bit.SplitUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, nmiEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, nmiEmulationVector+1)
		cpu.PC = bit.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, nmiNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, nmiNativeVector+1)
		cpu.PC = bit.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) reset() {
	cpu.setEFlag(true)
	cpu.D = 0x0000
	cpu.DBR = 0x00
	cpu.K = 0x00
	cpu.S = 0x01FF
	addressLo := cpu.memory.GetByteBank(0x00, resetEmulationVector)
	addressHi := cpu.memory.GetByteBank(0x00, resetEmulationVector+1)
	cpu.PC = bit.JoinUint16(addressLo, addressHi)
	cpu.dFlag = false
	cpu.iFlag = true

}

func (cpu *CPU) irq() {
	addressLo, addressHi := bit.SplitUint16(cpu.getPCRegister())
	if cpu.eFlag {
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, irqEmulationVector)
		addressHi := cpu.memory.GetByteBank(0x00, irqEmulationVector+1)
		cpu.PC = bit.JoinUint16(addressLo, addressHi)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(addressHi)
		cpu.pushStack(addressLo)
		cpu.php()
		cpu.K = 0x00
		addressLo := cpu.memory.GetByteBank(0x00, irqNativeVector)
		addressHi := cpu.memory.GetByteBank(0x00, irqNativeVector+1)
		cpu.PC = bit.JoinUint16(addressLo, addressHi)
	}
	cpu.dFlag = false
	cpu.iFlag = true

}

// bit16 performs a bitwise and for 16bits variables
func (cpu *CPU) bit16(data uint16, isImmediate bool) uint16 {
	result := cpu.getCRegister() & data

	// Last bit value
	if !isImmediate {
		cpu.nFlag = data&0x8000 != 0
		cpu.vFlag = data&0x4000 != 0
	}

	cpu.zFlag = result == 0

	return result
}

// bit8 performs a bitwise and for 8bits variables
func (cpu *CPU) bit8(data uint8, isImmediate bool) uint8 {
	result := cpu.getARegister() & data

	// Last bit value
	if !isImmediate {
		cpu.nFlag = data&0x80 != 0
		cpu.vFlag = data&0x40 != 0
	}

	cpu.zFlag = result == 0

	return result
}

// bit performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) bit(dataLo, dataHi uint8, isImmediate bool) {
	if cpu.mFlag {
		cpu.bit8(dataLo, isImmediate)
	} else {
		cpu.bit16(bit.JoinUint16(dataLo, dataHi), isImmediate)
	}
}

func (cpu *CPU) op24() {
	dataLo, dataHi := cpu.admDirect()
	cpu.bit(dataLo, dataHi, false)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op2C() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.bit(dataLo, dataHi, false)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op34() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.bit(dataLo, dataHi, false)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op3C() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.bit(dataLo, dataHi, false)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op89() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.bit(dataLo, dataHi, true)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) branch(cond bool, offset uint16) {
	cpu.PC += offset*bit.BoolToUint16(cond) + 2
	cpu.step(2 + bit.BoolToUint16(cond) + bit.BoolToUint16(cond)*bit.BoolToUint16(cpu.eFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) bcc(offset uint16) {
	cpu.branch(!cpu.cFlag, offset)
}

func (cpu *CPU) op90() {
	cpu.bcc(cpu.admRelative8())
}

func (cpu *CPU) bcs(offset uint16) {
	cpu.branch(cpu.cFlag, offset)
}

func (cpu *CPU) opB0() {
	cpu.bcs(cpu.admRelative8())
}

func (cpu *CPU) beq(offset uint16) {
	cpu.branch(cpu.zFlag, offset)
}

func (cpu *CPU) opF0() {
	cpu.beq(cpu.admRelative8())
}

func (cpu *CPU) bmi(offset uint16) {
	cpu.branch(cpu.nFlag, offset)
}

func (cpu *CPU) op30() {
	cpu.bmi(cpu.admRelative8())
}

func (cpu *CPU) bne(offset uint16) {
	cpu.branch(!cpu.zFlag, offset)
}

func (cpu *CPU) opD0() {
	cpu.bne(cpu.admRelative8())
}

func (cpu *CPU) bpl(offset uint16) {
	cpu.branch(!cpu.nFlag, offset)
}

func (cpu *CPU) op10() {
	cpu.bpl(cpu.admRelative8())
}

func (cpu *CPU) bra(offset uint16) {
	cpu.PC += offset + 2
	cpu.step(3 + bit.BoolToUint16(cpu.eFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op80() {
	cpu.bra(cpu.admRelative8())
}

func (cpu *CPU) bvc(offset uint16) {
	cpu.branch(!cpu.vFlag, offset)
}

func (cpu *CPU) op50() {
	cpu.bvc(cpu.admRelative8())
}

func (cpu *CPU) bvs(offset uint16) {
	cpu.branch(cpu.vFlag, offset)
}

func (cpu *CPU) op70() {
	cpu.bvs(cpu.admRelative8())
}

func (cpu *CPU) brl(offset uint16) {
	cpu.PC += offset + 3
	cpu.step(4)
}

func (cpu *CPU) op82() {
	cpu.brl(cpu.admRelative16())
}

// adc16 performs an add with carry 16bit operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) adc16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in adc needs to be implemented")
		//result = (cpu.getCRegister() & 0x000f) + (data & 0x000f) + bit.BoolToUint16(cpu.cFlag) + (cpu.getCRegister() & 0x00f0) + (data & 0x00f0) + (cpu.C & 0x0f00) + (data & 0x0f00) + (cpu.getCRegister() & 0xf000) + (data & 0xf000)

	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getCRegister() + data + bit.BoolToUint16(cpu.cFlag)
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Signed artihmetic overflow
		cpu.vFlag = (data^result)&^(data^cpu.getCRegister())&0x8000 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = (result < cpu.getCRegister()) || (result == cpu.getCRegister() && cpu.cFlag)

	}

	return result
}

// adc8 performs an add with carry 8bit operation the formula is: accumulator = accumulator + data + carry
func (cpu *CPU) adc8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in adc needs to be implemented")
	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getARegister() + data + bit.BoolToUint8(cpu.cFlag)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Signed artihmetic overflow
		cpu.vFlag = (data^result)&^(data^cpu.getARegister())&0x80 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = (result < cpu.getARegister()) || (result == cpu.getARegister() && cpu.cFlag)

	}

	return result
}

// sbc performs a substract with carry operation handling the 8/16 bit cases
func (cpu *CPU) adc(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.adc8(dataLo))
	} else {
		cpu.setCRegister(cpu.adc16(bit.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op61() {

	dataLo, dataHi := cpu.admPDirectX()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op63() {

	dataLo, dataHi := cpu.admStackS()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op65() {

	dataLo, dataHi := cpu.admDirect()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op67() {

	dataLo, dataHi := cpu.admBDirect()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op69() {

	dataLo, dataHi := cpu.admImmediateM()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op6D() {

	dataLo, dataHi := cpu.admAbsolute()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op6F() {

	dataLo, dataHi := cpu.admLong()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op71() {

	dataLo, dataHi := cpu.admPDirectY()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op72() {
	dataLo, dataHi := cpu.admPDirect()

	cpu.adc(dataLo, dataHi)

	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op73() {

	dataLo, dataHi := cpu.admPStackSY()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op75() {
	dataLo, dataHi := cpu.admDirectX()

	cpu.adc(dataLo, dataHi)

	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op77() {

	dataLo, dataHi := cpu.admBDirectY()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op79() {
	dataLo, dataHi := cpu.admAbsoluteY()

	cpu.adc(dataLo, dataHi)

	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.mFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op7D() {

	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op7F() {
	dataLo, dataHi := cpu.admLongX()
	cpu.adc(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// sbc16 performs a substract with carry 16bit operation the formula is: accumulator = accumulator - data - 1 + carry
func (cpu *CPU) sbc16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in sbc needs to be implemented")

	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getCRegister() - data - 1 + bit.BoolToUint16(cpu.cFlag)
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Signed artihmetic overflow
		cpu.vFlag = ((data+1)^cpu.getCRegister())&^((data+1)^result)&0x8000 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = cpu.getCRegister() >= data

	}

	return result
}

// sbc8 performs a substract with carry 8bit operation the formula is: accumulator = accumulator - data - 1 + carry
func (cpu *CPU) sbc8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		// Decimal mode on -> BCD arithmetic used
		panic("TODO, d flag in sbc needs to be implemented")

	} else {
		// Decimal mode off -> binary arithmetic used
		result = cpu.getARegister() - data - 1 + bit.BoolToUint8(cpu.cFlag)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Signed artihmetic overflow
		cpu.vFlag = ((data+1)^cpu.getARegister())&^((data+1)^result)&0x80 != 0
		cpu.zFlag = result == 0
		// Unsigned carry
		cpu.cFlag = cpu.getARegister() >= data

	}

	return result
}

// sbc performs a substract with carry operation handling the 8/16 bit cases
func (cpu *CPU) sbc(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.sbc8(dataLo))
	} else {
		cpu.setCRegister(cpu.sbc16(bit.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) opE1() {

	dataLo, dataHi := cpu.admPDirectX()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opE3() {

	dataLo, dataHi := cpu.admStackS()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opE5() {

	dataLo, dataHi := cpu.admDirect()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opE7() {

	dataLo, dataHi := cpu.admBDirect()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opE9() {

	dataLo, dataHi := cpu.admImmediateM()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opED() {

	dataLo, dataHi := cpu.admAbsolute()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opEF() {

	dataLo, dataHi := cpu.admLong()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opF1() {

	dataLo, dataHi := cpu.admPDirectY()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) opF2() {
	dataLo, dataHi := cpu.admPDirect()

	cpu.sbc(dataLo, dataHi)

	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opF3() {

	dataLo, dataHi := cpu.admPStackSY()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opF5() {
	dataLo, dataHi := cpu.admDirectX()

	cpu.sbc(dataLo, dataHi)

	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opF7() {

	dataLo, dataHi := cpu.admBDirectY()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opF9() {
	dataLo, dataHi := cpu.admAbsoluteY()

	cpu.sbc(dataLo, dataHi)

	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.mFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) opFD() {

	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.sbc(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) opFF() {
	dataLo, dataHi := cpu.admLongX()

	cpu.sbc(dataLo, dataHi)

	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// cmp16 does a 16bit comparison the accumulator to the data
func (cpu *CPU) cmp16(data uint16) {
	result := cpu.getCRegister() - data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getCRegister() >= data
}

// cmp8 does a 8bit comparison the accumulator to the data
func (cpu *CPU) cmp8(data uint8) {
	result := cpu.getARegister() - data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getARegister() >= data
}

// cmp compare the accumulator to the data handling the 16bit/8bit distinction
func (cpu *CPU) cmp(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.cmp8(dataLo)
	} else {
		cpu.cmp16(bit.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opC1() {

	dataLo, dataHi := cpu.admPDirectX()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opC3() {

	dataLo, dataHi := cpu.admStackS()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opC5() {

	dataLo, dataHi := cpu.admDirect()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opC7() {

	dataLo, dataHi := cpu.admBDirect()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opC9() {

	dataLo, dataHi := cpu.admImmediateM()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opCD() {

	dataLo, dataHi := cpu.admAbsolute()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opCF() {

	dataLo, dataHi := cpu.admLong()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opD1() {

	dataLo, dataHi := cpu.admPDirectY()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) + bit.BoolToUint16(cpu.xFlag)*(bit.BoolToUint16(cpu.pFlag)-1))
}

func (cpu *CPU) opD2() {

	dataLo, dataHi := cpu.admPDirect()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opD3() {

	dataLo, dataHi := cpu.admPStackSY()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opD5() {

	dataLo, dataHi := cpu.admDirectX()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opD7() {

	dataLo, dataHi := cpu.admBDirectY()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opD9() {

	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.xFlag)*(bit.BoolToUint16(cpu.pFlag)-1))
}

func (cpu *CPU) opDD() {

	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.xFlag)*(bit.BoolToUint16(cpu.pFlag)-1))
}

func (cpu *CPU) opDF() {

	dataLo, dataHi := cpu.admLongX()
	cpu.cmp(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// cpx16 does a 16bit comparison of the X register with the data
func (cpu *CPU) cpx16(data uint16) {
	result := cpu.getXRegister() - data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getXRegister() >= data
}

// cpx8 does a 8bit comparison of the X register with the data
func (cpu *CPU) cpx8(data uint8) {
	result := cpu.getXLRegister() - data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getXLRegister() >= data
}

// cpx compare the X register to the data handling the 16bit/8bit distinction
func (cpu *CPU) cpx(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.cpx8(dataLo)
	} else {
		cpu.cpx16(bit.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opE0() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.cpx(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.xFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) opE4() {
	dataLo, dataHi := cpu.admDirect()
	cpu.cpx(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opEC() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.cpx(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

// cpy16 does a 16bit comparison of the Y register with the data
func (cpu *CPU) cpy16(data uint16) {
	result := cpu.getYRegister() - data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getYRegister() >= data
}

// cpy8 does a 8bit comparison of the Y register with the data
func (cpu *CPU) cpy8(data uint8) {
	result := cpu.getYLRegister() - data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0
	// Unsigned carry
	cpu.cFlag = cpu.getYLRegister() >= data
}

// cpy compare the Y register to the data handling the 16bit/8bit distinction
func (cpu *CPU) cpy(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.cpy8(dataLo)
	} else {
		cpu.cpy16(bit.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opC0() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.cpy(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.xFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) opC4() {
	dataLo, dataHi := cpu.admDirect()
	cpu.cpy(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opCC() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.cpy(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) dec16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data - 1
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

func (cpu *CPU) dec8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data - 1
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

//op3A performs a decrement operation on the accumulator
func (cpu *CPU) op3A() {
	dataLo, dataHi := cpu.admAccumulator()
	if cpu.mFlag {
		cpu.setARegister(cpu.dec8(dataLo))
	} else {
		cpu.setCRegister(cpu.dec16(bit.JoinUint16(dataLo, dataHi)))
	}
	cpu.PC++
	cpu.step(2)
}

//opC6 performs a decrement operation on memory through direct addressing mode
func (cpu *CPU) opC6() {
	laddr, haddr := cpu.admDirectP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.dec16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

//opCE performs a decrement operation on memory through the absolute addressing mode
func (cpu *CPU) opCE() {
	laddr, haddr := cpu.admAbsoluteP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.dec16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

//opD6 performs a decrement operation on memory through direct,X addressing mode
func (cpu *CPU) opD6() {
	laddr, haddr := cpu.admDirectXP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.dec16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 2
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

//opDE performs a decrement operation on memory through absolute,X addressing mode
func (cpu *CPU) opDE() {
	laddr, haddr := cpu.admAbsoluteXP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.dec8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.dec16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 3
	cpu.step(9 - 2*bit.BoolToUint16(cpu.mFlag))
}

//opCA performs a decrement operation on the X register
func (cpu *CPU) opCA() {
	if cpu.xFlag {
		result := cpu.getXLRegister() - 1
		cpu.setXLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.X--
		// Last bit value
		cpu.nFlag = cpu.X&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.X == 0
	}
	cpu.PC++
	cpu.step(2)
}

//op88 performs a decrement operation on the Y register, immediate mode
func (cpu *CPU) op88() {
	if cpu.xFlag {
		result := cpu.getYLRegister() - 1
		cpu.setYLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.Y--
		// Last bit value
		cpu.nFlag = cpu.Y&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.Y == 0
	}
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) inc16(data uint16) uint16 {
	var result uint16
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data + 1
		// Last bit value
		cpu.nFlag = result&0x8000 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

func (cpu *CPU) inc8(data uint8) uint8 {
	var result uint8
	if cpu.dFlag {
		panic("TODO, d flag in dec needs to be implemented")
	} else {
		result = data + 1
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
		return result
	}
}

//op1A performs a increment operation on the accumulator
func (cpu *CPU) op1A() {
	dataLo, dataHi := cpu.admAccumulator()
	if cpu.mFlag {
		cpu.setARegister(cpu.inc8(dataLo))
	} else {
		cpu.setCRegister(cpu.inc16(bit.JoinUint16(dataLo, dataHi)))
	}
	cpu.PC++
	cpu.step(2)
}

//opE6 performs a increment operation on memory through direct addressing mode
func (cpu *CPU) opE6() {
	laddr, haddr := cpu.admDirectP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.inc16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

//opEE performs a increment operation through the absolute access mode
func (cpu *CPU) opEE() {
	laddr, haddr := cpu.admAbsoluteP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.inc16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

//opF6 performs a increment operation on memory through direct,X addressing mode
func (cpu *CPU) opF6() {
	laddr, haddr := cpu.admDirectXP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.inc16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 2
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

//opF6 performs a increment operation on memory through absolute,X addressing mode
func (cpu *CPU) opFE() {
	laddr, haddr := cpu.admAbsoluteXP()
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.inc8(dataLo), laddr)
	} else {
		resultLo, resultHi := bit.SplitUint16(cpu.inc16(bit.JoinUint16(dataLo, dataHi)))
		cpu.memory.SetByte(resultHi, haddr)
		cpu.memory.SetByte(resultLo, laddr)
	}
	cpu.PC += 3
	cpu.step(9 - 2*bit.BoolToUint16(cpu.mFlag))
}

//opE8 performs a increment operation on the X register, immediate mode
func (cpu *CPU) opE8() {
	if cpu.xFlag {
		result := cpu.getXLRegister() + 1
		cpu.setXLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.X++
		// Last bit value
		cpu.nFlag = cpu.X&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.X == 0
	}
	cpu.PC++
	cpu.step(2)
}

//opC8 performs a increment operation on the Y register, immediate mode
func (cpu *CPU) opC8() {
	if cpu.xFlag {
		result := cpu.getYLRegister() + 1
		cpu.setYLRegister(result)
		// Last bit value
		cpu.nFlag = result&0x80 != 0
		// Zero result flag
		cpu.zFlag = result == 0
	} else {
		cpu.Y++
		// Last bit value
		cpu.nFlag = cpu.Y&0x8000 != 0
		// Zero result flag
		cpu.zFlag = cpu.Y == 0
		cpu.step(2)
	}
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) xba() {
	temp := cpu.getBRegister()
	cpu.setBRegister(cpu.getARegister())
	cpu.setARegister(temp)
	cpu.nFlag = temp&0x80 != 0
	cpu.zFlag = temp == 0
	cpu.PC++
	cpu.step(3)
}

func (cpu *CPU) opEB() {
	cpu.xba()
}

func (cpu *CPU) xce() {
	temp := cpu.eFlag
	cpu.setEFlag(cpu.cFlag)
	cpu.cFlag = temp
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) opFB() {
	cpu.xce()
}

func (cpu *CPU) clc() {
	cpu.cFlag = false
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op18() {
	cpu.clc()
}

func (cpu *CPU) cld() {
	cpu.dFlag = false
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) opD8() {
	cpu.cld()
}

func (cpu *CPU) cli() {
	cpu.iFlag = false
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op58() {
	cpu.cli()
}

func (cpu *CPU) clv() {
	cpu.vFlag = false
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) opB8() {
	cpu.clv()
}

func (cpu *CPU) sec() {
	cpu.cFlag = true
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op38() {
	cpu.sec()
}

func (cpu *CPU) sed() {
	cpu.dFlag = true
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) opF8() {
	cpu.sed()
}

func (cpu *CPU) sei() {
	cpu.iFlag = true
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op78() {
	cpu.sei()
}

func (cpu *CPU) rep() {
	dataLo, _ := cpu.admImmediate8()
	cpu.cFlag = cpu.cFlag && dataLo&0x01 == 0
	cpu.zFlag = cpu.zFlag && dataLo&0x02 == 0
	cpu.iFlag = cpu.iFlag && dataLo&0x04 == 0
	cpu.dFlag = cpu.dFlag && dataLo&0x08 == 0
	if cpu.eFlag {
		cpu.bFlag = cpu.bFlag && dataLo&0x10 == 0
	} else {
		cpu.setXFlag(cpu.xFlag && dataLo&0x10 == 0)
	}
	cpu.mFlag = cpu.mFlag && dataLo&0x20 == 0
	cpu.vFlag = cpu.vFlag && dataLo&0x40 == 0
	cpu.nFlag = cpu.nFlag && dataLo&0x80 == 0
	cpu.PC += 2
	cpu.step(3)

}

func (cpu *CPU) opC2() {
	cpu.rep()
}

func (cpu *CPU) sep() {
	dataLo, _ := cpu.admImmediate8()
	cpu.cFlag = cpu.cFlag || dataLo&0x01 != 0
	cpu.zFlag = cpu.zFlag || dataLo&0x02 != 0
	cpu.iFlag = cpu.iFlag || dataLo&0x04 != 0
	cpu.dFlag = cpu.dFlag || dataLo&0x08 != 0
	if cpu.eFlag {
		cpu.bFlag = cpu.bFlag || dataLo&0x10 != 0
	} else {
		cpu.setXFlag(cpu.xFlag || dataLo&0x10 != 0)
	}
	cpu.mFlag = cpu.mFlag || dataLo&0x20 != 0
	cpu.vFlag = cpu.vFlag || dataLo&0x40 != 0
	cpu.nFlag = cpu.nFlag || dataLo&0x80 != 0
	cpu.PC += 2
	cpu.step(3)
}

func (cpu *CPU) opE2() {
	cpu.sep()
}

// stp STP stops the clock input of the 65C816,
func (cpu *CPU) stp() {
	log.Info("CPU has been shutdown")
	os.Exit(0)
}

func (cpu *CPU) opDB() {
	cpu.stp()
}

// wai stops the clock input of the 65C816,
func (cpu *CPU) wai() {
	cpu.waiting = true
}

func (cpu *CPU) opCB() {
	cpu.wai()
}

// jmp jumps to the address specified by the addressing mode
func (cpu *CPU) jmp(addr uint16) {
	cpu.PC = addr
}

// jmpLong jumps to the address specified by the long addressing
func (cpu *CPU) jmpLong(laddr uint16, haddr uint8) {
	cpu.K = haddr
	cpu.PC = laddr
}

func (cpu *CPU) op4C() {
	addr := cpu.admAbsoluteJ()
	cpu.jmp(addr)
	cpu.step(3)
}

func (cpu *CPU) op5C() {
	laddr, haddr := cpu.admLongJ()
	cpu.jmpLong(laddr, haddr)
	cpu.step(4)
}

func (cpu *CPU) op6C() {
	addr := cpu.admPAbsoluteJ()
	cpu.jmp(addr)
	cpu.step(5)
}

func (cpu *CPU) op7C() {
	addr := cpu.admPAbsoluteXJ()
	cpu.jmp(addr)
	cpu.step(6)
}

func (cpu *CPU) opDC() {
	laddr, haddr := cpu.admBAbsoluteJ()
	cpu.jmpLong(haddr, laddr)
	cpu.step(6)
}

// jsl jumps to a subroutine long
func (cpu *CPU) jsl(laddr uint16, haddr uint8) {
	laddr2, haddr2 := bit.SplitUint16(cpu.getPCRegister() + 3)
	cpu.pushStackNew24(laddr2, haddr2, cpu.getKRegister())

	cpu.jmpLong(laddr, haddr)
}

func (cpu *CPU) op22() {
	laddr, haddr := cpu.admLongJ()
	cpu.jsl(laddr, haddr)
	cpu.step(3)
}

// jsr jumps to a subroutine
func (cpu *CPU) jsr(addr uint16) {
	laddr, haddr := bit.SplitUint16(cpu.getPCRegister() + 2)

	cpu.pushStack(haddr)
	cpu.pushStack(laddr)

	cpu.jmp(addr)
}

// jsr jumps to a subroutine for new addressing mode
func (cpu *CPU) jsrNew(addr uint16) {
	laddr, haddr := bit.SplitUint16(cpu.getPCRegister() + 2)

	cpu.pushStackNew16(laddr, haddr)

	cpu.jmp(addr)
}

func (cpu *CPU) op20() {
	addr := cpu.admAbsoluteJ()
	cpu.jsr(addr)
	cpu.step(6)
}

func (cpu *CPU) opFC() {
	addr := cpu.admPAbsoluteXJ()
	cpu.jsrNew(addr)
	cpu.step(8)
}

// lda16 load data into the accumulator
func (cpu *CPU) lda16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data == 0
	cpu.setCRegister(data)
}

// lda8 load data into the lower bits of the accumulator
func (cpu *CPU) lda8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data == 0

	cpu.setARegister(data)
}

// lda load data into the accumulator
func (cpu *CPU) lda(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.lda8(dataLo)
	} else {
		cpu.lda16(bit.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA1() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opA3() {
	dataLo, dataHi := cpu.admStackS()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opA5() {
	dataLo, dataHi := cpu.admDirect()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opA7() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opA9() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opAD() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opAF() {
	dataLo, dataHi := cpu.admLong()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opB1() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) opB2() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opB3() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) opB5() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opB7() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opB9() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) opBD() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) opBF() {
	dataLo, dataHi := cpu.admLongX()
	cpu.lda(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// ldx16 load data into the x register
func (cpu *CPU) ldx16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data == 0

	cpu.setXRegister(data)
}

// ldx8 load data into the lower bits of the x register
func (cpu *CPU) ldx8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data == 0

	cpu.setXLRegister(data)
}

// ldx load data into the x register taking care of 16bit/8bit cases
func (cpu *CPU) ldx(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.ldx8(dataLo)
	} else {
		cpu.ldx16(bit.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA2() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.ldx(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.xFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) opA6() {
	dataLo, dataHi := cpu.admDirect()
	cpu.ldx(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opAE() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.ldx(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) opB6() {
	dataLo, dataHi := cpu.admDirectY()
	cpu.ldx(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opBE() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.ldx(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - 2*bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

// ldy16 load data into the y register
func (cpu *CPU) ldy16(data uint16) {

	// Last bit value
	cpu.nFlag = data&0x8000 != 0
	cpu.zFlag = data == 0

	cpu.setYRegister(data)
}

// ldy8 load data into the lower bits of the y register
func (cpu *CPU) ldy8(data uint8) {

	// Last bit value
	cpu.nFlag = data&0x80 != 0
	cpu.zFlag = data == 0

	cpu.setYLRegister(data)
}

// ldy load data into the y register taking care of 16bit/8bit cases
func (cpu *CPU) ldy(dataLo, dataHi uint8) {
	if cpu.xFlag {
		cpu.ldy8(dataLo)
	} else {
		cpu.ldy16(bit.JoinUint16(dataLo, dataHi))
	}
}

func (cpu *CPU) opA0() {
	dataLo, dataHi := cpu.admImmediateX()
	cpu.ldy(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.xFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) opA4() {
	dataLo, dataHi := cpu.admDirect()
	cpu.ldy(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opAC() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.ldy(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) opB4() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.ldy(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) opBC() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.ldy(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - 2*bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

// sta16 stores the accumulator in the memory
func (cpu *CPU) sta16(laddr, haddr uint32) {

	dataLo, dataHi := bit.SplitUint16(cpu.getCRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sta8 stores the lower part of the accumulator in the memory
func (cpu *CPU) sta8(addr uint32) {

	cpu.memory.SetByte(cpu.getARegister(), addr)
}

// sta stores the accumulator in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sta(laddr, haddr uint32) {
	if cpu.mFlag {
		cpu.sta8(laddr)
	} else {
		cpu.sta16(laddr, haddr)
	}
}

func (cpu *CPU) op81() {
	laddr, haddr := cpu.admPDirectXP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op83() {
	laddr, haddr := cpu.admStackSP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op85() {
	laddr, haddr := cpu.admDirectP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op87() {
	laddr, haddr := cpu.admBDirectP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op8D() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.sta(laddr, haddr)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op8F() {
	laddr, haddr := cpu.admLongP()
	cpu.sta(laddr, haddr)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op91() {
	laddr, haddr := cpu.admPDirectYP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op92() {
	laddr, haddr := cpu.admPDirectP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op93() {
	laddr, haddr := cpu.admPStackSYP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op95() {
	laddr, haddr := cpu.admDirectXP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op97() {
	laddr, haddr := cpu.admBDirectYP()
	cpu.sta(laddr, haddr)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op99() {
	laddr, haddr := cpu.admAbsoluteYP()
	cpu.sta(laddr, haddr)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op9D() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.sta(laddr, haddr)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op9F() {
	laddr, haddr := cpu.admLongXP()
	cpu.sta(laddr, haddr)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// stx16 stores the x register in the memory
func (cpu *CPU) stx16(laddr, haddr uint32) {

	dataLo, dataHi := bit.SplitUint16(cpu.getXRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stx8 stores the lower part of the x register in the memory
func (cpu *CPU) stx8(addr uint32) {

	cpu.memory.SetByte(cpu.getXLRegister(), addr)
}

// stx stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stx(laddr, haddr uint32) {
	if cpu.xFlag {
		cpu.stx8(laddr)
	} else {
		cpu.stx16(laddr, haddr)
	}
}

func (cpu *CPU) op86() {
	laddr, haddr := cpu.admDirectP()
	cpu.stx(laddr, haddr)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op8E() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.stx(laddr, haddr)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) op96() {
	laddr, haddr := cpu.admDirectYP()
	cpu.stx(laddr, haddr)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

// sty16 stores the x register in the memory
func (cpu *CPU) sty16(laddr, haddr uint32) {

	dataLo, dataHi := bit.SplitUint16(cpu.getYRegister())

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// sty8 stores the lower part of the x register in the memory
func (cpu *CPU) sty8(addr uint32) {

	cpu.memory.SetByte(cpu.getYLRegister(), addr)
}

// sty stores the x register in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) sty(laddr, haddr uint32) {
	if cpu.xFlag {
		cpu.sty8(laddr)
	} else {
		cpu.sty16(laddr, haddr)
	}
}

func (cpu *CPU) op84() {
	laddr, haddr := cpu.admDirectP()
	cpu.sty(laddr, haddr)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op8C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.sty(laddr, haddr)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

func (cpu *CPU) op94() {
	laddr, haddr := cpu.admDirectXP()
	cpu.sty(laddr, haddr)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

// stz16 stores 0 in the memory
func (cpu *CPU) stz16(laddr, haddr uint32) {

	dataLo, dataHi := bit.SplitUint16(0x0000)

	cpu.memory.SetByte(dataHi, haddr)
	cpu.memory.SetByte(dataLo, laddr)
}

// stz8 stores 0 in the memory
func (cpu *CPU) stz8(addr uint32) {

	cpu.memory.SetByte(0x00, addr)
}

// stz stores 0 in the memory taking care of the 16bit/8bit cases
func (cpu *CPU) stz(laddr, haddr uint32) {
	if cpu.mFlag {
		cpu.stz8(laddr)
	} else {
		cpu.stz16(laddr, haddr)
	}
}

func (cpu *CPU) op64() {
	laddr, haddr := cpu.admDirectP()
	cpu.stz(laddr, haddr)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op74() {
	laddr, haddr := cpu.admDirectXP()
	cpu.stz(laddr, haddr)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op9C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.stz(laddr, haddr)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op9E() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.stz(laddr, haddr)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// and16 performs a bitwise and for 16bits variables
func (cpu *CPU) and16(data uint16) uint16 {
	result := cpu.getCRegister() & data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	return result
}

// and8 performs a bitwise and for 8bits variables
func (cpu *CPU) and8(data uint8) uint8 {
	result := cpu.getARegister() & data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	return result
}

// and performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) and(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.and8(dataLo))
	} else {
		cpu.setCRegister(cpu.and16(bit.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op21() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op23() {
	dataLo, dataHi := cpu.admStackS()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op25() {
	dataLo, dataHi := cpu.admDirect()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op27() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op29() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.and(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op2D() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.and(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op2F() {
	dataLo, dataHi := cpu.admLong()
	cpu.and(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op31() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op32() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op33() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op35() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op37() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.and(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op39() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.and(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op3D() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.and(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op3F() {
	dataLo, dataHi := cpu.admLongX()
	cpu.and(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// eor16 performs a bitwise exclusive or for 16bits variables
func (cpu *CPU) eor16(data uint16) uint16 {
	result := cpu.getCRegister() ^ data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	return result
}

// eor8 performs a bitwise and for 8bits variables
func (cpu *CPU) eor8(data uint8) uint8 {
	result := cpu.getARegister() ^ data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	return result
}

// eor performs a bitwise and taking care of 16/8bits cases
func (cpu *CPU) eor(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.eor8(dataLo))
	} else {
		cpu.setCRegister(cpu.eor16(bit.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op41() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op43() {
	dataLo, dataHi := cpu.admStackS()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op45() {
	dataLo, dataHi := cpu.admDirect()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op47() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op49() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op4D() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op4F() {
	dataLo, dataHi := cpu.admLong()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op51() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op52() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op53() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op55() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op57() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op59() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op5D() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op5F() {
	dataLo, dataHi := cpu.admLongX()
	cpu.eor(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

// ora16 performs a bitwise or for 16bits variables
func (cpu *CPU) ora16(data uint16) uint16 {
	result := cpu.getCRegister() | data

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	return result
}

// ora8 performs a bitwise or for 8bits variables
func (cpu *CPU) ora8(data uint8) uint8 {
	result := cpu.getARegister() | data

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	return result
}

// ora performs a bitwise or taking care of 16/8bits cases
func (cpu *CPU) ora(dataLo, dataHi uint8) {
	if cpu.mFlag {
		cpu.setARegister(cpu.ora8(dataLo))
	} else {
		cpu.setCRegister(cpu.ora16(bit.JoinUint16(dataLo, dataHi)))
	}
}

func (cpu *CPU) op01() {
	dataLo, dataHi := cpu.admPDirectX()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op03() {
	dataLo, dataHi := cpu.admStackS()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op05() {
	dataLo, dataHi := cpu.admDirect()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op07() {
	dataLo, dataHi := cpu.admBDirect()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op09() {
	dataLo, dataHi := cpu.admImmediateM()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 3 - bit.BoolToUint16(cpu.mFlag)
	cpu.step(3 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op0D() {
	dataLo, dataHi := cpu.admAbsolute()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op0F() {
	dataLo, dataHi := cpu.admLong()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op11() {
	dataLo, dataHi := cpu.admPDirectY()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op12() {
	dataLo, dataHi := cpu.admPDirect()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op13() {
	dataLo, dataHi := cpu.admPStackSY()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(8 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op15() {
	dataLo, dataHi := cpu.admDirectX()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op17() {
	dataLo, dataHi := cpu.admBDirectY()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(7 - bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op19() {
	dataLo, dataHi := cpu.admAbsoluteY()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op1D() {
	dataLo, dataHi := cpu.admAbsoluteX()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag) - bit.BoolToUint16(cpu.xFlag) + bit.BoolToUint16(cpu.xFlag)*bit.BoolToUint16(cpu.pFlag))
}

func (cpu *CPU) op1F() {
	dataLo, dataHi := cpu.admLongX()
	cpu.ora(dataLo, dataHi)
	cpu.PC += 4
	cpu.step(6 - bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) mvn(SBank uint8, SAddress uint16, DBank uint8, DAddress uint16) {
	cpu.memory.SetByteBank(cpu.memory.GetByteBank(SBank, SAddress), DBank, DAddress)
	cpu.DBR = DBank
	cpu.C--
	if cpu.xFlag {
		cpu.setXLRegister(cpu.getXLRegister() + 1)
		cpu.setYLRegister(cpu.getYLRegister() + 1)
	} else {
		cpu.X++
		cpu.Y++
	}
	cpu.step(7)
	if cpu.getCRegister() == 0xFFFF {
		cpu.PC += 3
	}
}

func (cpu *CPU) op54() {
	SBank, SAddress, DBank, DAddress := cpu.admSourceDestination()
	cpu.mvn(SBank, SAddress, DBank, DAddress)
}

func (cpu *CPU) mvp(SBank uint8, SAddress uint16, DBank uint8, DAddress uint16) {
	cpu.memory.SetByteBank(cpu.memory.GetByteBank(SBank, SAddress), DBank, DAddress)
	cpu.DBR = DBank
	cpu.C--
	if cpu.xFlag {
		cpu.setXLRegister(cpu.getXLRegister() - 1)
		cpu.setYLRegister(cpu.getYLRegister() - 1)
	} else {
		cpu.X--
		cpu.Y--
	}
	cpu.step(7)
	if cpu.getCRegister() == 0xFFFF {
		cpu.PC += 3
	}
}

func (cpu *CPU) op44() {
	SBank, SAddress, DBank, DAddress := cpu.admSourceDestination()
	cpu.mvp(SBank, SAddress, DBank, DAddress)
}

func (cpu *CPU) nop() {
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) opEA() {
	cpu.nop()
}

func (cpu *CPU) wdm() {
	cpu.PC += 2
	cpu.step(2)
}

func (cpu *CPU) op42() {
	cpu.wdm()
}

func (cpu *CPU) rti() {
	cpu.plp()
	addressLo := cpu.pullStack()
	addressHi := cpu.pullStack()
	cpu.PC = bit.JoinUint16(addressLo, addressHi)
	if !cpu.eFlag {
		cpu.K = cpu.pullStack()
	}
	cpu.step(7 - bit.BoolToUint16(cpu.eFlag))

}

func (cpu *CPU) op40() {
	cpu.rti()
}

func (cpu *CPU) rtl() {
	PCLo, PCHi, K := cpu.pullStackNew24()
	cpu.K = K
	cpu.PC = bit.JoinUint16(PCLo, PCHi) + 1
	cpu.step(6)
}

func (cpu *CPU) op6B() {
	cpu.rtl()
}

func (cpu *CPU) rts() {
	PCLo := cpu.pullStack()
	PCHi := cpu.pullStack()
	cpu.PC = bit.JoinUint16(PCLo, PCHi) + 1
	cpu.step(6)
}

func (cpu *CPU) op60() {
	cpu.rts()
}

// asl16acc performs a left shift on the 16 bit accumulator
func (cpu *CPU) asl16acc() {
	result := cpu.getCRegister() << 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// asl8 performs a left shift on the lower 8 bit accumulator
func (cpu *CPU) asl8acc() {
	result := cpu.getARegister() << 1

	// Get the highbit before shifting
	cpu.cFlag = cpu.getARegister()&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// asl16data performs a left shift on the 16 bit data
func (cpu *CPU) asl16data(laddr, haddr uint32) {
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)

	data := bit.JoinUint16(dataLo, dataHi)

	result := data << 1

	// Get the highbit before shifting
	cpu.cFlag = data&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultLo, resultHi := bit.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddr)
	cpu.memory.SetByte(resultLo, laddr)
}

// asl8data performs a left shift on the 8 bit data
func (cpu *CPU) asl8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data << 1

	// Get the highbit before shifting
	cpu.cFlag = data&0x80 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// asl performs a left shift taking care of 16/8bits cases
func (cpu *CPU) asl(laddr, haddr uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.asl8acc()
		} else {
			cpu.asl8data(laddr)
		}
	} else {
		if isAcc {
			cpu.asl16acc()
		} else {
			cpu.asl16data(laddr, haddr)
		}
	}
}

func (cpu *CPU) op06() {
	laddr, haddr := cpu.admDirectP()
	cpu.asl(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op0A() {
	cpu.asl(0, 0, true)
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op0E() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.asl(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op16() {
	laddr, haddr := cpu.admDirectXP()
	cpu.asl(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op1E() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.asl(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(9 - 2*bit.BoolToUint16(cpu.mFlag))
}

// lsr16acc performs a right shift on the 16 bit accumulator
func (cpu *CPU) lsr16acc() {
	result := cpu.getCRegister() >> 1

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// lsr8 performs a right shift on the lower 8 bit accumulator
func (cpu *CPU) lsr8acc() {
	result := cpu.getARegister() >> 1

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getARegister()&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// lsr16data performs a right shift on the 16 bit data
func (cpu *CPU) lsr16data(haddr, laddr uint32) {
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)

	data := bit.JoinUint16(dataLo, dataHi)

	result := data >> 1

	// Get the lowbit before shifting
	cpu.cFlag = data&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultLo, resultHi := bit.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddr)
	cpu.memory.SetByte(resultLo, laddr)
}

// lsr8data performs a right shift on the 8 bit data
func (cpu *CPU) lsr8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data >> 1

	// Get the lowbit before shifting
	cpu.cFlag = data&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// lsr performs a right shift taking care of 16/8bits cases
func (cpu *CPU) lsr(laddr, haddr uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.lsr8acc()
		} else {
			cpu.lsr8data(laddr)
		}
	} else {
		if isAcc {
			cpu.lsr16acc()
		} else {
			cpu.lsr16data(laddr, haddr)
		}
	}
}

func (cpu *CPU) op46() {
	laddr, haddr := cpu.admDirectP()
	cpu.lsr(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op4A() {
	cpu.lsr(0, 0, true)
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op4E() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.lsr(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op56() {
	laddr, haddr := cpu.admDirectXP()
	cpu.lsr(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op5E() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.lsr(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(9 - 2*bit.BoolToUint16(cpu.mFlag))
}

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
func (cpu *CPU) rol16data(laddr, haddr uint32) {
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)

	data := bit.JoinUint16(dataLo, dataHi)

	result := data << 1

	if cpu.cFlag {
		result = result | 0x01
	}

	// Get the highbit before shifting
	cpu.cFlag = data&0x8000 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultLo, resultHi := bit.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddr)
	cpu.memory.SetByte(resultLo, laddr)
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
func (cpu *CPU) rol(laddr, haddr uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.rol8acc()
		} else {
			cpu.rol8data(laddr)
		}
	} else {
		if isAcc {
			cpu.rol16acc()
		} else {
			cpu.rol16data(laddr, haddr)
		}
	}
}

func (cpu *CPU) op26() {
	laddr, haddr := cpu.admDirectP()
	cpu.rol(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op2A() {
	cpu.rol(0, 0, true)
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op2E() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.rol(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op36() {
	laddr, haddr := cpu.admDirectXP()
	cpu.rol(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op3E() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.rol(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(9 - 2*bit.BoolToUint16(cpu.mFlag))
}

// ror16acc performs a rotate right on the 16 bit accumulator
func (cpu *CPU) ror16acc() {
	result := cpu.getCRegister() >> 1

	if cpu.cFlag {
		result = result | 0x8000
	}

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getCRegister()&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// ror8 performs a rotate right on the lower 8 bit accumulator
func (cpu *CPU) ror8acc() {
	result := cpu.getARegister() >> 1

	if cpu.cFlag {
		result = result | 0x80
	}

	// Get the lowbit before shifting
	cpu.cFlag = cpu.getARegister()&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

// ror16data performs a rotate right on the 16 bit data
func (cpu *CPU) ror16data(laddr, haddr uint32) {
	dataLo, dataHi := cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)

	data := bit.JoinUint16(dataLo, dataHi)

	result := data >> 1

	if cpu.cFlag {
		result = result | 0x8000
	}

	// Get the lowbit before shifting
	cpu.cFlag = data&0x0001 != 0

	// Last bit value
	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	resultLo, resultHi := bit.SplitUint16(data)

	cpu.memory.SetByte(resultHi, haddr)
	cpu.memory.SetByte(resultLo, laddr)
}

// ror8data performs a rotate right on the 8 bit data
func (cpu *CPU) ror8data(addr uint32) {
	data := cpu.memory.GetByte(addr)

	result := data >> 1

	if cpu.cFlag {
		result = result | 0x80
	}

	// Get the lowbit before shifting
	cpu.cFlag = data&0x01 != 0

	// Last bit value
	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.memory.SetByte(result, addr)
}

// ror performs a rotate right taking care of 16/8bits cases
func (cpu *CPU) ror(laddr, haddr uint32, isAcc bool) {
	if cpu.mFlag {
		if isAcc {
			cpu.ror8acc()
		} else {
			cpu.ror8data(laddr)
		}
	} else {
		if isAcc {
			cpu.ror16acc()
		} else {
			cpu.ror16data(laddr, haddr)
		}
	}
}

func (cpu *CPU) op66() {
	laddr, haddr := cpu.admDirectP()
	cpu.ror(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op6A() {
	cpu.ror(0, 0, true)
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) op6E() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.ror(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) op76() {
	laddr, haddr := cpu.admDirectXP()
	cpu.ror(laddr, haddr, false)
	cpu.PC += 2
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op7E() {
	laddr, haddr := cpu.admAbsoluteXP()
	cpu.ror(laddr, haddr, false)
	cpu.PC += 3
	cpu.step(9 - 2*bit.BoolToUint16(cpu.mFlag))
}

//p16 pushes the next 16-bit value into the stack
func (cpu *CPU) p16(dataLo, dataHi uint8) {
	cpu.pushStack(dataHi)
	cpu.pushStack(dataLo)
}

//p8 pushes the next 8-bit value into the stack
func (cpu *CPU) p8(data uint8) {
	cpu.pushStack(data)
}

// PEA instruction
func (cpu *CPU) opF4() {
	dataLo, dataHi := cpu.admImmediate16()
	cpu.pushStackNew16(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(5)
}

// PEI instruction
func (cpu *CPU) opD4() {
	dataLo, dataHi := cpu.admDirectNew()
	cpu.pushStackNew16(dataLo, dataHi)
	cpu.PC += 2
	cpu.step(6 + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

// PER instuction
func (cpu *CPU) op62() {
	dataLo, dataHi := cpu.admImmediate16()
	cpu.pushStackNew16(dataLo, dataHi)
	cpu.PC += 3
	cpu.step(6)
}

// pha16 push the accumulator onto the stack
func (cpu *CPU) pha16() {
	dataLo, dataHi := bit.SplitUint16(cpu.getCRegister())
	cpu.p16(dataLo, dataHi)
}

// pha8 push the lower bit of the accumulator onto the stack
func (cpu *CPU) pha8() {
	cpu.p8(cpu.getARegister())
}

func (cpu *CPU) pha() {
	if cpu.mFlag {
		cpu.pha8()
	} else {
		cpu.pha16()
	}
}

func (cpu *CPU) op48() {
	cpu.pha()
	cpu.PC++
	cpu.step(4 - bit.BoolToUint16(cpu.mFlag))
}

// PHB instruction
func (cpu *CPU) op8B() {
	cpu.pushStackNew8(cpu.getDBRRegister())
	cpu.PC++
	cpu.step(3)
}

// PHD instruction
func (cpu *CPU) op0B() {
	cpu.pushStackNew16(bit.SplitUint16(cpu.getDRegister()))
	cpu.PC++
	cpu.step(4)
}

// PHK instruction
func (cpu *CPU) op4B() {
	cpu.pushStackNew8(cpu.getKRegister())
	cpu.PC++
	cpu.step(3)
}

func (cpu *CPU) php() {
	P := bit.BoolToUint8(cpu.cFlag)*0x01 +
		bit.BoolToUint8(cpu.zFlag)*0x02 +
		bit.BoolToUint8(cpu.iFlag)*0x04 +
		bit.BoolToUint8(cpu.dFlag)*0x08 +
		bit.BoolToUint8(cpu.mFlag)*0x20 +
		bit.BoolToUint8(cpu.vFlag)*0x40 +
		bit.BoolToUint8(cpu.nFlag)*0x80
	if cpu.eFlag {
		P += bit.BoolToUint8(cpu.bFlag) * 0x10
	} else {
		P += bit.BoolToUint8(cpu.xFlag) * 0x10
	}
	cpu.pushStack(P)
}

func (cpu *CPU) op08() {
	cpu.php()
	cpu.PC++
	cpu.step(3)
}

// phx16 push the X register onto the stack
func (cpu *CPU) phx16() {
	dataLo, dataHi := bit.SplitUint16(cpu.getXRegister())
	cpu.p16(dataLo, dataHi)
}

// phx8 push the lower bit of the X register onto the stack
func (cpu *CPU) phx8() {
	cpu.p8(cpu.getXLRegister())
}

func (cpu *CPU) phx() {
	if cpu.xFlag {
		cpu.phx8()
	} else {
		cpu.phx16()
	}
}

func (cpu *CPU) opDA() {
	cpu.phx()
	cpu.PC++
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag))
}

// phy16 push the Y register onto the stack
func (cpu *CPU) phy16() {
	dataLo, dataHi := bit.SplitUint16(cpu.getYRegister())
	cpu.p16(dataLo, dataHi)
}

// phy8 push the lower bit of the Y register onto the stack
func (cpu *CPU) phy8() {
	cpu.p8(cpu.getYLRegister())
}

func (cpu *CPU) phy() {
	if cpu.xFlag {
		cpu.phy8()
	} else {
		cpu.phy16()
	}
}

func (cpu *CPU) op5A() {
	cpu.phy()
	cpu.PC++
	cpu.step(4 - bit.BoolToUint16(cpu.xFlag))
}

// pla16 pull the accumulator from the stack
func (cpu *CPU) pla16() {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()

	result := bit.JoinUint16(dataLo, dataHi)

	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setCRegister(result)
}

// pla8 pull the lower bits of the accumulator from the stack
func (cpu *CPU) pla8() {
	result := cpu.pullStack()

	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setARegister(result)
}

func (cpu *CPU) pla() {
	if cpu.mFlag {
		cpu.pla8()
	} else {
		cpu.pla16()
	}
}

func (cpu *CPU) op68() {
	cpu.pla()
	cpu.PC++
	cpu.step(5 - bit.BoolToUint16(cpu.mFlag))
}

// PLB instruction
func (cpu *CPU) opAB() {
	cpu.DBR = cpu.pullStackNew8()
	cpu.nFlag = cpu.getDBRRegister()&0x80 != 0
	cpu.zFlag = cpu.getDBRRegister() == 0
	cpu.PC++
	cpu.step(4)
}

// PLD instruction
func (cpu *CPU) op2B() {
	cpu.D = bit.JoinUint16(cpu.pullStackNew16())
	cpu.nFlag = cpu.getDRegister()&0x80 != 0
	cpu.zFlag = cpu.getDRegister() == 0
	cpu.PC++
	cpu.step(4)
}

func (cpu *CPU) plp() {
	P := cpu.pullStack()
	cpu.cFlag = P&0x01 != 0
	cpu.zFlag = P&0x02 != 0
	cpu.iFlag = P&0x04 != 0
	cpu.dFlag = P&0x08 != 0
	cpu.mFlag = P&0x20 != 0
	cpu.vFlag = P&0x40 != 0
	cpu.nFlag = P&0x80 != 0
	if cpu.eFlag {
		cpu.bFlag = P&0x10 != 0
		cpu.setXFlag(true)
		cpu.mFlag = true
	} else {
		cpu.setXFlag(P&0x10 != 0)
	}
}

func (cpu *CPU) op28() {
	cpu.plp()
	cpu.PC++
	cpu.step(4)
}

// plx16 pull the X register from the stack
func (cpu *CPU) plx16() {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()

	result := bit.JoinUint16(dataLo, dataHi)

	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setXRegister(result)
}

// plx8 pull the lower bits of the X register from the stack
func (cpu *CPU) plx8() {
	result := cpu.pullStack()

	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setXLRegister(result)
}

func (cpu *CPU) plx() {
	if cpu.xFlag {
		cpu.plx8()
	} else {
		cpu.plx16()
	}
}

func (cpu *CPU) opFA() {
	cpu.plx()
	cpu.PC++
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

// ply16 pull the Y register from the stack
func (cpu *CPU) ply16() {
	dataLo := cpu.pullStack()
	dataHi := cpu.pullStack()

	result := bit.JoinUint16(dataLo, dataHi)

	cpu.nFlag = result&0x8000 != 0
	cpu.zFlag = result == 0

	cpu.setYRegister(result)
}

// ply8 pull the lower bits of the Y register from the stack
func (cpu *CPU) ply8() {
	result := cpu.pullStack()

	cpu.nFlag = result&0x80 != 0
	cpu.zFlag = result == 0

	cpu.setYLRegister(result)
}

func (cpu *CPU) ply() {
	if cpu.xFlag {
		cpu.ply8()
	} else {
		cpu.ply16()
	}
}

func (cpu *CPU) op7A() {
	cpu.ply()
	cpu.PC++
	cpu.step(5 - bit.BoolToUint16(cpu.xFlag))
}

// trb16 test the bits of the data with the bits of the accumulator then reset the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) trb16(data uint16) uint16 {
	cpu.zFlag = cpu.getCRegister()&data == 0
	return data &^ cpu.getCRegister()

}

// trb8 test the bits of the data with the bits of the accumulator then reset the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) trb8(data uint8) uint8 {
	cpu.zFlag = cpu.getARegister()&data == 0
	return data &^ cpu.getARegister()
}

// trb test the bits of the data with the bits of the accumulator then reset the bits of the data that are ones in the accumulator handling the 8/16 case
func (cpu *CPU) trb(laddr, haddr uint32) {
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.trb8(cpu.memory.GetByte(laddr)), laddr)
	} else {
		result := cpu.trb16(bit.JoinUint16(cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)))
		resultLo, resultHi := bit.SplitUint16(result)
		cpu.memory.SetByte(resultLo, laddr)
		cpu.memory.SetByte(resultHi, haddr)
	}
}

func (cpu *CPU) op14() {
	laddr, haddr := cpu.admDirectP()
	cpu.trb(laddr, haddr)
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}
func (cpu *CPU) op1C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.trb(laddr, haddr)
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

// tsb16 test the bits of the data with the bits of the accumulator then set the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) tsb16(data uint16) uint16 {
	cpu.zFlag = cpu.getCRegister()&data == 0
	return data | cpu.getCRegister()

}

// tsb8 test the bits of the data with the bits of the accumulator then set the bits of the data that are ones in the accumulator 16bit operation
func (cpu *CPU) tsb8(data uint8) uint8 {
	cpu.zFlag = cpu.getARegister()&data == 0
	return data | cpu.getARegister()
}

// tsb test the bits of the data with the bits of the accumulator then set the bits of the data that are ones in the accumulator handling the 8/16 case
func (cpu *CPU) tsb(laddr, haddr uint32) {
	if cpu.mFlag {
		cpu.memory.SetByte(cpu.tsb8(cpu.memory.GetByte(laddr)), laddr)
	} else {
		result := cpu.tsb16(bit.JoinUint16(cpu.memory.GetByte(laddr), cpu.memory.GetByte(haddr)))
		resultLo, resultHi := bit.SplitUint16(result)
		cpu.memory.SetByte(resultLo, laddr)
		cpu.memory.SetByte(resultHi, haddr)
	}
}

func (cpu *CPU) op04() {
	laddr, haddr := cpu.admDirectP()
	cpu.tsb(laddr, haddr)
	cpu.PC += 2
	cpu.step(7 - 2*bit.BoolToUint16(cpu.mFlag) + bit.BoolToUint16(cpu.getDLRegister() == 0))
}

func (cpu *CPU) op0C() {
	laddr, haddr := cpu.admAbsoluteP()
	cpu.tsb(laddr, haddr)
	cpu.PC += 3
	cpu.step(8 - 2*bit.BoolToUint16(cpu.mFlag))
}

func (cpu *CPU) tcd() {
	cpu.D = cpu.C
	// Last bit value
	cpu.nFlag = cpu.D&0x8000 != 0
	cpu.zFlag = cpu.D == 0
}

func (cpu *CPU) op5B() {
	cpu.tcd()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tcs() {
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
	if cpu.eFlag {
		dataLo, _ := bit.SplitUint16(cpu.C)
		cpu.S = bit.JoinUint16(dataLo, 0x01)
	} else {
		cpu.S = cpu.C
	}
}

func (cpu *CPU) op1B() {
	cpu.tcs()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tdc() {
	cpu.C = cpu.D
	// Last bit value
	cpu.nFlag = cpu.C&0x8000 != 0
	cpu.zFlag = cpu.C == 0
}

func (cpu *CPU) op7B() {
	cpu.tdc()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tsc() {
	cpu.C = cpu.S
	// Last bit value
	cpu.nFlag = cpu.S&0x8000 != 0
	cpu.zFlag = cpu.S == 0
}

func (cpu *CPU) op3B() {
	cpu.tsc()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tax() {
	if cpu.xFlag {
		result := cpu.getARegister()
		cpu.setXLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.X = cpu.C
		cpu.nFlag = cpu.X&0x8000 != 0
		cpu.zFlag = cpu.X == 0
	}
}

func (cpu *CPU) opAA() {
	cpu.tax()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tay() {
	if cpu.xFlag {
		result := cpu.getARegister()
		cpu.setYLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.Y = cpu.C
		cpu.nFlag = cpu.Y&0x8000 != 0
		cpu.zFlag = cpu.Y == 0
	}
}

func (cpu *CPU) opA8() {
	cpu.tay()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tsx() {
	if cpu.xFlag {
		result := cpu.getSLRegister()
		cpu.setXLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.X = cpu.S
		cpu.nFlag = cpu.X&0x8000 != 0
		cpu.zFlag = cpu.X == 0
	}
}

func (cpu *CPU) opBA() {
	cpu.tsx()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) txa() {
	if cpu.mFlag {
		result := cpu.getXLRegister()
		cpu.setARegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.C = cpu.X
		cpu.nFlag = cpu.C&0x8000 != 0
		cpu.zFlag = cpu.C == 0
	}
}

func (cpu *CPU) op8A() {
	cpu.txa()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) txs() {
	if cpu.eFlag {
		result := cpu.getXLRegister()
		cpu.setSLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.S = cpu.X
		cpu.nFlag = cpu.S&0x8000 != 0
		cpu.zFlag = cpu.S == 0
	}
}

func (cpu *CPU) op9A() {
	cpu.txs()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) txy() {
	if cpu.xFlag {
		result := cpu.getXLRegister()
		cpu.setYLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.Y = cpu.X
		cpu.nFlag = cpu.Y&0x8000 != 0
		cpu.zFlag = cpu.Y == 0
	}
}

func (cpu *CPU) op9B() {
	cpu.txy()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tya() {
	if cpu.mFlag {
		result := cpu.getYLRegister()
		cpu.setARegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.C = cpu.Y
		cpu.nFlag = cpu.C&0x8000 != 0
		cpu.zFlag = cpu.C == 0
	}
}

func (cpu *CPU) op98() {
	cpu.tya()
	cpu.PC++
	cpu.step(2)
}

func (cpu *CPU) tyx() {
	if cpu.xFlag {
		result := cpu.getYLRegister()
		cpu.setXLRegister(result)
		cpu.nFlag = result&0x80 != 0
		cpu.zFlag = result == 0
	} else {
		cpu.X = cpu.Y
		cpu.nFlag = cpu.X&0x8000 != 0
		cpu.zFlag = cpu.X == 0
	}
}

func (cpu *CPU) opBB() {
	cpu.tyx()
	cpu.PC++
	cpu.step(2)
}

const (
	brkNativeVector    = 0xFFE6
	copNativeVector    = 0xFFE4
	brkEmulationVector = 0xFFFE
	copEmulationVector = 0xFFF4
)

func (cpu *CPU) brk() {
	laddr, haddr := bit.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.bFlag = true
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, brkEmulationVector)
		haddr := cpu.memory.GetByteBank(0x00, brkEmulationVector+1)
		cpu.PC = bit.JoinUint16(laddr, haddr)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, brkNativeVector)
		haddr := cpu.memory.GetByteBank(0x00, brkNativeVector+1)
		cpu.PC = bit.JoinUint16(laddr, haddr)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.step(8 - bit.BoolToUint16(cpu.eFlag))
}

func (cpu *CPU) op00() {
	cpu.brk()
}

func (cpu *CPU) cop() {
	laddr, haddr := bit.SplitUint16(cpu.getPCRegister() + 2)
	if cpu.eFlag {
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, copEmulationVector)
		haddr := cpu.memory.GetByteBank(0x00, copEmulationVector+1)
		cpu.PC = bit.JoinUint16(laddr, haddr)
	} else {
		cpu.pushStack(cpu.getKRegister())
		cpu.pushStack(haddr)
		cpu.pushStack(laddr)
		cpu.php()
		cpu.K = 0x00
		laddr := cpu.memory.GetByteBank(0x00, copNativeVector)
		haddr := cpu.memory.GetByteBank(0x00, copNativeVector+1)
		cpu.PC = bit.JoinUint16(laddr, haddr)
	}
	cpu.dFlag = false
	cpu.iFlag = true
	cpu.step(8 - bit.BoolToUint16(cpu.eFlag))

}

func (cpu *CPU) op02() {
	cpu.cop()
}

func (cpu *CPU) registerOpcodes() {
	cpu.opcodes[0x0] = cpu.op00
	cpu.opcodes[0x1] = cpu.op01
	cpu.opcodes[0x2] = cpu.op02
	cpu.opcodes[0x3] = cpu.op03
	cpu.opcodes[0x4] = cpu.op04
	cpu.opcodes[0x5] = cpu.op05
	cpu.opcodes[0x6] = cpu.op06
	cpu.opcodes[0x7] = cpu.op07
	cpu.opcodes[0x8] = cpu.op08
	cpu.opcodes[0x9] = cpu.op09
	cpu.opcodes[0xA] = cpu.op0A
	cpu.opcodes[0xB] = cpu.op0B
	cpu.opcodes[0xC] = cpu.op0C
	cpu.opcodes[0xD] = cpu.op0D
	cpu.opcodes[0xE] = cpu.op0E
	cpu.opcodes[0xF] = cpu.op0F
	cpu.opcodes[0x10] = cpu.op10
	cpu.opcodes[0x11] = cpu.op11
	cpu.opcodes[0x12] = cpu.op12
	cpu.opcodes[0x13] = cpu.op13
	cpu.opcodes[0x14] = cpu.op14
	cpu.opcodes[0x15] = cpu.op15
	cpu.opcodes[0x16] = cpu.op16
	cpu.opcodes[0x17] = cpu.op17
	cpu.opcodes[0x18] = cpu.op18
	cpu.opcodes[0x19] = cpu.op19
	cpu.opcodes[0x1A] = cpu.op1A
	cpu.opcodes[0x1B] = cpu.op1B
	cpu.opcodes[0x1C] = cpu.op1C
	cpu.opcodes[0x1D] = cpu.op1D
	cpu.opcodes[0x1E] = cpu.op1E
	cpu.opcodes[0x1F] = cpu.op1F
	cpu.opcodes[0x20] = cpu.op20
	cpu.opcodes[0x21] = cpu.op21
	cpu.opcodes[0x22] = cpu.op22
	cpu.opcodes[0x23] = cpu.op23
	cpu.opcodes[0x24] = cpu.op24
	cpu.opcodes[0x25] = cpu.op25
	cpu.opcodes[0x26] = cpu.op26
	cpu.opcodes[0x27] = cpu.op27
	cpu.opcodes[0x28] = cpu.op28
	cpu.opcodes[0x29] = cpu.op29
	cpu.opcodes[0x2A] = cpu.op2A
	cpu.opcodes[0x2B] = cpu.op2B
	cpu.opcodes[0x2C] = cpu.op2C
	cpu.opcodes[0x2D] = cpu.op2D
	cpu.opcodes[0x2E] = cpu.op2E
	cpu.opcodes[0x2F] = cpu.op2F
	cpu.opcodes[0x30] = cpu.op30
	cpu.opcodes[0x31] = cpu.op31
	cpu.opcodes[0x32] = cpu.op32
	cpu.opcodes[0x33] = cpu.op33
	cpu.opcodes[0x34] = cpu.op34
	cpu.opcodes[0x35] = cpu.op35
	cpu.opcodes[0x36] = cpu.op36
	cpu.opcodes[0x37] = cpu.op37
	cpu.opcodes[0x38] = cpu.op38
	cpu.opcodes[0x39] = cpu.op39
	cpu.opcodes[0x3A] = cpu.op3A
	cpu.opcodes[0x3B] = cpu.op3B
	cpu.opcodes[0x3C] = cpu.op3C
	cpu.opcodes[0x3D] = cpu.op3D
	cpu.opcodes[0x3E] = cpu.op3E
	cpu.opcodes[0x3F] = cpu.op3F
	cpu.opcodes[0x40] = cpu.op40
	cpu.opcodes[0x41] = cpu.op41
	cpu.opcodes[0x42] = cpu.op42
	cpu.opcodes[0x43] = cpu.op43
	cpu.opcodes[0x44] = cpu.op44
	cpu.opcodes[0x45] = cpu.op45
	cpu.opcodes[0x46] = cpu.op46
	cpu.opcodes[0x47] = cpu.op47
	cpu.opcodes[0x48] = cpu.op48
	cpu.opcodes[0x49] = cpu.op49
	cpu.opcodes[0x4A] = cpu.op4A
	cpu.opcodes[0x4B] = cpu.op4B
	cpu.opcodes[0x4C] = cpu.op4C
	cpu.opcodes[0x4D] = cpu.op4D
	cpu.opcodes[0x4E] = cpu.op4E
	cpu.opcodes[0x4F] = cpu.op4F
	cpu.opcodes[0x50] = cpu.op50
	cpu.opcodes[0x51] = cpu.op51
	cpu.opcodes[0x52] = cpu.op52
	cpu.opcodes[0x53] = cpu.op53
	cpu.opcodes[0x54] = cpu.op54
	cpu.opcodes[0x55] = cpu.op55
	cpu.opcodes[0x56] = cpu.op56
	cpu.opcodes[0x57] = cpu.op57
	cpu.opcodes[0x58] = cpu.op58
	cpu.opcodes[0x59] = cpu.op59
	cpu.opcodes[0x5A] = cpu.op5A
	cpu.opcodes[0x5B] = cpu.op5B
	cpu.opcodes[0x5C] = cpu.op5C
	cpu.opcodes[0x5D] = cpu.op5D
	cpu.opcodes[0x5E] = cpu.op5E
	cpu.opcodes[0x5F] = cpu.op5F
	cpu.opcodes[0x60] = cpu.op60
	cpu.opcodes[0x61] = cpu.op61
	cpu.opcodes[0x62] = cpu.op62
	cpu.opcodes[0x63] = cpu.op63
	cpu.opcodes[0x64] = cpu.op64
	cpu.opcodes[0x65] = cpu.op65
	cpu.opcodes[0x66] = cpu.op66
	cpu.opcodes[0x67] = cpu.op67
	cpu.opcodes[0x68] = cpu.op68
	cpu.opcodes[0x69] = cpu.op69
	cpu.opcodes[0x6A] = cpu.op6A
	cpu.opcodes[0x6B] = cpu.op6B
	cpu.opcodes[0x6C] = cpu.op6C
	cpu.opcodes[0x6D] = cpu.op6D
	cpu.opcodes[0x6E] = cpu.op6E
	cpu.opcodes[0x6F] = cpu.op6F
	cpu.opcodes[0x70] = cpu.op70
	cpu.opcodes[0x71] = cpu.op71
	cpu.opcodes[0x72] = cpu.op72
	cpu.opcodes[0x73] = cpu.op73
	cpu.opcodes[0x74] = cpu.op74
	cpu.opcodes[0x75] = cpu.op75
	cpu.opcodes[0x76] = cpu.op76
	cpu.opcodes[0x77] = cpu.op77
	cpu.opcodes[0x78] = cpu.op78
	cpu.opcodes[0x79] = cpu.op79
	cpu.opcodes[0x7A] = cpu.op7A
	cpu.opcodes[0x7B] = cpu.op7B
	cpu.opcodes[0x7C] = cpu.op7C
	cpu.opcodes[0x7D] = cpu.op7D
	cpu.opcodes[0x7E] = cpu.op7E
	cpu.opcodes[0x7F] = cpu.op7F
	cpu.opcodes[0x80] = cpu.op80
	cpu.opcodes[0x81] = cpu.op81
	cpu.opcodes[0x82] = cpu.op82
	cpu.opcodes[0x83] = cpu.op83
	cpu.opcodes[0x84] = cpu.op84
	cpu.opcodes[0x85] = cpu.op85
	cpu.opcodes[0x86] = cpu.op86
	cpu.opcodes[0x87] = cpu.op87
	cpu.opcodes[0x88] = cpu.op88
	cpu.opcodes[0x89] = cpu.op89
	cpu.opcodes[0x8A] = cpu.op8A
	cpu.opcodes[0x8B] = cpu.op8B
	cpu.opcodes[0x8C] = cpu.op8C
	cpu.opcodes[0x8D] = cpu.op8D
	cpu.opcodes[0x8E] = cpu.op8E
	cpu.opcodes[0x8F] = cpu.op8F
	cpu.opcodes[0x90] = cpu.op90
	cpu.opcodes[0x91] = cpu.op91
	cpu.opcodes[0x92] = cpu.op92
	cpu.opcodes[0x93] = cpu.op93
	cpu.opcodes[0x94] = cpu.op94
	cpu.opcodes[0x95] = cpu.op95
	cpu.opcodes[0x96] = cpu.op96
	cpu.opcodes[0x97] = cpu.op97
	cpu.opcodes[0x98] = cpu.op98
	cpu.opcodes[0x99] = cpu.op99
	cpu.opcodes[0x9A] = cpu.op9A
	cpu.opcodes[0x9B] = cpu.op9B
	cpu.opcodes[0x9C] = cpu.op9C
	cpu.opcodes[0x9D] = cpu.op9D
	cpu.opcodes[0x9E] = cpu.op9E
	cpu.opcodes[0x9F] = cpu.op9F
	cpu.opcodes[0xA0] = cpu.opA0
	cpu.opcodes[0xA1] = cpu.opA1
	cpu.opcodes[0xA2] = cpu.opA2
	cpu.opcodes[0xA3] = cpu.opA3
	cpu.opcodes[0xA4] = cpu.opA4
	cpu.opcodes[0xA5] = cpu.opA5
	cpu.opcodes[0xA6] = cpu.opA6
	cpu.opcodes[0xA7] = cpu.opA7
	cpu.opcodes[0xA8] = cpu.opA8
	cpu.opcodes[0xA9] = cpu.opA9
	cpu.opcodes[0xAA] = cpu.opAA
	cpu.opcodes[0xAB] = cpu.opAB
	cpu.opcodes[0xAC] = cpu.opAC
	cpu.opcodes[0xAD] = cpu.opAD
	cpu.opcodes[0xAE] = cpu.opAE
	cpu.opcodes[0xAF] = cpu.opAF
	cpu.opcodes[0xB0] = cpu.opB0
	cpu.opcodes[0xB1] = cpu.opB1
	cpu.opcodes[0xB2] = cpu.opB2
	cpu.opcodes[0xB3] = cpu.opB3
	cpu.opcodes[0xB4] = cpu.opB4
	cpu.opcodes[0xB5] = cpu.opB5
	cpu.opcodes[0xB6] = cpu.opB6
	cpu.opcodes[0xB7] = cpu.opB7
	cpu.opcodes[0xB8] = cpu.opB8
	cpu.opcodes[0xB9] = cpu.opB9
	cpu.opcodes[0xBA] = cpu.opBA
	cpu.opcodes[0xBB] = cpu.opBB
	cpu.opcodes[0xBC] = cpu.opBC
	cpu.opcodes[0xBD] = cpu.opBD
	cpu.opcodes[0xBE] = cpu.opBE
	cpu.opcodes[0xBF] = cpu.opBF
	cpu.opcodes[0xC0] = cpu.opC0
	cpu.opcodes[0xC1] = cpu.opC1
	cpu.opcodes[0xC2] = cpu.opC2
	cpu.opcodes[0xC3] = cpu.opC3
	cpu.opcodes[0xC4] = cpu.opC4
	cpu.opcodes[0xC5] = cpu.opC5
	cpu.opcodes[0xC6] = cpu.opC6
	cpu.opcodes[0xC7] = cpu.opC7
	cpu.opcodes[0xC8] = cpu.opC8
	cpu.opcodes[0xC9] = cpu.opC9
	cpu.opcodes[0xCA] = cpu.opCA
	cpu.opcodes[0xCB] = cpu.opCB
	cpu.opcodes[0xCC] = cpu.opCC
	cpu.opcodes[0xCD] = cpu.opCD
	cpu.opcodes[0xCE] = cpu.opCE
	cpu.opcodes[0xCF] = cpu.opCF
	cpu.opcodes[0xD0] = cpu.opD0
	cpu.opcodes[0xD1] = cpu.opD1
	cpu.opcodes[0xD2] = cpu.opD2
	cpu.opcodes[0xD3] = cpu.opD3
	cpu.opcodes[0xD4] = cpu.opD4
	cpu.opcodes[0xD5] = cpu.opD5
	cpu.opcodes[0xD6] = cpu.opD6
	cpu.opcodes[0xD7] = cpu.opD7
	cpu.opcodes[0xD8] = cpu.opD8
	cpu.opcodes[0xD9] = cpu.opD9
	cpu.opcodes[0xDA] = cpu.opDA
	cpu.opcodes[0xDB] = cpu.opDB
	cpu.opcodes[0xDC] = cpu.opDC
	cpu.opcodes[0xDD] = cpu.opDD
	cpu.opcodes[0xDE] = cpu.opDE
	cpu.opcodes[0xDF] = cpu.opDF
	cpu.opcodes[0xE0] = cpu.opE0
	cpu.opcodes[0xE1] = cpu.opE1
	cpu.opcodes[0xE2] = cpu.opE2
	cpu.opcodes[0xE3] = cpu.opE3
	cpu.opcodes[0xE4] = cpu.opE4
	cpu.opcodes[0xE5] = cpu.opE5
	cpu.opcodes[0xE6] = cpu.opE6
	cpu.opcodes[0xE7] = cpu.opE7
	cpu.opcodes[0xE8] = cpu.opE8
	cpu.opcodes[0xE9] = cpu.opE9
	cpu.opcodes[0xEA] = cpu.opEA
	cpu.opcodes[0xEB] = cpu.opEB
	cpu.opcodes[0xEC] = cpu.opEC
	cpu.opcodes[0xED] = cpu.opED
	cpu.opcodes[0xEE] = cpu.opEE
	cpu.opcodes[0xEF] = cpu.opEF
	cpu.opcodes[0xF0] = cpu.opF0
	cpu.opcodes[0xF1] = cpu.opF1
	cpu.opcodes[0xF2] = cpu.opF2
	cpu.opcodes[0xF3] = cpu.opF3
	cpu.opcodes[0xF4] = cpu.opF4
	cpu.opcodes[0xF5] = cpu.opF5
	cpu.opcodes[0xF6] = cpu.opF6
	cpu.opcodes[0xF7] = cpu.opF7
	cpu.opcodes[0xF8] = cpu.opF8
	cpu.opcodes[0xF9] = cpu.opF9
	cpu.opcodes[0xFA] = cpu.opFA
	cpu.opcodes[0xFB] = cpu.opFB
	cpu.opcodes[0xFC] = cpu.opFC
	cpu.opcodes[0xFD] = cpu.opFD
	cpu.opcodes[0xFE] = cpu.opFE
	cpu.opcodes[0xFF] = cpu.opFF
}
