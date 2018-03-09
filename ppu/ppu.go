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
	status         *status
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
	ppu.registers[0x00] = ppu.inidisp
	ppu.registers[0x01] = ppu.obsel
	ppu.registers[0x02] = ppu.oamaddl
	ppu.registers[0x03] = ppu.oamaddh
	ppu.registers[0x04] = ppu.oamdata
	ppu.registers[0x05] = ppu.bgmode
	ppu.registers[0x06] = ppu.mosaic
	ppu.registers[0x07] = ppu.bg1sc
	ppu.registers[0x08] = ppu.bg2sc
	ppu.registers[0x09] = ppu.bg3sc
	ppu.registers[0x0A] = ppu.bg4sc
	ppu.registers[0x0B] = ppu.bg12nba
	ppu.registers[0x0C] = ppu.bg34nba
	ppu.registers[0x0D] = ppu.bg1hofs
	ppu.registers[0x0E] = ppu.bg1vofs
	ppu.registers[0x0F] = ppu.bg2hofs
	ppu.registers[0x10] = ppu.bg2vofs
	ppu.registers[0x11] = ppu.bg3hofs
	ppu.registers[0x12] = ppu.bg3vofs
	ppu.registers[0x13] = ppu.bg4hofs
	ppu.registers[0x14] = ppu.bg4vofs
	ppu.registers[0x15] = ppu.vmain
	ppu.registers[0x16] = ppu.vmaddl
	ppu.registers[0x17] = ppu.vmaddh
	ppu.registers[0x18] = ppu.vmdatal
	ppu.registers[0x19] = ppu.vmdatah
	ppu.registers[0x1A] = ppu.m7sel
	ppu.registers[0x1B] = ppu.m7a
	ppu.registers[0x1C] = ppu.m7b
	ppu.registers[0x1D] = ppu.m7c
	ppu.registers[0x1E] = ppu.m7d
	ppu.registers[0x1F] = ppu.m7x
	ppu.registers[0x20] = ppu.m7y
	ppu.registers[0x21] = ppu.cgadd
	ppu.registers[0x22] = ppu.cgdata
	ppu.registers[0x23] = ppu.w12sel
	ppu.registers[0x24] = ppu.w34sel
	ppu.registers[0x25] = ppu.wobjsel
	ppu.registers[0x26] = ppu.wh0
	ppu.registers[0x27] = ppu.wh1
	ppu.registers[0x28] = ppu.wh2
	ppu.registers[0x29] = ppu.wh3
	ppu.registers[0x2A] = ppu.wbglog
	ppu.registers[0x2B] = ppu.wobjlog
	ppu.registers[0x2C] = ppu.tm
	ppu.registers[0x2D] = ppu.ts
	ppu.registers[0x2E] = ppu.tmw
	ppu.registers[0x2F] = ppu.tsw
	ppu.registers[0x30] = ppu.cgwsel
	ppu.registers[0x31] = ppu.cgadsub
	ppu.registers[0x32] = ppu.coldata
	ppu.registers[0x33] = ppu.setini
	//ppu.registers[0x34] = ppu.mpyl
	//ppu.registers[0x35] = ppu.mpym
	//ppu.registers[0x36] = ppu.mpyh
	ppu.registers[0x37] = ppu.slhv
	ppu.registers[0x38] = ppu.oamdata
	ppu.registers[0x39] = ppu.vmdatal
	ppu.registers[0x3A] = ppu.vmdatah
	ppu.registers[0x3B] = ppu.cgdata
	ppu.registers[0x3C] = ppu.ophct
	ppu.registers[0x3D] = ppu.opvct
	ppu.registers[0x3E] = ppu.stat77
	ppu.registers[0x3F] = ppu.stat78
	return ppu
}

type register func(uint8) uint8
