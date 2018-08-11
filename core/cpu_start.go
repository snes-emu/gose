// +build !debug

package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (cpu *CPU) Start() {
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
			cpu.opcodes[opcode]()
		}
	}
}
