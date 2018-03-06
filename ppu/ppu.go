package ppu

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram           *vram           // vram represents the VideoRAM (64KB)
	oam            *oam            // oam represents the object attribute memory (512 + 32 Bytes)
	cgram          *cgram          // cgram represents the color graphics ram and stores the color palette with 256 color entries
	backgroundData *backgroundData // background data
	colorMath      *colorMath      // Color math parameters
	m7             *m7             // mode 7 parameters

	registers [0x40]register // registers represents the ppu registers as methods
	window    [2]*window
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
	return ppu
}

type register func(uint8) uint8
