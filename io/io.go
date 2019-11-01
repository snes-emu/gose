package io

const UNUSED_REGISTER = "UNUSED_REGISTER"

// Register represents an read/write io register
type Register struct {
	Read  func() uint8
	Write func(uint8)
	Name  string
}

type RegisterFactory struct {
	hook func(string)
}

func NewRegisterFactory() *RegisterFactory {
	return NewRegisterFactoryWithHook(nil)
}

func NewRegisterFactoryWithHook(hook func(string)) *RegisterFactory {
	return &RegisterFactory{
		hook: hook,
	}
}

func (rf *RegisterFactory) NewRegister(read func() uint8, write func(uint8), name ...string) *Register {
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
	rread := read
	rwrite := write
	if rf.hook != nil {
		rread = func() uint8 {
			rf.hook(regname)
			return read()
		}

		rwrite = func(data uint8) {
			rf.hook(regname)
			write(data)
		}
	}

	r := &Register{
		rread,
		rwrite,
		regname,
	}
	return r
}

func unusedRead() uint8 {
	return 0
}

func unusedWrite(_ uint8) {
}
