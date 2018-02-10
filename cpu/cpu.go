package cpu

type CPU struct {
	name string
}

type cpuOperation func()

var opcodes []cpuOperation
