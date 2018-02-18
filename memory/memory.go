package memory

const bankNumber = 256
const offsetMask = 0xFFFF
const sramSize = 0x8000
const wramSize = 0x20000

type Memory struct {
	main [bankNumber][]uint8
	sram [sramSize]uint8
	wram [wramSize]uint8
}

func New() *Memory {
	memory := &Memory{}
	for bank := 0; bank < bankNumber; bank++ {
		memory.main[bank] = make([]byte, 1<<16)
	}
	return memory
}

func (memory *Memory) LoadROM(ROM []byte, ROMType int) {

	// only LoROM for now
	if ROMType == 0 {
		for bank := 0x00; bank < 0x80; bank++ {
			memory.main[bank] = make([]byte, 0xFFFF+1)
			for offset := 0x8000; offset < 0x10000; offset++ {
				memory.main[bank][offset] = ROM[offset+(bank-1)*0x8000]
			}
		}

		for bank := 0x80; bank < 0x100; bank++ {
			memory.main[bank] = memory.main[bank-0x80]
		}
	}
}

func (memory Memory) GetByte(index uint32) uint8 {
	return memory.main[index>>16][index&offsetMask]
}

func (memory Memory) GetByteBank(K uint8, address uint16) uint8 {
	return memory.main[K][address]
}

func (memory Memory) SetByte(value uint8, index uint32) {
	memory.main[index>>16][index&offsetMask] = value
}

func (memory Memory) SetByteBank(value uint8, K uint8, address uint16) {
	memory.main[K][address] = value
}
