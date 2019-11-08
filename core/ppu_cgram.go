package core

import (
	"image/color"

	"github.com/snes-emu/gose/bit"
	"github.com/snes-emu/gose/render"
)

const cgramSize = 0x200

// cgram represents the color graphics ram and stores the color palette with 256 color entries
// Colors are stored in 2-byte words like:
// _bbbbbgg gggrrrrr
// The last bit is not used
type cgram struct {
	bytes [cgramSize]byte // bytes represents the raw 512 bytes (for 256 color entries)
	addr  uint16          // store the cgram address over 512 byte (not the Word addr !)
	// lsb temporary variable for the cgdata register, it's used when we want to write a new color:
	// - we first set the addr where we want to write (for instance the color number 67, we will have 61 * 2 = 122 stored in addr
	// - we call cgdata to write once (this will store the provided data in the lsb)
	// - we call cgdata one more time and then we will write (data << 8 | lsb) to the registry
	lsb uint8
}

// 2121 - Color index (0..255). This is a WORD-address (2-byte steps), allowing to access 256 words (512 bytes). Writing to this register resets the 1st/2nd access flipflop (for 2122h/213Bh) to 1st access.
func (ppu *PPU) cgadd(addr uint8) {
	ppu.cgram.addr = 2 * uint16(addr)
}

// 2122 - CGDATA - Palette CGRAM Data Write (W)
func (ppu *PPU) cgdata(data uint8) {
	if ppu.cgram.addr%2 == 0 {
		// Write to the temporary variable
		ppu.cgram.lsb = data
	} else {
		// addr - 1 because we increment even if we wrote in the lsb
		ppu.cgram.write(ppu.cgram.addr-1, ppu.cgram.lsb, data)
	}
	ppu.cgram.incrAddr()
}

// 213B - RDCGRAM - Palette CGRAM Data Read (R)
func (ppu *PPU) rdcgram() uint8 {
	res := ppu.cgram.read(ppu.cgram.addr)
	ppu.cgram.incrAddr()
	return res
}

func (cg *cgram) read(addr uint16) uint8 {
	return cg.bytes[addr]
}

func (cg *cgram) write(addr uint16, low uint8, high uint8) {
	// This could be (high & 0x7f) but since last bit is never used in color palette it's not an issue
	cg.bytes[addr+1] = high
	cg.bytes[addr] = low
}

func (cg *cgram) incrAddr() {
	cg.addr = (cg.addr + 1) % 512
}

// Palette exports the content of the cgram
func (ppu *PPU) Palette() color.Palette {
	var palette color.Palette

	for i := 0; i < len(ppu.cgram.bytes)/2; i++ {
		palette = append(palette, render.BGR555{Color: bit.JoinUint16(ppu.cgram.bytes[2*i], ppu.cgram.bytes[2*i+1])})
	}
	return palette
}

// colorIndex returns the color index for a tile's pixel inside the sprite palette
// tileAddress is the tile address in the vram
// colorDepth is the color depth of the current tile (should be 2 for 4 colors, 4 for 16 colors or 8 for 256 colors)
// x is the x position of the pixel inside the tile (from 0 to 7 included)
// y is the y position of the pixel inside the tile (from 0 to 7 included)
//
// Each 8x8 tile occupies 16, 32, or 64 bytes (for 4, 16, or 256 colors). BG tiles can be 4/16/256 colors (depending on BG Mode), OBJs are always 16 color.

//   Color Bits (Planes)     Upper Row ........... Lower Row
//   Plane 0 stored in bytes 00h,02h,04h,06h,08h,0Ah,0Ch,0Eh ;\for 4/16/256 colors
//   Plane 1 stored in bytes 01h,03h,05h,07h,09h,0Bh,0Dh,0Fh ;/
//   Plane 2 stored in bytes 10h,12h,14h,16h,18h,1Ah,1Ch,1Eh ;\for 16/256 colors
//   Plane 3 stored in bytes 11h,13h,15h,17h,19h,1Bh,1Dh,1Fh ;/
//   Plane 4 stored in bytes 20h,22h,24h,26h,28h,2Ah,2Ch,2Eh ;\
//   Plane 5 stored in bytes 21h,23h,25h,27h,29h,2Bh,2Dh,2Fh ; for 256 colors
//   Plane 6 stored in bytes 30h,32h,34h,36h,38h,3Ah,3Ch,3Eh ;
//   Plane 7 stored in bytes 31h,33h,35h,37h,39h,3Bh,3Dh,3Fh ;/
//   In each byte, bit7 is left-most, bit0 is right-most.
//   Plane 0 is the LSB of color number.
func (ppu *PPU) colorIndex(tileAddress, colorDepth, x, y uint16) uint8 {
	var colorIndex uint8

	base := tileAddress + 2*y
	shift := (7 - x)
	// 4 colors
	colorIndex += (ppu.vram.bytes[base+0x00] >> shift & 1) << 0
	colorIndex += (ppu.vram.bytes[base+0x01] >> shift & 1) << 1

	// 16 colors
	if colorDepth >= 4 {
		colorIndex += (ppu.vram.bytes[base+0x10] >> shift & 1) << 2
		colorIndex += (ppu.vram.bytes[base+0x11] >> shift & 1) << 3
	}

	// 256 colors
	if colorDepth >= 8 {
		colorIndex += (ppu.vram.bytes[base+0x20] >> shift & 1) << 4
		colorIndex += (ppu.vram.bytes[base+0x21] >> shift & 1) << 5
		colorIndex += (ppu.vram.bytes[base+0x30] >> shift & 1) << 6
		colorIndex += (ppu.vram.bytes[base+0x31] >> shift & 1) << 7
	}
	return colorIndex
}

// tileSpriteRowColors returns the colors to use for a given tile's row
// tileAddress is the tile address in the vram
// y is the row number inside the tile (from 0 to 7 included)
// palette is the palette we should use
func (ppu *PPU) tileSpriteRowColor(tileAddress, y uint16, palette uint8) [TILE_SIZE]render.BGR555 {
	colors := [TILE_SIZE]render.BGR555{}

	for x := uint16(0); x < TILE_SIZE; x++ {
		colors[x] = ppu.tileSpriteColor(tileAddress, x, y, palette)
	}

	return colors
}

// tileRowColors returns the colors to use for a given tile's row
// tileAddress is the tile address in the vram
// hSize is the horizontal size of the tile: either 8 or 16
// y is the row number inside the tile (from 0 to 7 included)
// palette is the palette we should use
func (ppu *PPU) tileRowColor(tileAddress, colorDepth, hSize, y uint16, palette uint8) []render.BGR555 {
	colors := make([]render.BGR555, hSize)

	for x := uint16(0); x < hSize; x++ {
		colors[x] = ppu.tileColor(tileAddress, colorDepth, x, y, palette)
	}

	return colors
}

// tileSpriteColor is like tileColor but for sprites
func (ppu *PPU) tileSpriteColor(tileAddress, x, y uint16, palette uint8) render.BGR555 {
	// For sprites the color depth is always 4
	return ppu.tileColor(tileAddress, 4, x, y, palette)
}

// tileColor returns the color to use for a given tile's pixel in the given sprite
// tileAddress is the tile address in the vram
// colorDepth is the number of bits used per color
// x is the x position of the pixel inside the tile (from 0 to 7 included)
// y is the y position of the pixel inside the tile (from 0 to 7 included)
// palette is the palette we should use
func (ppu *PPU) tileColor(tileAddress, colorDepth, x, y uint16, palette uint8) render.BGR555 {
	idx := ppu.colorIndex(tileAddress, colorDepth, x, y)

	if idx == 0 {
		return render.BGR555{Transparent: true}
	}

	// Sprite colors are stored in the CGRAM starting at palette 8
	colorWordAddr := 2 * uint16(palette+idx)
	return render.BGR555{
		Color: bit.JoinUint16(ppu.cgram.bytes[colorWordAddr], ppu.cgram.bytes[colorWordAddr+1]),
	}.ApplyBrightness(ppu.display.brightness)
}

func (ppu *PPU) backdropPixel() render.Pixel {
	return render.Pixel{
		Visible: true,
		Color: render.BGR555{
			Color: bit.JoinUint16(ppu.cgram.bytes[0], ppu.cgram.bytes[1]),
		},
	}
}
