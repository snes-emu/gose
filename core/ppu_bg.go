package core

import (
	"fmt"

	"github.com/snes-emu/gose/bit"
)

type backgroundData struct {
	bg          [4]*bg // BG array containing the 4 backgrounds
	scrollPrev1 uint8  // temporary variable for bg scrolling
	scrollPrev2 uint8  // temporary variable for bg scrolling
	screenMode  uint8  // Screen mode from 0 to 7
	mosaicSize  uint8  // Size of block in mosaic mode (0=Smallest/1x1, 0xF=Largest/16x16)
}

// BG stores data about a background
type bg struct {
	tileSizeFlag     bool   // false 8x8 tiles, true 16x16 tiles
	mosaic           bool   // mosaic mode enabled
	priority         bool   // Only useful for BG3
	screenSize       uint8  // 0=32x32, 1=64x32, 2=32x64, 3=64x64 tiles
	tileMapBaseAddr  uint8  // base address for tile map in VRAM (in 1k word steps, 2k byte steps)
	tileSetBaseAddr  uint8  // base address for tile set in VRAM (in 4k word steps, 8k byte steps)
	horizontalScroll uint16 // horizontal scroll in pixel
	verticalScroll   uint16 // vertical scroll in pixel
	windowMask1      uint8  // mask for window 1 (0..1=Disable, 2=Inside, 3=Outside)
	windowMask2      uint8  // mask for window 2 (0..1=Disable, 2=Inside, 3=Outside)
	windowMaskLogic  uint8  // 0=OR, 1=AND, 2=XOR, 3=XNOR)
	mainScreenWindow bool   // Disable window area on main screen
	subScreenWindow  bool   // Disable windows area on sub screen
	mainScreen       bool   // Enable layer on main screen
	subScreen        bool   // Enable layer on sub screen
	colorMath        bool   // Flag to control colors on the BG (False: Display RAW Main Screen as such (without math), True: Apply math on Mainscreen)
}

// 2105h - BGMODE - BG Mode and BG Character Size (W)
func (ppu *PPU) bgmode(data uint8) {
	ppu.backgroundData.screenMode = data & 7
	ppu.backgroundData.bg[2].priority = data&8 != 0
	for i := uint8(0); i < 4; i++ {
		ppu.backgroundData.bg[i].tileSizeFlag = data&(1<<(4+i)) != 0
	}

}

// 2106h - MOSAIC - Mosaic Size and Mosaic Enable (W)
func (ppu *PPU) mosaic(data uint8) {
	for i := uint8(0); i < 4; i++ {
		ppu.backgroundData.bg[i].mosaic = data&(1<<i) != 0
	}
	ppu.backgroundData.mosaicSize = data >> 4
}

// 2107h -  210Ah - BG?SC - BG? Screen Base and Screen Size (W)
// 7-2  SC Base Address in VRAM (in 1K-word steps, aka 2K-byte steps)
// 1-0  SC Size (0=One-Screen, 1=V-Mirror, 2=H-Mirror, 3=Four-Screen)
// (0=32x32, 1=64x32, 2=32x64, 3=64x64 tiles)
// (0: SC0 SC0    1: SC0 SC1  2: SC0 SC0  3: SC0 SC1   )
// (   SC0 SC0       SC0 SC1     SC1 SC1     SC2 SC3   )
func (bg *bg) bgsc(data uint8) {
	bg.screenSize = data & 3
	bg.tileMapBaseAddr = data >> 2
}

// 2107h - BG1SC - BG1 Screen Base and Screen Size (W)
func (ppu *PPU) bg1sc(data uint8) {
	ppu.backgroundData.bg[0].bgsc(data)
}

// 2108h - BG2SC - BG2 Screen Base and Screen Size (W)
func (ppu *PPU) bg2sc(data uint8) {
	ppu.backgroundData.bg[1].bgsc(data)
}

// 2109h - BG3SC - BG3 Screen Base and Screen Size (W)
func (ppu *PPU) bg3sc(data uint8) {
	ppu.backgroundData.bg[2].bgsc(data)
}

// 210Ah - BG4SC - BG4 Screen Base and Screen Size (W)
func (ppu *PPU) bg4sc(data uint8) {
	ppu.backgroundData.bg[3].bgsc(data)
}

// 210Bh/210Ch - BG12NBA/BG34NBA - BG Character Data Area Designation (W)
func (ppu *PPU) bg12nba(data uint8) {
	// TODO: use util there
	ppu.backgroundData.bg[0].tileSetBaseAddr = data & 0x0F
	ppu.backgroundData.bg[1].tileSetBaseAddr = data >> 4
}

func (ppu *PPU) bg34nba(data uint8) {
	// TODO: use util there
	ppu.backgroundData.bg[2].tileSetBaseAddr = data & 0x0F
	ppu.backgroundData.bg[3].tileSetBaseAddr = data >> 4
}

// 210Dh - BG1HOFS - BG1 Horizontal Scroll (X) (W)
func (ppu *PPU) bg1hofs(data uint8) {
	ppu.backgroundData.bg[0].horizontalScroll = bit.JoinUint16(0x00, data) | uint16((ppu.backgroundData.scrollPrev1 &^ 7)) | uint16(ppu.backgroundData.scrollPrev2&7)
	ppu.backgroundData.scrollPrev1 = data
	ppu.backgroundData.scrollPrev2 = data
	ppu.m7hofs(data)
}

// 210Eh - BG1VOFS - BG1 Vertical Scroll (Y) (W)
func (ppu *PPU) bg1vofs(data uint8) {
	ppu.backgroundData.bg[0].horizontalScroll = bit.JoinUint16(0x00, data) | uint16(ppu.backgroundData.scrollPrev1)
	ppu.backgroundData.scrollPrev1 = data
	ppu.m7vofs(data)
}

// 210Fh - BG2HOFS - BG2 Horizontal Scroll (X) (W)
func (ppu *PPU) bg2hofs(data uint8) {
	ppu.backgroundData.bg[1].horizontalScroll = bit.JoinUint16(0x00, data) | uint16((ppu.backgroundData.scrollPrev1 &^ 7)) | uint16(ppu.backgroundData.scrollPrev2&7)
	ppu.backgroundData.scrollPrev1 = data
	ppu.backgroundData.scrollPrev2 = data
}

// 2110h - BG2VOFS - BG2 Vertical Scroll (Y) (W)
func (ppu *PPU) bg2vofs(data uint8) {
	ppu.backgroundData.bg[1].horizontalScroll = bit.JoinUint16(0x00, data) | uint16(ppu.backgroundData.scrollPrev1)
	ppu.backgroundData.scrollPrev1 = data
}

// 2111h - BG3HOFS - BG3 Horizontal Scroll (X) (W)
func (ppu *PPU) bg3hofs(data uint8) {
	ppu.backgroundData.bg[2].horizontalScroll = bit.JoinUint16(0x00, data) | uint16((ppu.backgroundData.scrollPrev1 &^ 7)) | uint16(ppu.backgroundData.scrollPrev2&7)
	ppu.backgroundData.scrollPrev1 = data
	ppu.backgroundData.scrollPrev2 = data
}

// 2112h - BG3VOFS - BG3 Vertical Scroll (Y) (W)
func (ppu *PPU) bg3vofs(data uint8) {
	ppu.backgroundData.bg[2].horizontalScroll = bit.JoinUint16(0x00, data) | uint16(ppu.backgroundData.scrollPrev1)
	ppu.backgroundData.scrollPrev1 = data
}

// 2113h - BG4HOFS - BG4 Horizontal Scroll (X) (W)
func (ppu *PPU) bg4hofs(data uint8) {
	ppu.backgroundData.bg[3].horizontalScroll = bit.JoinUint16(0x00, data) | uint16((ppu.backgroundData.scrollPrev1 &^ 7)) | uint16(ppu.backgroundData.scrollPrev2&7)
	ppu.backgroundData.scrollPrev1 = data
	ppu.backgroundData.scrollPrev2 = data
}

// 2114h - BG4VOFS - BG4 Vertical Scroll (Y) (W)
func (ppu *PPU) bg4vofs(data uint8) {
	ppu.backgroundData.bg[3].horizontalScroll = bit.JoinUint16(0x00, data) | uint16(ppu.backgroundData.scrollPrev1)
	ppu.backgroundData.scrollPrev1 = data
}

// tileMapAddress returns the byte address in the VRAM of the tile we are looking for in the tilemap
// See here: https://wiki.superfamicom.org/backgrounds
func (bg *bg) tileMapAddress(x uint16, y uint16) uint32 {
	// TODO: verify that, not sure at all about this
	var mapIndex uint32
	if bg.screenSize&0x1 != 0 {
		mapIndex += uint32(x / 32)
	}
	if bg.screenSize&0x2 != 0 {
		mapIndex += uint32(y / 32)
	}

	base := uint32(bg.tileMapBaseAddr)

	return (base+mapIndex)<<11 +
		((uint32(y) % 32) << 6) + //a row of 32 tile is 64 = 1<<6 bytes
		((uint32(x) % 32) << 1) //a tile is 2 = 1<<1 bytes
}

func (ppu *PPU) tileFromBackground(background uint8, x uint16, y uint16) tile {
	bg := ppu.backgroundData.bg[background]
	addr := bg.tileMapAddress(x, y)
	// raw contains:
	// vhopppcc cccccccc
	// v/h        = Vertical/Horizontal flip this tile.
	// 	o          = Tile priority.
	// 	ppp        = Tile palette. The number of entries in the palette depends on the Mode and the BG.
	// 	cccccccccc = Tile number.
	// See: https://wiki.superfamicom.org/backgrounds
	raw := bit.JoinUint16(ppu.vram.bytes[addr], ppu.vram.bytes[addr+1])

	hSize, vSize := bg.tileSize()

	return tile{
		vFlip:         raw&0x8000 != 0,
		hFlip:         raw&0x4000 != 0,
		priority:      raw&0x2000 != 0,
		palette:       uint8((raw >> 10) & 0x7),
		number:        raw & 0x3ff,
		firstTileAddr: uint32(bg.tileSetBaseAddr) << 13,
		colorDepth:    ppu.colorDepth(background),
		hSize:         hSize,
		vSize:         vSize,
	}
}

//tileSize returns the size in pixel of tiles in the background
func (bg *bg) tileSize() (uint16, uint16) {
	hSize, vSize := uint16(8), uint16(8)
	if bg.tileSizeFlag {
		hSize, vSize = 16, 16
	}

	return hSize, vSize
}

// colorDepth returns the number of bits used for the colors in the background
func (ppu *PPU) colorDepth(background uint8) uint8 {
	panicString := "in mode %d, only background 1,2,3 are valid, attempted to use background %d"
	switch ppu.backgroundData.screenMode {
	case 0:
		return 2
	case 1:
		switch background {
		case 0, 1:
			return 4
		case 2:
			return 2
		default:
			panic(fmt.Sprintf(panicString, ppu.backgroundData.screenMode, background+1))
		}
	case 2:
		switch background {
		case 0, 1:
			return 4
		default:
			panic(fmt.Sprintf(panicString, ppu.backgroundData.screenMode, background+1))
		}
	case 3, 4:
		switch background {
		case 0:
			return 8
		case 1:
			return 4
		default:
			panic(fmt.Sprintf(panicString, ppu.backgroundData.screenMode, background+1))
		}
	case 5:
		switch background {
		case 0:
			return 4
		case 1:
			return 2
		default:
			panic(fmt.Sprintf(panicString, ppu.backgroundData.screenMode, background+1))
		}
	case 6:
		switch background {
		case 0:
			return 4
		default:
			panic(fmt.Sprintf(panicString, ppu.backgroundData.screenMode, background+1))
		}
	case 7:
		switch background {
		case 0:
			return 8
		}
	default:
		panic(fmt.Sprintf("invalid mode requested: %d", ppu.backgroundData.screenMode))
	}

	//should never happen
	return 0
}

//validBackgrounds are the backgrounds that can be used for the current screen mode
func (ppu *PPU) validBackgrounds() []uint8 {
	bgs := []uint8{0}
	mode := ppu.backgroundData.screenMode
	if mode < 6 {
		bgs = append(bgs, 1)
	}
	if mode < 2 {
		bgs = append(bgs, 2)
	}
	if mode == 0 {
		bgs = append(bgs, 3)
	}

	return bgs
}
