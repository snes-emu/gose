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

func ReadUint32(HH uint8, MM uint8, LL uint8) uint32 {
	var ret uint32
	buf := bytes.NewBuffer([]byte{LL, MM, HH})
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func ReadUint16(HH uint8, LL uint8) uint16 {
	var ret uint16
	buf := bytes.NewBuffer([]byte{LL, HH})
	binary.Read(buf, binary.LittleEndian, &ret)
	return ret
}

func WriteUint16(MM uint16) (uint8, uint8) {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &MM)
	ret := buf.Bytes()
	return ret[1], ret[0]
}

func WriteUint32(MM uint32) (uint8, uint8, uint8) {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &MM)
	ret := buf.Bytes()
	return ret[2], ret[1], ret[0]
}
