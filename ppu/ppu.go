package ppu

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram      *vram          // vram represents the VideoRAM (64KB)
	oam       *oam           // oam represents the object attribute memory (512 + 32 Bytes)
	cgram     [0x200]byte    // cgram represents the color graphics ram and stores the color palette with 256 color entries
	registers [0x40]register // registers represents the ppu registers as methods

	cgramAddr uint16 // store the cgram address over 512 byte (not the Word addr !)
	cgramLsb  uint8  // temporary variable for the cgdata register

	backgroundData *backgroundData // background data
	colorMath      *colorMath      // Color math parameters

	m7 *m7 // mode 7 parameters
}

type register func(uint8) uint8
