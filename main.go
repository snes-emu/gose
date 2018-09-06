package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/snes-emu/gose/core"
	"github.com/veandco/go-sdl2/sdl"
)

var VERSION string

var texture *sdl.Texture
var renderer *sdl.Renderer
var window *sdl.Window

func main() {
	fmt.Printf("Staring gose, version: %s\n", VERSION)

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a rom file to open")
		os.Exit(1)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	var err error
	window, renderer, err = sdl.CreateWindowAndRenderer(core.HMax, core.VMaxNTSC, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_BGR555, sdl.TEXTUREACCESS_STREAMING, core.HMax, core.VMaxNTSC)
	if err != nil {
		panic(err)
	}
	emu := core.New(render)
	emu.ReadROM(flag.Arg(0))
	emu.CPU.Start()
}

func render(screen []byte) {
	pixels, _, _ := texture.Lock(nil)
	for i := 0; i < len(pixels); i++ {
		pixels[i] = screen[i]
	}
	texture.Unlock()
	renderer.Copy(texture, nil, nil)
	renderer.Present()

}
