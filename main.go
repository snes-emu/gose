// +build !debug

package main

import (
	"flag"
	_ "net/http/pprof"

	"github.com/snes-emu/gose/core"
)

var filename string
var pprof bool

func flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.Parse()
}

func main() {
	flags()

	emu := core.New()
	emu.ReadROM(filename)
	emu.CPU.Start()
}
