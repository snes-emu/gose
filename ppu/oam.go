package ppu

type sprite struct {
	x            uint16 // x coordinate of the upper left tile
	y            uint16 // y coordinate of the upper left tile
	tileAddress  uint16 // address of first tile in the VRAM
	paletteIndex uint16 // index of the palette
	priority     uint8  // priority of the sprite
	hFlip        bool   // horizontal flip
	vFlip        bool   // vertical flip
	hSize        uint8  // horizontal sprite size
	vSize        uint8  // vertical sprite size
}

type oam struct {
	bytes           [0x200 + 0x20]byte // oam represents the object attribute memory (512 + 32 Bytes)
	addr            uint16             // the OAM addr p------b aaaaaaaa  (p is the Obj Priority activation bit and the rest represents the oam addr) stored as ba aaaaaaaf where f is the flip
	lastWrittenAddr uint16             // variable to hold the last written oam.addr
	priorityBit     bool               // Hold addr flip (even or odd part of a word)
	lsb             uint8              // temporary variable for the oamdata register

	objectSize            uint8  // index representing object size in pixel
	objectTileBaseAddress uint16 // Tile used for sprites base address in VRAM
	objectTileGapAddress  uint16 // Gap between object tile 0x0FF and 0x100 in VRAM
	windowMask1           uint8  // mask for window 1 (0..1=Disable, 2=Inside, 3=Outside)
	windowMask2           uint8  // mask for window 2 (0..1=Disable, 2=Inside, 3=Outside)
	windowMaskLogic       uint8  // 0=OR, 1=AND, 2=XOR, 3=XNOR)
	mainScreenWindow      bool   // Disable window area on main screen
	subScreenWindow       bool   // Disable windows area on sub screen
	mainScreen            bool   // Enable layer on main screen
	subScreen             bool   // Enable layer on sub screen
}

// 2102 - oam.aDDL
func (ppu *PPU) oamaddl(data uint8) uint8 {
	ppu.oam.addr = (ppu.oam.lastWrittenAddr & 0x0200) | (uint16(data) << 1)
	ppu.oam.lastWrittenAddr = ppu.oam.addr
	return 0
}

// 2103 - oam.aDDH
func (ppu *PPU) oamaddh(data uint8) uint8 {
	ppu.oam.priorityBit = data&0x80 != 0
	ppu.oam.addr = (uint16(data) << 9) | (ppu.oam.lastWrittenAddr & 0x01fe)
	ppu.oam.lastWrittenAddr = ppu.oam.addr
	return 0
}

// 2104 - OAMDATA - OAM Data Write (W)
func (ppu *PPU) oamdata(data uint8) uint8 {
	if ppu.oam.addr%2 == 0 {
		// Write to the temporary variable
		ppu.oam.lsb = data
	}
	if ppu.oam.addr > 0x1FF {
		ppu.oam.bytes[ppu.oam.addr] = data
	} else if ppu.oam.addr%2 == 1 {
		// Remove the Obj Priority activation bit and keep only the b aaaaaaaa part
		ppu.oam.bytes[ppu.oam.addr-1] = ppu.oam.lsb
		ppu.oam.bytes[ppu.oam.addr] = data
	}
	// Increment the address
	ppu.oam.addr = (ppu.oam.addr + 1) % 544
	return 0
}

// 2138 - RDOAM - OAM Data Read (R)
func (ppu *PPU) rdoam(_ uint8) uint8 {
	res := ppu.oam.bytes[ppu.oam.addr]
	ppu.oam.addr = (ppu.oam.addr + 1) % 544
	return res
}

// 2101h - OBSEL - Object Size and Object Base (W)
func (ppu *PPU) obsel(data uint8) uint8 {
	ppu.oam.objectSize = (data >> 5)
	ppu.oam.objectTileBaseAddress = uint16(data&0x7) << 14
	ppu.oam.objectTileGapAddress = uint16((data>>3)&0x3) << 13
	return 0
}

var spriteSizeTable = [16][2]uint8{
	{8, 8},
	{8, 8},
	{8, 8},
	{16, 16},
	{16, 16},
	{32, 32},
	{16, 32},
	{16, 32},
	{16, 16},
	{32, 32},
	{64, 64},
	{32, 32},
	{64, 64},
	{64, 64},
	{32, 64},
	{32, 32},
}

func (ppu PPU) decodeSprite(i uint16) sprite {
	sprite := sprite{}
	sprite.x = uint16(ppu.oam.bytes[4*i]) | uint16(ppu.oam.bytes[0x200+i/4]&(1<<(2*(i%4)))<<(8-2*(i%4)))
	sprite.y = uint16(ppu.oam.bytes[4*i+1])
	sprite.tileAddress = ((ppu.oam.objectTileBaseAddress << 14) + (uint16(ppu.oam.bytes[4*i+2]) << 5) + (uint16(ppu.oam.bytes[4*i+3])&0x1)*(ppu.oam.objectTileGapAddress+1)<<13) & 0xFFFE
	sprite.paletteIndex = (uint16(ppu.oam.bytes[4*i+3]) & 0xE) >> 1
	sprite.priority = (ppu.oam.bytes[4*i+3] & 0x30) >> 4
	sprite.hFlip = ppu.oam.bytes[4*i+3]&0x40 != 0
	sprite.vFlip = ppu.oam.bytes[4*i+3]&0x80 != 0
	size := spriteSizeTable[ppu.oam.objectSize|((ppu.oam.bytes[0x200+i/4]&(1<<(2*(i%4)+1)))>>(1<<(2*(i%4)+1)))<<4]
	sprite.hSize = size[0]
	sprite.vSize = size[1]
	return sprite
}
