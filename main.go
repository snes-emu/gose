package main

import (
	"flag"
	"fmt"
	"github.com/snes-emu/gose/render"
	"os"
	"os/signal"
	"syscall"

	"github.com/snes-emu/gose/log"

	"github.com/snes-emu/gose/config"
	"github.com/snes-emu/gose/core"
	"github.com/snes-emu/gose/debugger"
	"go.uber.org/zap"
)

// VERSION set at compile time
var VERSION string

func main() {
	config.Init()

	log.Info("starting gose", zap.String("version", VERSION))

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a rom file to open")
		os.Exit(1)
	}

	// TODO: fix dimension
	renderer, err := render.NewSDLRenderer(core.WIDTH, core.HEIGHT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating renderer: %s", err)
		os.Exit(1)
	}
	emu := core.New(renderer)
	emu.ReadROM(flag.Arg(0))

	if config.DebugServer() {
		log.Info("starting the debugger")
		db := debugger.New(emu, fmt.Sprintf("localhost:%d", config.DebugPort()))
		db.Start()
		emu.StartPaused()
	} else {
		emu.Start()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	emu.Stop()
	log.Info("emulation stopped")
}
