package render

import "encoding/json"

// Pixel represents a pixel
type Pixel struct {
	Color    BGR555
	Visible  bool
	Priority uint8
}

//BGR555 represents 16-bit opaque color, each channel uses 5 bits with red in the least significant bits
type BGR555 struct {
	Color       uint16
	Transparent bool
}

// ApplyBrightness apply the given brightness from 0 to 15
// TODO
func (c BGR555) ApplyBrightness(bness uint8) BGR555 {
	return c
}

// RGBA implements the color.Color interface
func (c BGR555) RGBA() (r, g, b, a uint32) {
	a = 0xFFFF
	if c.Transparent {
		a = 0
	}
	r = uint32(c.Color&0x1F) << 3
	g = uint32((c.Color&0x3E0)>>5) << 3
	b = uint32((c.Color&0x7C00)>>10) << 3

	return
}

//MarshalJSON implements the json.Marshaler interface
func (c BGR555) MarshalJSON() ([]byte, error) {
	r, g, b, a := c.RGBA()
	return json.Marshal(map[string]uint32{
		"r": r,
		"g": g,
		"b": b,
		"a": a,
	})
}
