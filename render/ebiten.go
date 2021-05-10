// +build !ci

package render

import (
	"fmt"
	"image"
	"image/png"

	"github.com/snes-emu/gose/log"
	"go.uber.org/zap"

	"github.com/gobuffalo/packr/v2"
	"github.com/hajimehoshi/ebiten/v2"
)

//EbitenRenderer is a Renderer implementation using ebiten
type EbitenRenderer struct {
	width           int
	height          int
	scale           float64
	offscreenBuffer *ebiten.Image
	drawOptions     *ebiten.DrawImageOptions
	running         bool
}

//newEbitenRenderer creates a ebiten renderer
func newEbitenRenderer(width, height int) (Renderer, error) {
	//We use this offscreen buffer because we don't want our SNES main loop to be tied to the ebiten one
	//NewImage always returns a nil error
	offscreenBuffer := ebiten.NewImage(width, height)
	er := &EbitenRenderer{
		width:           width,
		height:          height,
		scale:           2.0,
		offscreenBuffer: offscreenBuffer,
		drawOptions:     &ebiten.DrawImageOptions{},
	}

	ebiten.SetWindowIcon(getWindowLogos())
	ebiten.SetWindowTitle("Gose")
	ebiten.SetWindowSize(er.width*int(er.scale), er.height*int(er.scale))
	ebiten.SetRunnableOnUnfocused(true)

	return er, nil
}

//Render updates the offscreen buffer with the new SNES screen content
func (er *EbitenRenderer) Render(screen *Screen) {
	//consecutive Set calls are efficient
	for x := 0; x < er.width; x++ {
		for y := 0; y < er.height; y++ {
			er.offscreenBuffer.Set(x, y, screen.At(x, y))
		}
	}
}

//SetRomTitle stores the new title if ebiten is not yet running or set the title directly
func (er *EbitenRenderer) SetRomTitle(title string) {
	ebiten.SetWindowTitle(fmt.Sprintf("Gose - %s", title))

}

func (er *EbitenRenderer) Update() error {
	return nil
}

//update copies the content of the offscreenBuffer to the screen
func (er *EbitenRenderer) Draw(screen *ebiten.Image) {
	screen.DrawImage(er.offscreenBuffer, er.drawOptions)
}

func (er *EbitenRenderer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return er.width, er.height
}

//Run starts the ebiten main loop
//should be called on the main thread
func (er *EbitenRenderer) Run() {
	err := ebiten.RunGame(er)
	if err != nil {
		log.Fatal("ebiten crashed", zap.Error(err))
	}
}

//Stop implements the Renderer interface
func (er *EbitenRenderer) Stop() {

}

func getWindowLogos() []image.Image {
	logoBox := packr.New("logos", "../logos")

	var logos []image.Image

	for _, filename := range logoBox.List() {
		logoFile, err := logoBox.Resolve(filename)
		if err != nil {
			continue
		}

		logo, err := png.Decode(logoFile)
		if err != nil {
			continue
		}

		logos = append(logos, logo)

	}

	return logos
}

func init() {
	register("ebiten", newEbitenRenderer)
}
