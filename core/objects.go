package core

import "fmt"

//baseTile represents an 8x8 pixels tile
type baseTile struct {
	addr uint16 // address of the 8x8 tile in the VRAM

	// index of the color palette to use.
	// for background tiles, the number of entries in the palette depends on the mode and the background)
	// there are 16 available color palettes, but only 8 available to sprites
	palette    uint8
	colorDepth uint8 // number of bits used to addres the colors
	mode7      bool  // whether this tile is part of a mode7 background
}

//baseTileSize returns the size of a base tile in bytes depending on its color depth
func baseTileSize(colorDepth uint8) uint16 {
	//base tile size in bits = number of pixels * number of bits per pixel = 64 * colorDepth
	//base tile size in bytes = base tile size in bits / 8 = 8 * colorDepth
	return 8 * uint16(colorDepth)
}

// bgTile represents a variable size background tile
type bgTile struct {
	baseTile

	priority     bool   // Tile priority, for tiles this is only encoded in 1 bit (whereas for sprites it's encoded into 2 bits)
	hFlip, vFlip bool   // horizontal and vertical flips
	hSize, vSize uint16 // horizontal and vertical size
}

func (bgt *bgTile) tileAt(xTile, yTile uint16) baseTile {
	return baseTile{
		addr:       bgt.addr + (xTile+(yTile<<4))*baseTileSize(bgt.colorDepth),
		colorDepth: bgt.colorDepth,
		palette:    bgt.palette,
		mode7:      bgt.mode7,
	}
}

// sprite defines how a sprite is handled by the super-nes
// A sprite is composed by 8x8 tiles and can have the following sizes:
// 8x8, 16x16, 32x32, 64x64
// 16x32, 32x64
// We only store the address of the first tile composing the sprite, and we
// use the size of the sprite to determine the other tiles to use, with the following rules:
// - address of a tile to the right -> current address + 0x1
// - address of a tile to the bottom -> current address + 0x10
type sprite struct {
	baseTile

	x uint16 // x coordinate of the upper left tile composing the sprite
	y uint16 // y coordinate of the upper left tile composing the sprite

	priority uint8 // priority of the sprite (used to superpose multiple sprites / backgrounds)

	hFlip, vFlip bool   // horizontal and vertical flips
	hSize, vSize uint16 // horizontal and vertical sizes
}

// return true if the given sprite intersects the v-line
func (s *sprite) IntersectsLine(vCounter uint16) bool {
	return vCounter >= s.y && vCounter < s.y+s.vSize
}

// tileAt returns the tileAt at the given coordinate in the sprite
func (s *sprite) tileAt(xTile uint16, yTile uint16) baseTile {
	return baseTile{
		addr:       s.addr + (yTile<<4+xTile)*baseTileSize(4),
		colorDepth: s.colorDepth,
		palette:    s.palette,
	}
}

// spriteSize returns the horizontal and vertical sizes for a sprite given the oam config, it uses the following table:
// Val Small  Large
// 0 = 8x8    16x16
// 1 = 8x8    32x32
// 2 = 8x8    64x64
// 3 = 16x16  32x32
// 4 = 16x16  64x64
// 5 = 32x32  64x64
// 6 = 16x32  32x64 (undocumented)
// 7 = 16x32  32x32 (undocumented)
func spriteSize(isLarge bool, objectSize uint8) (uint16, uint16) {
	// Large
	if isLarge {
		switch objectSize {
		case 0:
			return 16, 16
		case 1, 3, 7:
			return 32, 32
		case 2, 4, 5:
			return 64, 64
		case 6:
			return 32, 64
		}
	}

	// Small
	switch objectSize {
	case 0, 1, 2:
		return 8, 8
	case 3, 4:
		return 16, 16
	case 5:
		return 32, 32
	case 6, 7:
		return 16, 32
	}

	panic(fmt.Sprintf("Invalid object size: %d", objectSize))
}
