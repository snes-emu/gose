package rom

type ROM struct {
	rawData   []byte
	GameTitle string
}

// ParseROM parses a ROM file representation in bytes and return a representation
func ParseROM(data []byte) *ROM {
	rom := &ROM{
		rawData: data,
	}

	title := data[0xfc0:0xfd4]

	rom.GameTitle = string(title)

	return rom
}
