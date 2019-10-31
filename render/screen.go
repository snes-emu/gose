package render

import "fmt"

type Screen struct {
	Pixels []Pixel
	Width  uint16
	Height uint16
}

func NewScreen(width, height uint16) *Screen {
	return &Screen{
		Pixels: make([]Pixel, width*height),
		Width:  width,
		Height: height,
	}
}

func (s *Screen) SetPixelLine(line uint16, pixels []Pixel) {
	if line < s.Height {
		start := int(line * s.Height)
		for i, pix := range pixels[:s.Width] {
			s.Pixels[start+i] = pix
		}
	} else {
		panic(fmt.Sprintf("Screen not big enough ! can't set pixels at line %d in screen having only %d lines", line, s.Height))
	}
}
