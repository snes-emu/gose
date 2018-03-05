package ppu

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram      *vram              // vram represents the VideoRAM (64KB)
	oam       [0x200 + 0x20]byte // oam represents the object attribute memory (512 + 32 Bytes)
	cgram     [0x200]byte        // cgram represents the color graphics ram and stores the color palette with 256 color entries
	registers [0x40]register     // registers represents the ppu registers as methods

	cgramAddr uint16 // store the cgram address over 512 byte (not the Word addr !)
	cgramLsb  uint8  // temporary variable for the cgdata register

	oamAddr            uint16 // the OAM addr p------b aaaaaaaa  (p is the Obj Priority activation bit and the rest represents the oam addr) stored as ba aaaaaaaf where f is the flip
	oamLastWrittenAddr uint16 // variable to hold the last written oamAddr
	oamPriorityBit     bool   // Hold addr flip (even or odd part of a word)
	oamLsb             uint8  // temporary variable for the oamdata register

	objectSize            uint8  // index representing object size in pixel
	objectTileBaseAddress uint16 // Tile used for sprites base address in VRAM
	objectTileGapAddress  uint16 // Gap between object tile 0x0FF and 0x100 in VRAM

	backgroundData *backgroundData // background data
	colorMath      *colorMath      // Color math parameters
}

type register func(uint8) uint8
