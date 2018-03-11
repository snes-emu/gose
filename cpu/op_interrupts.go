package cpu

import (
	"fmt"
	"os"
)

// stp STP stops the clock input of the 65C816,
func (cpu *CPU) stp() {
	fmt.Print("Cpu has been shutdown")
	os.Exit(0)
}

func (cpu *CPU) opDB() {
	cpu.stp()
}

// wai stops the clock input of the 65C816,
func (cpu *CPU) wai() {
	cpu.waiting = true
}

func (cpu *CPU) opCB() {
	cpu.wai()
}
