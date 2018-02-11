package memory

type Memory struct {
	bytes []uint8
}

func (memory Memory) GetByte(index uint32) uint8 {
	return memory.bytes[index]
}

func (memory Memory) GetByteBank(K uint8, address uint16) uint8 {
	return memory.bytes[uint32(K)<<16 + +uint32(address)]
}
