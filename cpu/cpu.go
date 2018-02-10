package cpu

type CPU struct {
	name string
	P    int16
}

type cpuOperation func()

var opcodes []cpuOperation
