package core

// oam represents the object attribute memory, two tables (512 + 32 Bytes)
type oam struct {
	bytes           [0x200 + 0x20]byte // raw bytes for the object attribute memory
	addr            uint16             // the OAM addr p------b aaaaaaaa  (p is the Obj Priority activation bit and the rest represents the oam addr) stored as ba aaaaaaaf where f is the flip
	lastWrittenAddr uint16             // variable to hold the last written oam.addr
	priorityBit     bool               // Hold addr flip (even or odd part of a word)
	lsb             uint8              // temporary variable for the oamdata register

	objectSize       uint8  // index representing object size in pixel
	baseAddr         uint16 // Tile used for sprites base address in VRAM (16K bytes steps, 8k words)
	nameSelect       uint16 // Gap between object tile 0x0FF and 0x100 in VRAM (8K bytes steps, 4k words)
	windowMask1      uint8  // mask for window 1 (0..1=Disable, 2=Inside, 3=Outside)
	windowMask2      uint8  // mask for window 2 (0..1=Disable, 2=Inside, 3=Outside)
	windowMaskLogic  uint8  // 0=OR, 1=AND, 2=XOR, 3=XNOR)
	mainScreenWindow bool   // Disable window area on main screen
	subScreenWindow  bool   // Disable windows area on sub screen
	mainScreen       bool   // Enable layer on main screen
	subScreen        bool   // Enable layer on sub screen
}

// TODO: have a look at "reload"
// 2102 - oam.aDDL
func (ppu *PPU) oamaddl(data uint8) {
	ppu.oam.addr = (ppu.oam.lastWrittenAddr & 0x0200) | uint16(data)
	ppu.oam.lastWrittenAddr = ppu.oam.addr
}

// 2103 - oam.aDDH
func (ppu *PPU) oamaddh(data uint8) {
	ppu.oam.priorityBit = data&0x80 != 0
	ppu.oam.addr = (uint16(data&0x1) << 8) | (ppu.oam.lastWrittenAddr & 0xff)
	ppu.oam.lastWrittenAddr = ppu.oam.addr
}

// 2104 - OAMDATA - OAM Data Write (W)
// Reads and Writes to EVEN and ODD byte-addresses work as follows:
//
// Write to EVEN address      -->  set OAM_Lsb = Data    ;memorize value
// Write to ODD address<200h  -->  set WORD[addr-1] = Data*256 + OAM_Lsb
// Write to ANY address>1FFh  -->  set BYTE[addr] = Data
// Read from ANY address      -->  return BYTE[addr]
func (ppu *PPU) oamdata(data uint8) {
	if ppu.oam.addr%2 == 0 {
		// Write to the temporary variable
		ppu.oam.lsb = data
	}

	// Check if we are going to write in the first or second table
	// 0x1FF == 511
	if ppu.oam.addr > 0x1FF {
		// Writing in the second table
		ppu.oam.bytes[ppu.oam.addr] = data
	} else if ppu.oam.addr%2 == 1 {
		// Writing in the first table
		ppu.oam.write(ppu.oam.addr-1, ppu.oam.lsb, data)
	}
	// Increment the address
	ppu.oam.incrAddr()
}

// 2138 - RDOAM - OAM Data Read (R)
func (ppu *PPU) rdoam() uint8 {
	res := ppu.oam.read(ppu.oam.addr)
	ppu.oam.incrAddr()
	return res
}

// 2101h - OBSEL - Object Size and Object Base (W)
// 7-5   OBJ Size Selection  (0-5, see below) (6-7=Reserved)
// Val Small  Large
// 0 = 8x8    16x16    ;Caution:
// 1 = 8x8    32x32    ;In 224-lines mode, OBJs with 64-pixel height
// 2 = 8x8    64x64    ;may wrap from lower to upper screen border.
// 3 = 16x16  32x32    ;In 239-lines mode, the same problem applies
// 4 = 16x16  64x64    ;also for OBJs with 32-pixel height.
// 5 = 32x32  64x64
// 6 = 16x32  32x64 (undocumented)
// 7 = 16x32  32x32 (undocumented)
// (Ie. a setting of 0 means Small OBJs=8x8, Large OBJs=16x16 pixels)
// (Whether an OBJ is "small" or "large" is selected by a bit in OAM)
// 4-3   Gap between OBJ 0FFh and 100h (0=None) (4K-word steps) (8K-byte steps)
// 2-0   Base Address for OBJ Tiles 000h..0FFh  (8K-word steps) (16K-byte steps)
func (ppu *PPU) obsel(data uint8) {
	ppu.oam.objectSize = (data >> 5)
	ppu.oam.nameSelect = uint16((data >> 3) & 0x3)
	ppu.oam.baseAddr = uint16(data & 0x7)
}

func (o *oam) read(addr uint16) uint8 {
	return o.bytes[addr]
}

func (o *oam) write(addr uint16, low uint8, high uint8) {
	o.bytes[addr+1] = high
	o.bytes[addr] = low
}

func (o *oam) incrAddr() {
	o.addr = (o.addr + 1) % 544
}

// allSprites returns all the sprites currently stored in the OAM
func (o *oam) allSprites() []sprite {
	sprites := make([]sprite, 128)

	for i := range sprites {
		sprites[i] = o.sprite(uint16(i))
	}

	return sprites
}

// intersectingSprites returns all the sprites currently intersecting the v-line
func (o *oam) intersectingSprites(vCounter uint16) []sprite {
	sprites := make([]sprite, 0, 128)

	for i := 0; i < 128; i++ {
		s := o.sprite(uint16(i))
		if s.IntersectsLine(vCounter) {
			sprites = append(sprites, o.sprite(uint16(i)))
		}
	}

	return sprites
}

// sprite gets the sprite at the given index
// the oam stores 128 entries in the following format:
//
// Table 1 (4-bytes per sprite) -> 512 bytes:
// Byte 1:    xxxxxxxx    x: X coordinate
// Byte 2:    yyyyyyyy    y: Y coordinate
// Byte 3:    cccccccc    c: Starting tile #
// Byte 4:    vhoopppc    v: vertical flip h: horizontal flip  o: priority bits  p: palette
//
// Table 2 (2 bits per sprite) -> 32 bytes:
// Bit 0: upper X coordinate bit
// Bit 1: Sprite size (0 is small, 1 is large)
func (o *oam) sprite(idx uint16) sprite {
	// raw1 contains the 4 bytes from the Table 1 for the given index
	raw1 := o.bytes[4*idx : 4*(idx+1)]

	// raw2 contains the two bits of interest from the Table 2 for the given index
	raw2 := (o.bytes[0x200+(idx/4)] >> (2 * (idx % 4))) & 0x3

	attrs := raw1[3]

	sprite := sprite{}

	// sprites always use 16 colors
	sprite.colorDepth = 4

	// Read x, y, and firstTileAddress low word
	sprite.x = uint16(raw1[0])
	sprite.y = uint16(raw1[1])
	tileIdx := uint16(raw1[2])

	// Add upper bits
	sprite.x |= uint16(raw2&0x1) << 8
	tileIdx |= uint16(attrs&0x1) << 8

	// Vertical and Horizontal flips
	sprite.hFlip = attrs&0x40 != 0
	sprite.vFlip = attrs&0x80 != 0

	// Priority and palette
	sprite.priority = (attrs >> 4) & 0x3
	// Sprite palette starts at 128
	sprite.palette = 128 + (16 * ((attrs >> 1) & 0x7))

	// Sprite size
	isLarge := (raw2 & 0x2) != 0
	sprite.hSize, sprite.vSize = spriteSize(isLarge, o.objectSize)

	// Base Address is in 16K bytes steps
	// NameSelect is in 8k bytes steps
	// Formula is:
	// ((Base << 14) + (index << 5) + (N ? ((Name)<<13) : 0))
	// Where N is the upper bit of the tile index
	// The & 0x7fff is just to limit the range to 32KB
	// See: https://wiki.superfamicom.org/sprites
	// The formula in wiki.superfamicom.com is given as word address (hence 2 bytes)
	// that's why they limit the result to 32KB
	sprite.addr = (o.baseAddr << 14) + (tileIdx << 5)
	if attrs&0x1 != 0 {
		sprite.addr += o.nameSelect << 13
	}

	return sprite
}
