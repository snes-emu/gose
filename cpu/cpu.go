package cpu

type CPU struct {
	name   string
	P      int16
	cycles int
}

type cpuOperation func()

var opcodes []cpuOperation
