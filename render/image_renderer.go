package render

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/snes-emu/gose/log"
	"go.uber.org/zap"
)

var _ Renderer = &ImageRenderer{}

type ImageRenderer struct {
	width  int
	height int
	path   string
	Image  *image.RGBA
}

func NewImageRenderer(width, height int, path string) *ImageRenderer {
	var image = image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{X: width, Y: height}})
	return &ImageRenderer{Image: image, width: width, height: height, path: path}
}

func (i *ImageRenderer) Render(screen *Screen) {
	for x := 0; x < i.width; x++ {
		for y := 0; y < i.height; y++ {
			i.Image.Set(x, y, screen.At(x, y))
		}
	}
}

func (i *ImageRenderer) saveToFile() error {
	file, err := os.Create(i.path)
	if err != nil {
		return fmt.Errorf("failed to create file at: %s: %w", i.path, err)
	}

	err = png.Encode(file, i.Image)
	if err != nil {
		file.Close()
		return fmt.Errorf("failed to encode screen as png: %w", err)
	}

	return file.Close()
}

func (i ImageRenderer) Stop() {
	err := i.saveToFile()
	if err != nil {
		log.Error("Failed to save image file", zap.Error(err))
	} else {
		log.Info("Saved screen at path", zap.String("path", i.path))
	}
}

func (i ImageRenderer) SetRomTitle(string) {
	// no-op
}

func (i ImageRenderer) Run() {
	// no-op
}
