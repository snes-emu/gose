package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/snes-emu/gose/config"
	"github.com/snes-emu/gose/core"
	"github.com/snes-emu/gose/debugger"
	"go.uber.org/zap"
)

// VERSION set at compile time
var VERSION string

func main() {
	// TODO allow to configure the logger / be in quiet mode
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to instatiate the logger, logging won't work")
	} else {
		zap.ReplaceGlobals(logger)
	}

	lg := zap.L()
	lg.Info("starting gose", zap.String("version", VERSION))

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a rom file to open")
		os.Exit(1)
	}

	emu := core.New()
	emu.ReadROM(flag.Arg(0))
	emu.Start()

	if config.DebugServer() {
		lg.Info("starting the debugger")
		db := debugger.New(emu, fmt.Sprintf("localhost:%d", config.DebugPort()))
		db.Start()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	emu.Stop()
	lg.Info("emulation stopped")
}
