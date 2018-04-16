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

func JoinUint32(LL uint8, MM uint8, HH uint8) uint32 {
	return uint32(LL) | uint32(MM)<<8 | uint32(HH)<<16
}

func JoinUint16(LL uint8, HH uint8) uint16 {
	return uint16(LL) | uint16(HH)<<8
}

func SplitUint16(MM uint16) (uint8, uint8) {
	return uint8(MM), uint8(MM >> 8)
}

func SplitUint32(MM uint32) (uint8, uint8, uint8) {
	return uint8(MM), uint8(MM >> 8), uint8(MM >> 16)
}
