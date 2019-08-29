package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/snes-emu/gose/config"
	"github.com/snes-emu/gose/core"
)

var VERSION string

func main() {
	fmt.Printf("Staring gose, version: %s\n", VERSION)

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a rom file to open")
		os.Exit(1)
	}

	emu := core.New()
	emu.ReadROM(flag.Arg(0))
	emu.Start()

	if config.DebugServer() {

		fmt.Println("start debugger")
		debugger.Launch()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	emu.Stop()
	fmt.Printf("Emulator exited")
}
