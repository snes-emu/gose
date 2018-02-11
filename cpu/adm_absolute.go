package cpu

import (
	"bytes"
	"encoding/binary"
)

func readUint32(HH uint8, MM uint8, LL uint8) uint32 {
	var ret uint32
	buf := bytes.NewBuffer([]byte{LL, MM, HH})
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

// ABSOLUTE addressing mode to use only for JMP	and JSR instructions
func (cpu CPU) admAbsoluteJ(HH uint8, LL uint8) uint32 {
	return readUint32(cpu.getKRegister(), HH, LL)
}

//ABSOLUTE addressing mode
func (cpu CPU) admAbsolute(HH uint8, LL uint8) (uint8, uint8) {
	address := readUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteX(HH uint8, LL uint8) (uint8, uint8) {
	address := readUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + uint32(cpu.getXRegister()) + 1), cpu.memory.GetByte(address + uint32(cpu.getXRegister()))
}

// ABSOLUTE,X addressing mode
func (cpu CPU) admAbsoluteY(HH uint8, LL uint8) (uint8, uint8) {
	address := readUint32(cpu.getKRegister(), HH, LL)
	return cpu.memory.GetByte(address + uint32(cpu.getYRegister()) + 1), cpu.memory.GetByte(address + uint32(cpu.getYRegister()))
}

// (ABSOLUTE) addressing mode
func (cpu CPU) admPAbsolute(HH uint8, LL uint8) uint32 {
	address := readUint32(0x00, HH, LL)
	return readUint32(cpu.getKRegister(), cpu.memory.GetByte(address+1), cpu.memory.GetByte(address))
}

// [ABSOLUTE] addressing mode
func (cpu CPU) admBAbsolute(HH uint8, LL uint8) uint32 {
	address := readUint32(0x00, HH, LL)
	return readUint32(cpu.memory.GetByte(address+2), cpu.memory.GetByte(address+1), cpu.memory.GetByte(address))
}

// (ABSOLUTE,X) addressing mode
func (cpu CPU) admPAbsoluteX(HH uint8, LL uint8) uint32 {
	address := readUint32(0x00, HH, LL)
	return readUint32(cpu.getKRegister(), cpu.memory.GetByte(address+uint32(cpu.getXRegister())+1), cpu.memory.GetByte(address+uint32(cpu.getXRegister())))
}
