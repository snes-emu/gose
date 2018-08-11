package main

import (
	"flag"

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
	emu.CPU.Start()
}
