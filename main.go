package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/snes-emu/gose/core"
	"github.com/snes-emu/gose/ppu"
	"github.com/snes-emu/gose/rom"
	"github.com/veandco/go-sdl2/sdl"
)

var filename string

func Flags() {
	flag.StringVar(&filename, "filename", "", "filename of the ROM to load")
	flag.Parse()
}

func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()
	Flags()
	ROM, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	rom, err := rom.ParseROM(ROM)
	fmt.Println(rom.Title)

	if err != nil {
		log.Fatalf("There were a problem while importing the ROM: %v", err)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(ppu.HMax, ppu.VMaxNTSC, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_BGR555, sdl.TEXTUREACCESS_STREAMING, ppu.HMax, ppu.VMaxNTSC)
	if err != nil {
		panic(err)
	}

	Ppu := ppu.New()
	mem := core.NewMemory(Ppu)
	mem.LoadROM(*rom)
	cpu := core.NewCPU(mem)
	cpu.Init()
	for {
		cpu.Execute(1364)
		Ppu.RenderLine()
		if Ppu.VCounter == 0 {
			pixels, _, err := texture.Lock(nil)
			if err != nil {
				panic(err)
			}
			for i := 0; i < len(pixels); i++ {
				pixels[i] = ppu.Screen[i]
			}
			texture.Unlock()
			renderer.Copy(texture, nil, nil)
			renderer.Present()
			fmt.Println("frame !")
		}

	}
}
