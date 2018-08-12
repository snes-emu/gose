package io

// Register represents an read/write io register
type Register struct {
	Read  func() uint8
	Write func(uint8)
	Name  string
}

const UNUSED_REGISTER = "UNUSED_REGISTER"

func NewRegister(read func() uint8, write func(uint8), name ...string) *Register {
	regname := UNUSED_REGISTER

	if read == nil {
		read = unusedRead
	}
	if write == nil {
		write = unusedWrite
	}

	if len(name) > 0 {
		regname = name[0]
	}

	r := &Register{
		read,
		write,
		regname,
	}
	return r
}
func unusedRead() uint8 {
	return 0
}

func unusedWrite(_ uint8) {
}
