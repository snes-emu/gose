package io

// Register represents an read/write io register
type Register struct {
	read  func() uint8
	write func(uint8)
}

func NewRegister(read func() uint8, write func(uint8)) *Register {
	if read == nil {
		read = unusedRead
	}
	if write == nil {
		write = unusedWrite
	}
	r := &Register{
		read,
		write,
	}
	return r
}
func unusedRead() uint8 {
	return 0
}

func unusedWrite(_ uint8) {
}
