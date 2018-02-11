package cpu

import (
	"unsafe"
)

func (cpu CPU) admDirect8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL)
	return cpu.memory.GetByte(address)
}

func (cpu CPU) admDirect(LL uint8) (uint8, uint8) {
	laddress := cpu.getDRegister() + uint16(LL)
	var uint32 address
	*(*uint16)(unsafe.Pointer(&address)) = laddress
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectX8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
	return cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectX(LL uint8) (uint8, uint8) {
	laddress := cpu.getDRegister() + uint16(LL) + cpu.getXRegister()
	var uint32 address
	*(*uint16)(unsafe.Pointer(&address)) = laddress
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectY8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getYLRegister())
	return cpu.memory.GetByte(address)
}

func (cpu CPU) admDirectY(LL uint8) (uint8, uint8) {
	laddress := cpu.getDRegister() + uint16(LL) + cpu.getYRegister()
	var uint32 address
	*(*uint16)(unsafe.Pointer(&address)) = laddress
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admPDirect8(LL uint8) uint8 {
	address := readUint32(0x00, cpu.getDHRegister(), LL)
	ll := cpu.memory.GetByte(address)
	hh := cpu.memory.GetByte(address + 1)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

func (cpu CPU) admPDirect(LL uint8) (uint8, uint8) {
	laddress := cpu.getDRegister() + uint16(LL)
	var uint32 address
	*(*uint16)(unsafe.Pointer(&address)) = laddress
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}

func (cpu CPU) admBDirect(LL uint8) (uint8, uint8) {
	laddress := cpu.getDRegister() + uint16(LL)
	var uint32 address
	*(*uint16)(unsafe.Pointer(&address)) = laddress
	ll := cpu.memory.GetByte(address)
	*(*uint16)(unsafe.Pointer(&address)) = laddress + 1
	mm := cpu.memory.GetByte(address)
	*(*uint16)(unsafe.Pointer(&address)) = laddress + 2
	hh := cpu.memory.GetByte(address)
	pointer := readUint32(hh, mm, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

func (cpu CPU) admPDirectX8(LL uint8) (uint8, uint8) {
	address := readUint32(0x00, cpu.getDHRegister(), LL+cpu.getXLRegister())
	ll := cpu.memory.GetByte(address)
	hh := cpu.memory.GetByte(address + 1)
	pointer := readUint32(cpu.getDBRRegister(), hh, ll)
	return cpu.memory.GetByte(pointer + 1), cpu.memory.GetByte(pointer)
}

func (cpu CPU) admPDirectX(LL uint8) (uint8, uint8) {
	laddress := cpu.getDRegister() + uint16(LL) + cpu.getXRegister()
	var uint32 address
	*(*uint16)(unsafe.Pointer(&address)) = laddress
	return cpu.memory.GetByte(address + 1), cpu.memory.GetByte(address)
}
