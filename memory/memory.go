package memory

type Memory struct {
	bytes []uint8
}

func (memory Memory) GetByte(index uint32) uint8 {
	return memory.bytes[index]
}
