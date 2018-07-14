package ppu

import "github.com/snes-emu/gose/io"

const (
	// HMax represents max H counter value
	HMax = 339
	// VBSNTSC represents VBlank start counter value in NTSC
	VBSNTSC = 224
	// VBSPAL represents VBlank start counter value in PAL
	VBSPAL = 239
	// VMaxNTSC represents max V counter value in NTSC
	VMaxNTSC = 261
	// VMaxPAL represents max V counter value in PAL
	VMaxPAL = 311
)

// PPU represents the Picture Processing Unit of the SNES
type PPU struct {
	vram           *vram              // vram represents the VideoRAM (64KB)
	oam            *oam               // oam represents the object attribute memory (512 + 32 Bytes)
	cgram          *cgram             // cgram represents the color graphics ram and stores the color palette with 256 color entries
	backgroundData *backgroundData    // background data
	colorMath      *colorMath         // Color math parameters
	m7             *m7                // mode 7 parameters
	display        *display           // display parameters
	window         [2]*window         // window parameters
	status         *status            // store ppu status
	Registers      [0x40]*io.Register // Registers represents the ppu registers as methods

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
	ppu.Registers[0x00] = io.NewRegister(nil, ppu.inidisp)
	ppu.Registers[0x01] = io.NewRegister(nil, ppu.obsel)
	ppu.Registers[0x02] = io.NewRegister(nil, ppu.oamaddl)
	ppu.Registers[0x03] = io.NewRegister(nil, ppu.oamaddh)
	ppu.Registers[0x04] = io.NewRegister(nil, ppu.oamdata)
	ppu.Registers[0x05] = io.NewRegister(nil, ppu.bgmode)
	ppu.Registers[0x06] = io.NewRegister(nil, ppu.mosaic)
	ppu.Registers[0x07] = io.NewRegister(nil, ppu.bg1sc)
	ppu.Registers[0x08] = io.NewRegister(nil, ppu.bg2sc)
	ppu.Registers[0x09] = io.NewRegister(nil, ppu.bg3sc)
	ppu.Registers[0x0A] = io.NewRegister(nil, ppu.bg4sc)
	ppu.Registers[0x0B] = io.NewRegister(nil, ppu.bg12nba)
	ppu.Registers[0x0C] = io.NewRegister(nil, ppu.bg34nba)
	ppu.Registers[0x0D] = io.NewRegister(nil, ppu.bg1hofs)
	ppu.Registers[0x0E] = io.NewRegister(nil, ppu.bg1vofs)
	ppu.Registers[0x0F] = io.NewRegister(nil, ppu.bg2hofs)
	ppu.Registers[0x10] = io.NewRegister(nil, ppu.bg2vofs)
	ppu.Registers[0x11] = io.NewRegister(nil, ppu.bg3hofs)
	ppu.Registers[0x12] = io.NewRegister(nil, ppu.bg3vofs)
	ppu.Registers[0x13] = io.NewRegister(nil, ppu.bg4hofs)
	ppu.Registers[0x14] = io.NewRegister(nil, ppu.bg4vofs)
	ppu.Registers[0x15] = io.NewRegister(nil, ppu.vmain)
	ppu.Registers[0x16] = io.NewRegister(nil, ppu.vmaddl)
	ppu.Registers[0x17] = io.NewRegister(nil, ppu.vmaddh)
	ppu.Registers[0x18] = io.NewRegister(nil, ppu.vmdatal)
	ppu.Registers[0x19] = io.NewRegister(nil, ppu.vmdatah)
	ppu.Registers[0x1A] = io.NewRegister(nil, ppu.m7sel)
	ppu.Registers[0x1B] = io.NewRegister(nil, ppu.m7a)
	ppu.Registers[0x1C] = io.NewRegister(nil, ppu.m7b)
	ppu.Registers[0x1D] = io.NewRegister(nil, ppu.m7c)
	ppu.Registers[0x1E] = io.NewRegister(nil, ppu.m7d)
	ppu.Registers[0x1F] = io.NewRegister(nil, ppu.m7x)
	ppu.Registers[0x20] = io.NewRegister(nil, ppu.m7y)
	ppu.Registers[0x21] = io.NewRegister(nil, ppu.cgadd)
	ppu.Registers[0x22] = io.NewRegister(nil, ppu.cgdata)
	ppu.Registers[0x23] = io.NewRegister(nil, ppu.w12sel)
	ppu.Registers[0x24] = io.NewRegister(nil, ppu.w34sel)
	ppu.Registers[0x25] = io.NewRegister(nil, ppu.wobjsel)
	ppu.Registers[0x26] = io.NewRegister(nil, ppu.wh0)
	ppu.Registers[0x27] = io.NewRegister(nil, ppu.wh1)
	ppu.Registers[0x28] = io.NewRegister(nil, ppu.wh2)
	ppu.Registers[0x29] = io.NewRegister(nil, ppu.wh3)
	ppu.Registers[0x2A] = io.NewRegister(nil, ppu.wbglog)
	ppu.Registers[0x2B] = io.NewRegister(nil, ppu.wobjlog)
	ppu.Registers[0x2C] = io.NewRegister(nil, ppu.tm)
	ppu.Registers[0x2D] = io.NewRegister(nil, ppu.ts)
	ppu.Registers[0x2E] = io.NewRegister(nil, ppu.tmw)
	ppu.Registers[0x2F] = io.NewRegister(nil, ppu.tsw)
	ppu.Registers[0x30] = io.NewRegister(nil, ppu.cgwsel)
	ppu.Registers[0x31] = io.NewRegister(nil, ppu.cgadsub)
	ppu.Registers[0x32] = io.NewRegister(nil, ppu.coldata)
	ppu.Registers[0x33] = io.NewRegister(nil, ppu.setini)
	ppu.Registers[0x34] = io.NewRegister(ppu.mpyl, nil)
	ppu.Registers[0x35] = io.NewRegister(ppu.mpym, nil)
	ppu.Registers[0x36] = io.NewRegister(ppu.mpyh, nil)
	ppu.Registers[0x37] = io.NewRegister(ppu.slhv, nil)
	ppu.Registers[0x38] = io.NewRegister(ppu.rdoam, nil)
	ppu.Registers[0x39] = io.NewRegister(ppu.rdvraml, nil)
	ppu.Registers[0x3A] = io.NewRegister(ppu.rdvramh, nil)
	ppu.Registers[0x3B] = io.NewRegister(ppu.rdcgram, nil)
	ppu.Registers[0x3C] = io.NewRegister(ppu.ophct, nil)
	ppu.Registers[0x3D] = io.NewRegister(ppu.opvct, nil)
	ppu.Registers[0x3E] = io.NewRegister(ppu.stat77, nil)
	ppu.Registers[0x3F] = io.NewRegister(ppu.stat78, nil)
	return ppu
}
