package rom

import "fmt"

type ROM struct {
	data  []byte
	Title string
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

	// Set rom Title
	if rom.isLo() {
		rom.Title = string(rom.data[0x7fc0:0x7fd4])
	} else {
		rom.Title = string(rom.data[0xffc0:0xffd4])
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
