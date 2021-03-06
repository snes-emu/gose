package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/snes-emu/gose/render"

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
	log.Init()

	log.Info("starting gose", zap.String("version", VERSION))

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a rom file to open")
		os.Exit(1)
	}

	if config.ImagePath() != "" {
		renderer := render.NewImageRenderer(core.WIDTH, core.HEIGHT, config.ImagePath())
		emu := core.New(renderer, true)
		emu.ReadROM(flag.Arg(0))
		emu.Start()
		emu.StepAndWait(config.NCycles())
		emu.Stop()
		renderer.Stop() // TODO: move into emu.Stop()
	} else {
		renderer, err := render.NewRenderer(int(core.WIDTH), int(core.HEIGHT))
		if err != nil {
			log.Fatal("failed to init renderer", zap.Error(err))
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

		go func() {
			<-sigs
			emu.Stop()
			log.Info("emulation stopped")
			os.Exit(0)
		}()

		//should be run on the main thread
		renderer.Run()
	}
}
