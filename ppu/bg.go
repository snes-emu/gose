package ppu

import (
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
	tileSize           bool   // false 8x8 tiles, true 16x16 tiles
	mosaic             bool   // mosaic mode enabled
	priority           bool   // Only useful for BG3
	screenSize         uint8  // 0=32x32, 1=64x32, 2=32x64, 3=64x64 tiles
	tileMapBaseAddress uint16 // base address for tile map in VRAM
	tileSetBaseAddress uint16 // base address for tile set in VRAM
	horizontalScroll   uint16 // horizontal scroll in pixel
	verticalScroll     uint16 // vertical scroll in pixel
	windowMask1        uint8  // mask for window 1 (0..1=Disable, 2=Inside, 3=Outside)
	windowMask2        uint8  // mask for window 2 (0..1=Disable, 2=Inside, 3=Outside)
	windowMaskLogic    uint8  // 0=OR, 1=AND, 2=XOR, 3=XNOR)
	mainScreenWindow   bool   // Disable window area on main screen
	subScreenWindow    bool   // Disable windows area on sub screen
	mainScreen         bool   // Enable layer on main screen
	subScreen          bool   // Enable layer on sub screen
	colorMath          bool   // Flag to control colors on the BG (False: Display RAW Main Screen as such (without math), True: Apply math on Mainscreen)
}

// 2105h - BGMODE - BG Mode and BG Character Size (W)
func (ppu *PPU) bgmode(data uint8) {
	ppu.backgroundData.screenMode = data & 7
	ppu.backgroundData.bg[2].priority = data&8 != 0
	for i := uint8(0); i < 4; i++ {
		ppu.backgroundData.bg[i].tileSize = data&(1<<(4+i)) != 0
	}

}

// 2106h - MOSAIC - Mosaic Size and Mosaic Enable (W)
func (ppu *PPU) mosaic(data uint8) {
	for i := uint8(0); i < 4; i++ {
		ppu.backgroundData.bg[i].mosaic = data&(1<<i) != 0
	}
	ppu.backgroundData.mosaicSize = data >> 4
}

// 2107h - BG1SC - BG1 Screen Base and Screen Size (W)
func (ppu *PPU) bg1sc(data uint8) {
	ppu.backgroundData.bg[0].screenSize = data & 3
	ppu.backgroundData.bg[0].tileMapBaseAddress = bit.JoinUint16(0x00, data&^uint8(3))
}

// 2108h - BG2SC - BG2 Screen Base and Screen Size (W)
func (ppu *PPU) bg2sc(data uint8) {
	ppu.backgroundData.bg[1].screenSize = data & 3
	ppu.backgroundData.bg[1].tileMapBaseAddress = bit.JoinUint16(0x00, data&^uint8(3))
}

// 2109h - BG3SC - BG3 Screen Base and Screen Size (W)
func (ppu *PPU) bg3sc(data uint8) {
	ppu.backgroundData.bg[2].screenSize = data & 3
	ppu.backgroundData.bg[2].tileMapBaseAddress = bit.JoinUint16(0x00, data&^uint8(3))
}

// 210Ah - BG4SC - BG4 Screen Base and Screen Size (W)
func (ppu *PPU) bg4sc(data uint8) {
	ppu.backgroundData.bg[3].screenSize = data & 3
	ppu.backgroundData.bg[3].tileMapBaseAddress = bit.JoinUint16(0x00, data&^uint8(3))
}

// 210Bh/210Ch - BG12NBA/BG34NBA - BG Character Data Area Designation (W)
func (ppu *PPU) bg12nba(data uint8) {
	// TODO: use util there
	ppu.backgroundData.bg[0].tileSetBaseAddress = uint16(data&0x0F) << 12
	ppu.backgroundData.bg[1].tileSetBaseAddress = bit.JoinUint16(0x00, data&0xF0)
}

func (ppu *PPU) bg34nba(data uint8) {
	// TODO: use util there
	ppu.backgroundData.bg[2].tileSetBaseAddress = uint16(data&0x0F) << 12
	ppu.backgroundData.bg[3].tileSetBaseAddress = bit.JoinUint16(0x00, data&0xF0)
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
