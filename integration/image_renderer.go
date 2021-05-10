package integration

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/snes-emu/gose/render"
)

var _ render.Renderer = &ImageRenderer{}

type ImageRenderer struct {
	width  int
	height int
	Image  *image.RGBA
}

func NewImageRenderer(width, height int) *ImageRenderer {
	image := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	return &ImageRenderer{Image: image, width: width, height: height}
}

func (i *ImageRenderer) Render(screen *render.Screen) {
	for x := 0; x < i.width; x++ {
		for y := 0; y < i.height; y++ {
			i.Image.Set(x, y, screen.At(x, y))
		}
	}
	fmt.Println("Rendered")
}

func (i *ImageRenderer) SaveToFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file at: %s: %w", filepath, err)
	}

	err = png.Encode(file, i.Image)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to encode screen as png: %w", err)
	}

	return file.Close()
}

func (i ImageRenderer) Stop() {
	// no-op
}

func (i ImageRenderer) SetRomTitle(string) {
	// no-op
}

func (i ImageRenderer) Run() {
	// no-op
}
