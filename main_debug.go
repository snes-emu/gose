// +build debug

package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"

	"github.com/snes-emu/gose/core"
)

var filename string

func flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.Parse()
}

func main() {
	flags()

	emu := core.New()
	emu.ReadROM(filename)
	go debugServer()
	emu.CPU.StartDebug()
}

func debugServer() {
	http.ListenAndServe("localhost:8080", nil)
}
