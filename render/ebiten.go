package render

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

//EbitenRenderer is a Renderer implementation using ebiten
type EbitenRenderer struct {
	width           int
	height          int
	scale           float64
	title           string
	offscreenBuffer *ebiten.Image
	drawOptions     *ebiten.DrawImageOptions
	running         bool
}

//NewEbitenRenderer creates a ebiten renderer
func NewEbitenRenderer(width, height int) *EbitenRenderer {
	//We use this offscreen buffer because we don't want our SNES main loop to be tied to the ebiten one
	offscreenBuffer, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	er := &EbitenRenderer{
		width:           width,
		height:          height,
		scale:           2.0,
		title:           "Gose",
		offscreenBuffer: offscreenBuffer,
		drawOptions:     &ebiten.DrawImageOptions{},
	}

	return er
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
	title = fmt.Sprintf("Gose - %s", title)
	if er.running {
		ebiten.SetWindowTitle(title)
	} else {
		er.title = title
	}
}

//update copies the content of the offscreenBuffer to the screen
func (er *EbitenRenderer) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	return screen.DrawImage(er.offscreenBuffer, er.drawOptions)
}

//Run starts the ebiten main loop
//should be called on the main thread
func (er *EbitenRenderer) Run() error {
	return ebiten.Run(er.update, er.width, er.height, er.scale, er.title)
}

//Stop implements the Renderer interface
func (er *EbitenRenderer) Stop() {

}
