package cpu

import (
	"fmt"
	"reflect"
)

func (cpu2 CPU) compare(cpu CPU) error {
	msg := ""

	if cpu.C != cpu2.C {
		msg += fmt.Sprintf("Accumulator value not matching, expected: %v, received: %v\n", cpu.C, cpu2.C)
	}

	if cpu.DBR != cpu2.DBR {
		msg += fmt.Sprintf("Data bank register value not matching, expected: %v, received: %v\n", cpu.DBR, cpu2.DBR)
	}

	if cpu.D != cpu2.D {
		msg += fmt.Sprintf("Direct register value not matching, expected: %v, received: %v\n", cpu.D, cpu2.D)
	}

	if cpu.K != cpu2.K {
		msg += fmt.Sprintf("Bank register value not matching, expected: %v, received: %v\n", cpu.K, cpu2.K)
	}

	if cpu.PC != cpu2.PC {
		msg += fmt.Sprintf("Program counter value not matching, expected: %v, received: %v\n", cpu.PC, cpu2.PC)
	}

	if cpu.S != cpu2.S {
		msg += fmt.Sprintf("Stack pointer value not matching, expected: %v, received: %v\n", cpu.S, cpu2.S)
	}

	if cpu.X != cpu2.X {
		msg += fmt.Sprintf("X register value not matching, expected: %v, received: %v\n", cpu.X, cpu2.X)
	}

	if cpu.Y != cpu2.Y {
		msg += fmt.Sprintf("Y register value not matching, expected: %v, received: %v\n", cpu.Y, cpu2.Y)
	}

	if cpu.eFlag != cpu2.eFlag {
		msg += fmt.Sprintf("Emulation flag value not matching, expected: %v, received: %v\n", cpu.eFlag, cpu2.eFlag)
	}

	if cpu.nFlag != cpu2.nFlag {
		msg += fmt.Sprintf("Negative flag value not matching, expected: %v, received: %v\n", cpu.nFlag, cpu2.nFlag)
	}

	if cpu.vFlag != cpu2.vFlag {
		msg += fmt.Sprintf("Overflow flag value not matching, expected: %v, received: %v\n", cpu.vFlag, cpu2.vFlag)
	}

	if cpu.mFlag != cpu2.mFlag {
		msg += fmt.Sprintf("Accumulator/Memory width flag value not matching, expected: %v, received: %v\n", cpu.mFlag, cpu2.mFlag)
	}

	if cpu.bFlag != cpu2.bFlag {
		msg += fmt.Sprintf("Break flag value not matching, expected: %v, received: %v\n", cpu.bFlag, cpu2.bFlag)
	}

	if cpu.xFlag != cpu2.xFlag {
		msg += fmt.Sprintf("Index register flag value not matching, expected: %v, received: %v\n", cpu.xFlag, cpu2.xFlag)
	}

	if cpu.dFlag != cpu2.dFlag {
		msg += fmt.Sprintf("Decimal mode flag value not matching, expected: %v, received: %v\n", cpu.dFlag, cpu2.dFlag)
	}

	if cpu.iFlag != cpu2.iFlag {
		msg += fmt.Sprintf("Interrupt disable flag value not matching, expected: %v, received: %v\n", cpu.iFlag, cpu2.iFlag)
	}

	if cpu.zFlag != cpu2.zFlag {
		msg += fmt.Sprintf("Zero flag value not matching, expected: %v, received: %v\n", cpu.zFlag, cpu2.zFlag)
	}

	if cpu.pFlag != cpu2.pFlag {
		msg += fmt.Sprintf("Page boundary crossed virtual flag value not matching, expected: %v, received: %v\n", cpu.pFlag, cpu2.pFlag)
	}

	if !reflect.DeepEqual(cpu.memory, cpu2.memory) {
		msg += fmt.Sprintf("Memories are not matching")
	}

	if msg != "" {
		return fmt.Errorf(msg)
	}

	return nil
}
