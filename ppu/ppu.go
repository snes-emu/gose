package ppu

const (
	// HMax represents max H counter value
	HMax = 339
	// VMaxNTSC represents max V counter value in NTSC
	VMaxNTSC = 261
	// VMaxPAL represents max V counter value in PAL
	VMaxPAL = 311
)

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram           *vram           // vram represents the VideoRAM (64KB)
	oam            *oam            // oam represents the object attribute memory (512 + 32 Bytes)
	cgram          *cgram          // cgram represents the color graphics ram and stores the color palette with 256 color entries
	backgroundData *backgroundData // background data
	colorMath      *colorMath      // Color math parameters
	m7             *m7             // mode 7 parameters
	display        *display        // display parameters
	window         [2]*window      // window parameters
	status         *status         // store ppu status
	Registers      [0x40]register  // Registers represents the ppu registers as methods

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
	for i := 0; i < 4; i++ {
		ppu.backgroundData.bg[i] = &bg{}
	}
	ppu.colorMath = &colorMath{}
	ppu.m7 = &m7{}
	ppu.display = &display{}
	ppu.window[0] = &window{}
	ppu.window[1] = &window{}
	ppu.status = &status{}
	ppu.Registers[0x00] = ppu.inidisp
	ppu.Registers[0x01] = ppu.obsel
	ppu.Registers[0x02] = ppu.oamaddl
	ppu.Registers[0x03] = ppu.oamaddh
	ppu.Registers[0x04] = ppu.oamdata
	ppu.Registers[0x05] = ppu.bgmode
	ppu.Registers[0x06] = ppu.mosaic
	ppu.Registers[0x07] = ppu.bg1sc
	ppu.Registers[0x08] = ppu.bg2sc
	ppu.Registers[0x09] = ppu.bg3sc
	ppu.Registers[0x0A] = ppu.bg4sc
	ppu.Registers[0x0B] = ppu.bg12nba
	ppu.Registers[0x0C] = ppu.bg34nba
	ppu.Registers[0x0D] = ppu.bg1hofs
	ppu.Registers[0x0E] = ppu.bg1vofs
	ppu.Registers[0x0F] = ppu.bg2hofs
	ppu.Registers[0x10] = ppu.bg2vofs
	ppu.Registers[0x11] = ppu.bg3hofs
	ppu.Registers[0x12] = ppu.bg3vofs
	ppu.Registers[0x13] = ppu.bg4hofs
	ppu.Registers[0x14] = ppu.bg4vofs
	ppu.Registers[0x15] = ppu.vmain
	ppu.Registers[0x16] = ppu.vmaddl
	ppu.Registers[0x17] = ppu.vmaddh
	ppu.Registers[0x18] = ppu.vmdatal
	ppu.Registers[0x19] = ppu.vmdatah
	ppu.Registers[0x1A] = ppu.m7sel
	ppu.Registers[0x1B] = ppu.m7a
	ppu.Registers[0x1C] = ppu.m7b
	ppu.Registers[0x1D] = ppu.m7c
	ppu.Registers[0x1E] = ppu.m7d
	ppu.Registers[0x1F] = ppu.m7x
	ppu.Registers[0x20] = ppu.m7y
	ppu.Registers[0x21] = ppu.cgadd
	ppu.Registers[0x22] = ppu.cgdata
	ppu.Registers[0x23] = ppu.w12sel
	ppu.Registers[0x24] = ppu.w34sel
	ppu.Registers[0x25] = ppu.wobjsel
	ppu.Registers[0x26] = ppu.wh0
	ppu.Registers[0x27] = ppu.wh1
	ppu.Registers[0x28] = ppu.wh2
	ppu.Registers[0x29] = ppu.wh3
	ppu.Registers[0x2A] = ppu.wbglog
	ppu.Registers[0x2B] = ppu.wobjlog
	ppu.Registers[0x2C] = ppu.tm
	ppu.Registers[0x2D] = ppu.ts
	ppu.Registers[0x2E] = ppu.tmw
	ppu.Registers[0x2F] = ppu.tsw
	ppu.Registers[0x30] = ppu.cgwsel
	ppu.Registers[0x31] = ppu.cgadsub
	ppu.Registers[0x32] = ppu.coldata
	ppu.Registers[0x33] = ppu.setini
	ppu.Registers[0x34] = ppu.mpyl
	ppu.Registers[0x35] = ppu.mpym
	ppu.Registers[0x36] = ppu.mpyh
	ppu.Registers[0x37] = ppu.slhv
	ppu.Registers[0x38] = ppu.oamdata
	ppu.Registers[0x39] = ppu.vmdatal
	ppu.Registers[0x3A] = ppu.vmdatah
	ppu.Registers[0x3B] = ppu.cgdata
	ppu.Registers[0x3C] = ppu.ophct
	ppu.Registers[0x3D] = ppu.opvct
	ppu.Registers[0x3E] = ppu.stat77
	ppu.Registers[0x3F] = ppu.stat78
	return ppu
}

type register func(uint8) uint8
