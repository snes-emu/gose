// +build debug

package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (cpu *CPU) StartDebug() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigs:
			fmt.Printf("Emulator exited")
			os.Exit(0)
		default:
			K := cpu.getKRegister()
			PC := cpu.getPCRegister()
			opcode := cpu.memory.GetByteBank(K, PC)
			cpu.logState()
			cpu.opcodes[opcode]()
		}
	}
}

func (cpu *CPU) logState() {
	fmt.Printf("%02X/%04X:  %02X opcode     A: %04X X: %04X Y: %04X D: %04X DB: %02X S: %04X\n", cpu.getKRegister(), cpu.getPCRegister(), cpu.memory.GetByteBank(cpu.getKRegister(), cpu.getPCRegister()), cpu.getCRegister(), cpu.getXRegister(), cpu.getYRegister(), cpu.getDRegister(), cpu.getDBRRegister(), cpu.getSRegister())
}
