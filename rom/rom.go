package rom

import "fmt"

const (
	LoROM = iota
	HiROM
	ExLoROM
	ExHiROM
)

type ROM struct {
	data     []byte // Raw bytes of the rom
	Title    string // Rom title
	size     uint   // Size of the rom
	isFast   bool   // Whether or not the ROM is of type Fast
	sramSize uint   // SRAM size
	Type     uint   // Type of the Rom (LoROM, HiROM, ExLoROM, ExHiROM)
}

// ParseROM parses a ROM file representation in bytes and return a representation
func ParseROM(data []byte) (*ROM, error) {
	rom := &ROM{
		data: data,
	}

	// SMC header should be of len 0 or 512
	smcHeaderSize := len(rom.data) % 0x400
	if smcHeaderSize != 0 && smcHeaderSize != 512 {
		return nil, fmt.Errorf("The smc header of this rom is not conventional (len: %v)", smcHeaderSize)
	}

	// Remove smc header
	rom.data = rom.data[smcHeaderSize:]

	// Set rom parameters
	if rom.isLo() {
		rom.Title = string(rom.data[0x7fc0:0x7fd4])
		rom.isFast = rom.data[0x7fd5]&0x30 != 0
		rom.size = 0x400 << rom.data[0x7fd7]
		rom.sramSize = 0x400 << rom.data[0x7fd8]
		rom.Type = LoROM
	} else {
		rom.Title = string(rom.data[0xffc0:0xffd4])
		rom.isFast = rom.data[0xffd5]&0x30 != 0
		rom.size = 0x400 << rom.data[0xffd7]
		rom.sramSize = 0x400 << rom.data[0xffd8]
		rom.Type = HiROM
	}

	return rom, nil
}

// isLo checks if the ROM is of type LoROM
func (rom ROM) isLo() bool {
	for _, c := range rom.data[0x7fc0:0x7fd4] {
		// check if all chars are ascii characters
		if c > 0x7f || c < 0x1f {
			return false
		}
	}

	return true
}
