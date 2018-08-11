package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/snes-emu/gose/core"
)

var VERSION string

func main() {
	fmt.Printf("Staring gose, version: %s\n", VERSION)

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a rom file to open")
		os.Exit(1)
	}

	emu := core.New()
	emu.ReadROM(flag.Arg(0))
	emu.CPU.Start()
}
