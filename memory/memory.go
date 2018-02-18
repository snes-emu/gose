package memory

const bankNumber = 256
const offsetMask = 0xFFFF
const sramSize = 0x8000
const wramSize = 0x20000
const loROM = 0
const hiROM = 1
const exLoROM = 2
const exHiROM = 3

// Memory struct containing SNES working RAM, cartridge static RAM, special hardware registers and default memory buffer for ROM
type Memory struct {
	main    [bankNumber][]uint8
	sram    [sramSize]uint8
	wram    [wramSize]uint8
	romType int
}

// New creates a Memory struct and initialize it
func New() *Memory {
	memory := &Memory{}
	for bank := 0; bank < bankNumber; bank++ {
		memory.main[bank] = make([]byte, 1<<16)
	}
	return memory
}

// LoadROM takes a memory buffer and load it into memory depending ROM type
func (memory *Memory) LoadROM(ROM []byte) {

	// only LoROM for now
	if memory.romType == loROM {
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

//GetByte gets a byte by its complete address
func (memory Memory) GetByte(index uint32) uint8 {
	K := index >> 16
	offset := index & offsetMask
	return memory.GetByteBank(uint8(K), uint16(offset))
}

//GetByteBank gets a byte by memory bank and offset
func (memory Memory) GetByteBank(K uint8, offset uint16) uint8 {
	switch memory.romType {
	case loROM:
		if K < 0x40 || (0x7F < K && K < 0xC0) {
			if offset < 0x2000 {
				return memory.wram[offset]
			}
		} else if offset < 0x8000 && ((0x6F < K && K < 0x7E) || (0xEF < K && K < 0xFE)) {
			return memory.sram[offset]
		} else if K > 0x7D && K < 0x80 {
			return memory.wram[offset+uint16(K)-0x7E]
		} else if 0xFD < K && offset < 0x8000 {
			return memory.sram[offset]
		}
		return memory.main[K][offset]
	default:
		return 0x00
	}
}

//SetByte sets a byte by its complete address
func (memory *Memory) SetByte(value uint8, index uint32) {
	K := index >> 16
	offset := index & offsetMask
	memory.SetByteBank(value, uint8(K), uint16(offset))
}

//SetByteBank sets a byte by memory bank and offset
func (memory *Memory) SetByteBank(value uint8, K uint8, offset uint16) {
	if K < 0x40 {
		if offset < 0x2000 {
			memory.wram[offset] = value
		}
	} else if K > 0x6F && K < 0x7E && offset < 0x8000 {
		memory.sram[offset] = value
	} else if K > 0x7D && K < 0x80 {
		memory.wram[offset+uint16(K)-0x7E] = value
	}
}
