package render

import (
	"fmt"
	"github.com/snes-emu/gose/bit"
	"github.com/snes-emu/gose/log"
	"github.com/veandco/go-sdl2/sdl"
	"go.uber.org/zap"
)

var _ Renderer = &SDLRenderer{}

type SDLRenderer struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
}

func NewSDLRenderer(width, height int32) (*SDLRenderer, error) {
	sr := &SDLRenderer{}
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, fmt.Errorf("failed to init SDL renderer: %w", err)
	}

	var err error
	sr.window, sr.renderer, err = sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, fmt.Errorf("failed to create SDL window and renderer: %w", err)
	}

	sr.texture, err = sr.renderer.CreateTexture(sdl.PIXELFORMAT_BGR555, sdl.TEXTUREACCESS_STREAMING, width, height)
	if err != nil {
		return nil, fmt.Errorf("failed to create SDL texture: %w", err)
	}

	sr.window.SetTitle("Gose - SNES Emulator")
	// TODO: icon
	// TODO: poll close

	return sr, nil
}

func (sr *SDLRenderer) Render(screen *Screen) {
	raw, _, err := sr.texture.Lock(nil)
	if err != nil {
		// TODO: better error handling here ?
		log.Error("error locking texture", zap.Error(err))
		return
	}

	for i, p := range screen.Pixels {
		raw[2*i] = bit.LowByte(p.Color.Color)
		raw[2*i+1] = bit.HighByte(p.Color.Color)
	}

	sr.texture.Unlock()
	err = sr.renderer.Copy(sr.texture, nil, nil)
	if err != nil {
		log.Error("error copying texture", zap.Error(err))
	}
	sr.renderer.Present()
}

func (sr *SDLRenderer) Stop() {
	if sr.window != nil {
		sr.window.Destroy()
	}

	if sr.renderer != nil {
		sr.renderer.Destroy()
	}

	if sr.texture != nil {
		sr.texture.Destroy()
	}

	sdl.Quit()
}
