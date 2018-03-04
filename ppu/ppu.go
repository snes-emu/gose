package ppu

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram      [0x10000]byte      // vram represents the VideoRAM (64KB)
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

	bg            [4]BG // BG array containing the 4 backgrounds
	bgScrollPrev1 uint8 // temporary variable for bg scrolling
	bgScrollPrev2 uint8 // temporary variable for bg scrolling
	bgScreenMode  uint8 // Screen mode from 0 to 7
	mosaicSize    uint8 // Size of block in mosaic mode (0=Smallest/1x1, 0xF=Largest/16x16)

	m7VerticalFlip                         bool   // Vertical flip flag used in mode7
	m7HorizontalFlip                       bool   // Horizontal flip flag used in mode7
	m7ScreenOver                           uint8  // Mode 7 screen over variable (possible values are 0,1,2 or 3)
	m7Cache                                uint16 // Mode 7 cache value used in various registers like M7A or M7HOFS
	m7aParam, m7bParam, m7cParam, m7dParam uint16 // Rotation/scaling parameters used in mode 7

}

type register func(uint8) uint8
