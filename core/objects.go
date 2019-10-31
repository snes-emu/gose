package core

import "fmt"

// tile defines how a tile is handled by the super-nes
type tile struct {
	number       uint16 // tile number (used to find it in the vram)
	palette      uint8  // index of the color palette to use (the number of entries in the palette depends on the mode and the background)
	priority     bool   // Tile priority, for tiles this is only encoded in 1 bit (whereas for sprites it's encoded into 2 bits)
	hFlip, vFlip bool   // horizontal and vertical flips
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
	x uint16 // x coordinate of the upper left tile composing the sprite
	y uint16 // y coordinate of the upper left tile composing the sprite

	firstTileAddr uint16 // address of the first tile composing the sprite in the VRAM

	palette  uint8 // index of the color palette to use (there are 16 available color palettes, but only 8 available to sprites)
	priority uint8 // priority of the sprite (used to superpose multiple sprites / backgrounds)

	hFlip, vFlip bool   // horizontal and vertical flips
	hSize, vSize uint16 // horizontal and vertical sizes
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
