package core

// oam represents the object attribute memory, two tables (512 + 32 Bytes)
type oam struct {
	bytes           [0x200 + 0x20]byte // raw bytes for the object attribute memory
	addr            uint16             // the OAM addr p------b aaaaaaaa  (p is the Obj Priority activation bit and the rest represents the oam addr) stored as ba aaaaaaaf where f is the flip
	lastWrittenAddr uint16             // variable to hold the last written oam.addr
	priorityBit     bool               // Hold addr flip (even or odd part of a word)
	lsb             uint8              // temporary variable for the oamdata register

	objectSize            uint8  // index representing object size in pixel
	objectTileBaseAddress uint16 // Tile used for sprites base address in VRAM (16K bytes steps, 8k words)
	objectTileGapAddress  uint16 // Gap between object tile 0x0FF and 0x100 in VRAM (8K bytes steps, 4k words)
	windowMask1           uint8  // mask for window 1 (0..1=Disable, 2=Inside, 3=Outside)
	windowMask2           uint8  // mask for window 2 (0..1=Disable, 2=Inside, 3=Outside)
	windowMaskLogic       uint8  // 0=OR, 1=AND, 2=XOR, 3=XNOR)
	mainScreenWindow      bool   // Disable window area on main screen
	subScreenWindow       bool   // Disable windows area on sub screen
	mainScreen            bool   // Enable layer on main screen
	subScreen             bool   // Enable layer on sub screen
}

// TODO: have a look at "reload"
// 2102 - oam.aDDL
func (ppu *PPU) oamaddl(data uint8) {
	ppu.oam.addr = (ppu.oam.lastWrittenAddr & 0x0200) | (uint16(data) << 1)
	ppu.oam.lastWrittenAddr = ppu.oam.addr
}

// 2103 - oam.aDDH
func (ppu *PPU) oamaddh(data uint8) {
	ppu.oam.priorityBit = data&0x80 != 0
	ppu.oam.addr = (uint16(data & 0x1) << 8) | (ppu.oam.lastWrittenAddr & 0xff)
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
	ppu.oam.objectTileGapAddress = uint16((data >> 3) & 0x3)
	ppu.oam.objectTileBaseAddress = uint16(data & 0x7)
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
