package render

import (
	"fmt"
	"image"
	"image/color"
)

type Screen struct {
	Pixels []Pixel
	Width  uint16
	Height uint16
	model  color.Model
}

func NewScreen(width, height uint16) *Screen {
	return &Screen{
		Pixels: make([]Pixel, width*height),
		Width:  width,
		Height: height,
		model:  color.ModelFunc(bgr555ModelFunc),
	}
}

func (s *Screen) SetPixelLine(line uint16, pixels []Pixel) {
	if line < s.Height {
		start := int(line * s.Width)
		for i, pix := range pixels[:s.Width] {
			s.Pixels[start+i] = pix
		}
	} else {
		panic(fmt.Sprintf("Screen not big enough ! can't set pixels at line %d in screen having only %d lines", line, s.Height))
	}
}

func (s *Screen) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(s.Width), int(s.Height))
}

func (s *Screen) At(x, y int) color.Color {
	return s.Pixels[y*int(s.Width)+x].Color
}

func (s *Screen) ColorModel() color.Model {
	return s.model
}
