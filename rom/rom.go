package rom

import "fmt"

const (
	// LoROM type
	LoROM = iota
	// HiROM type
	HiROM
	// ExLoROM type
	ExLoROM
	// ExHiROM type
	ExHiROM

	maxSRAMSize = 9
)

// ROM struct
type ROM struct {
	Data     []byte // Raw bytes of the rom
	Title    string // Rom title
	size     uint   // Size of the rom
	isFast   bool   // Whether or not the ROM is of type Fast
	SRAMSize uint   // SRAM size
	Type     uint   // Type of the Rom (LoROM, HiROM, ExLoROM, ExHiROM)
}

// ParseROM parses a ROM file representation in bytes and return a representation
func ParseROM(data []byte) (*ROM, error) {
	rom := &ROM{
		Data: data,
	}

	// SMC header should be of len 0 or 512
	smcHeaderSize := len(rom.Data) % 0x400
	if smcHeaderSize != 0 && smcHeaderSize != 512 {
		return nil, fmt.Errorf("The smc header of this rom is not conventional (len: %v)", smcHeaderSize)
	}

	// Remove smc header
	rom.Data = rom.Data[smcHeaderSize:]

	// Set rom parameters
	var (
		headerAddr int
		romType    uint
	)
	if rom.isLo() {
		headerAddr = 0x7fc0
		romType = LoROM

	} else {
		headerAddr = 0xffc0
		romType = HiROM
	}

	rom.Title = string(rom.Data[headerAddr : headerAddr+21])
	rom.isFast = rom.Data[headerAddr+21]&0x30 != 0
	rom.size = 0x400 << rom.Data[headerAddr+23]
	sramSize := rom.Data[headerAddr+24]
	//sram is between 0 and 512kB
	if sramSize != 0 {
		if sramSize > maxSRAMSize {
			sramSize = maxSRAMSize
		}
		rom.SRAMSize = 0x400 << sramSize
	}
	rom.Type = romType

	return rom, nil
}

// isLo checks if the ROM is of type LoROM
func (rom *ROM) isLo() bool {
	for _, c := range rom.Data[0x7fc0:0x7fd4] {
		// check if all chars are ascii characters
		if c > 0x7f || c < 0x1f {
			return false
		}
	}

	return true
}
