package bit

// BoolToUint16 provides a conversion from bool to uint16
func BoolToUint16(f bool) uint16 {
	if f {
		return 1
	}

	return 0
}

// BoolToUint8 provides a conversion from bool to uint8
func BoolToUint8(f bool) uint8 {
	if f {
		return 1
	}

	return 0
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

func SetLowByte(w uint16, ll uint8) uint16 {
	return (w & 0xff00) | uint16(ll)
}

func SetHighByte(w uint16, ll uint8) uint16 {
	return (w & 0x00ff) | uint16(ll)<<8
}

func LowByte(x uint16) uint8 {
	return uint8(x & 0xff)
}

func HighByte(x uint16) uint8 {
	return uint8(x >> 8)
}
