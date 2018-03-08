package ppu

const (
	HMax     = 339 // max H counter value
	VMaxNTSC = 261 // max V counter value in NTSC
	VMaxPAL  = 311 // max V counter value in PAL
)

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram           *vram           // vram represents the VideoRAM (64KB)
	oam            *oam            // oam represents the object attribute memory (512 + 32 Bytes)
	cgram          *cgram          // cgram represents the color graphics ram and stores the color palette with 256 color entries
	backgroundData *backgroundData // background data
	colorMath      *colorMath      // Color math parameters
	m7             *m7             // mode 7 parameters
	display        *display
	window         [2]*window
	registers      [0x40]register // registers represents the ppu registers as methods

	hCounter uint16
	vCounter uint16
}

// New initializes a PPU struct and returns it
func New() *PPU {
	ppu := &PPU{}
	ppu.vram = &vram{}
	ppu.oam = &oam{}
	ppu.cgram = &cgram{}
	ppu.backgroundData = &backgroundData{}
	ppu.colorMath = &colorMath{}
	ppu.m7 = &m7{}
	ppu.display = &display{}
	return ppu
}

type register func(uint8) uint8
