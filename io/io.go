package io

// Register represents an read/write io register
type Register func(uint8) uint8

// UnusedRegister represents an unused register in the memory map
func UnusedRegister(_ uint8) uint8 {
	return 0
}
