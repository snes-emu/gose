package utils

// BoolToUint16 provides a conversion from bool to uint16
var BoolToUint16 = map[bool]uint16{
	true:  1,
	false: 0,
}

// BoolToUint8 provides a conversion from bool to uint8
var BoolToUint8 = map[bool]uint8{
	true:  1,
	false: 0,
}

// GetBit8  nth bit of a uint8
func GetBit8(m, n uint8) uint8 {
	return (m >> n) % 2
}

// GetBit16  nth bit of a uint8
func GetBit16(m, n uint16) uint16 {
	return (m >> n) % 2
}
