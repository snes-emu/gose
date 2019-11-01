package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/snes-emu/gose/render"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/snes-emu/gose/log"

	"github.com/snes-emu/gose/config"
	"github.com/snes-emu/gose/core"
	"github.com/snes-emu/gose/debugger"
	"go.uber.org/zap"
)

// VERSION set at compile time
var VERSION string

func main() {
	var exitcode int
	sdl.Main(func() {
		exitcode = Main()
	})
	os.Exit(exitcode)
}

func Main() int {
	config.Init()

	log.Info("starting gose", zap.String("version", VERSION))

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a rom file to open")
		return 1
	}

	// TODO: fix dimension
	renderer, err := render.NewSDLRenderer(core.WIDTH, core.HEIGHT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating renderer: %s", err)
		return 1
	}
	emu := core.New(renderer, config.DebugServer())
	emu.ReadROM(flag.Arg(0))

	if config.DebugServer() {
		log.Info("starting the debugger")
		db := debugger.New(emu, fmt.Sprintf("localhost:%d", config.DebugPort()))
		db.Start()
	}

	emu.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	emu.Stop()
	log.Info("emulation stopped")
	return 0
}
