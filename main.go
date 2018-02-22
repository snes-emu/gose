package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/snes-emu/gose/memory"
	"github.com/snes-emu/gose/rom"
)

var filename string

func Flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.Parse()
}

func main() {
	Flags()
	ROM, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	rom, err := rom.ParseROM(ROM)

	if err != nil {
		log.Fatalf("There were a problem while importing the ROM: %v", err)
	}

	mem := memory.New()
	mem.LoadROM(*rom)
}
