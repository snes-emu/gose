package memory

const bankNumber = 256
const offsetMask = 0xFFFF

type Memory struct {
	bank [bankNumber][]uint8
}

func New() *Memory {
	memory := &Memory{}
	for bank := 0; bank < bankNumber; bank++ {
		memory.bank[bank] = make([]byte, 1<<16)
	}
	return memory
}

func (memory Memory) GetByte(index uint32) uint8 {
	return memory.bank[index>>16][index&offsetMask]
}

func (memory Memory) GetByteBank(K uint8, address uint16) uint8 {
	return memory.bank[K][address]
}

func (memory Memory) SetByte(value uint8, index uint32) {
	memory.bank[index>>16][index&offsetMask] = value
}

func (memory Memory) SetByteBank(value uint8, K uint8, address uint16) {
	memory.bank[K][address] = value
}
