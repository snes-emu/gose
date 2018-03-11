package utils

import (
	"bytes"
	"encoding/binary"
)

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
	var ret uint32
	buf := bytes.NewBuffer([]byte{LL, MM, HH, 0x00})
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func JoinUint16(LL uint8, HH uint8) uint16 {
	var ret uint16
	buf := bytes.NewBuffer([]byte{LL, HH})
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func SplitUint16(MM uint16) (uint8, uint8) {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &MM)
	ret := buf.Bytes()
	return ret[0], ret[1]
}

func SplitUint32(MM uint32) (uint8, uint8, uint8) {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &MM)
	ret := buf.Bytes()
	return ret[0], ret[1], ret[2]
}
