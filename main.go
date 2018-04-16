package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"

	"github.com/snes-emu/gose/core"
)

var filename string
var debug bool

func Flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.BoolVar(&debug, "debug", false, "Enable debug output and pprof server on localhost:8080/debug/pprof")
	flag.Parse()
}

func main() {
	Flags()

	emu := core.New()
	emu.ReadROM(filename)
	if debug {
		go debugServer()
		emu.CPU.StartDebug()
	} else {
		emu.CPU.Start()
	}
}

func debugServer() {
	http.ListenAndServe("localhost:8080", nil)
}
