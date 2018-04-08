package main

import (
	"flag"

	"github.com/snes-emu/gose/core"
)

var filename string

func Flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.Parse()
}

func main() {
	Flags()

	mem := core.NewMemory()
	mem.LoadROM(*rom)
}
