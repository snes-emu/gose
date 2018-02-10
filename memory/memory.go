package memory

type Memory struct {
	bytes []uint8
}

func (memory Memory) getByte(index uint16) uint8 {
	return memory.bytes[index]
}
