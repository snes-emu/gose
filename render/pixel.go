package render

import (
	"encoding/json"
	"image/color"
)

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
// If the brightness is 0 we return a black color
// otherwise we replace the colors by color * brightness + 1 / 16
func (c BGR555) ApplyBrightness(bness uint8) BGR555 {
	if bness == 0 {
		return BGR555{
			Color:       0,
			Transparent: c.Transparent,
		}
	}

	r := int(c.Color) & 0x1f
	g := int(c.Color>>5) & 0x1f
	b := int(c.Color>>10) & 0x1f

	b32 := int(bness)

	r = (r * (1 + b32)) >> 4
	g = (g * (1 + b32)) >> 4
	b = (b * (1 + b32)) >> 4

	return BGR555{
		Color:       uint16(b<<10 | g<<5 | r),
		Transparent: c.Transparent,
	}
}

// RGBA implements the color.Color interface
func (c BGR555) RGBA() (r, g, b, a uint32) {
	a = 0xFFFF
	if c.Transparent {
		a = 0
	}
	r = uint32(c.Color&0x1F) << 11
	g = uint32((c.Color&0x3E0)>>5) << 11
	b = uint32((c.Color&0x7C00)>>10) << 11

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

func bgr555ModelFunc(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	return BGR555{
		Color:       uint16((b>>11)<<10 | (g>>11)<<5 | (r >> 11)),
		Transparent: false,
	}
}
