package io

const UNUSED_REGISTER = "UNUSED_REGISTER"

const READ = "read"
const WRITE = "write"

type registerHook func(reg string, typ string, data uint8)

// Register represents an read/write io register
type Register struct {
	Read  func() uint8
	Write func(uint8)
	Name  string
}

type RegisterFactory struct {
	hook registerHook
}

func NewRegisterFactory() *RegisterFactory {
	return NewRegisterFactoryWithHook(nil)
}

func NewRegisterFactoryWithHook(hook registerHook) *RegisterFactory {
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
			res := read()
			rf.hook(regname, READ, res)
			return res
		}

		rwrite = func(data uint8) {
			write(data)
			rf.hook(regname, WRITE, data)
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
