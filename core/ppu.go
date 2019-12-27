package core

import (
	"github.com/snes-emu/gose/io"
	"github.com/snes-emu/gose/render"
)

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

	cpu         *CPU
	renderer    render.Renderer
	screen      *render.Screen
	blackScreen *render.Screen
}

// New initializes a PPU struct and returns it
func newPPU(renderer render.Renderer, rf *io.RegisterFactory) *PPU {
	ppu := &PPU{}
	ppu.renderer = renderer
	// TODO: fix dimensions
	ppu.screen = render.NewScreen(WIDTH, HEIGHT)
	ppu.blackScreen = render.NewScreen(WIDTH, HEIGHT)
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

	ppu.Registers[0x00] = rf.NewRegister(nil, ppu.inidisp, "INIDISP")
	ppu.Registers[0x01] = rf.NewRegister(nil, ppu.obsel, "OBSEL")
	ppu.Registers[0x02] = rf.NewRegister(nil, ppu.oamaddl, "OAMADDL")
	ppu.Registers[0x03] = rf.NewRegister(nil, ppu.oamaddh, "OAMADDH")
	ppu.Registers[0x04] = rf.NewRegister(nil, ppu.oamdata, "OAMDATA")
	ppu.Registers[0x05] = rf.NewRegister(nil, ppu.bgmode, "BGMODE")
	ppu.Registers[0x06] = rf.NewRegister(nil, ppu.mosaic, "MOSAIC")
	ppu.Registers[0x07] = rf.NewRegister(nil, ppu.bg1sc, "BG1SC")
	ppu.Registers[0x08] = rf.NewRegister(nil, ppu.bg2sc, "BG2SC")
	ppu.Registers[0x09] = rf.NewRegister(nil, ppu.bg3sc, "BG3SC")
	ppu.Registers[0x0A] = rf.NewRegister(nil, ppu.bg4sc, "BG4SC")
	ppu.Registers[0x0B] = rf.NewRegister(nil, ppu.bg12nba, "BG12NBA")
	ppu.Registers[0x0C] = rf.NewRegister(nil, ppu.bg34nba, "BG34NBA")
	ppu.Registers[0x0D] = rf.NewRegister(nil, ppu.bg1hofs, "BG1HOFS")
	ppu.Registers[0x0E] = rf.NewRegister(nil, ppu.bg1vofs, "BG1VOFS")
	ppu.Registers[0x0F] = rf.NewRegister(nil, ppu.bg2hofs, "BG2HOFS")
	ppu.Registers[0x10] = rf.NewRegister(nil, ppu.bg2vofs, "BG2VOFS")
	ppu.Registers[0x11] = rf.NewRegister(nil, ppu.bg3hofs, "BG3HOFS")
	ppu.Registers[0x12] = rf.NewRegister(nil, ppu.bg3vofs, "BG3VOFS")
	ppu.Registers[0x13] = rf.NewRegister(nil, ppu.bg4hofs, "BG4HOFS")
	ppu.Registers[0x14] = rf.NewRegister(nil, ppu.bg4vofs, "BG4VOFS")
	ppu.Registers[0x15] = rf.NewRegister(nil, ppu.vmain, "VMAIN")
	ppu.Registers[0x16] = rf.NewRegister(nil, ppu.vmaddl, "VMADDL")
	ppu.Registers[0x17] = rf.NewRegister(nil, ppu.vmaddh, "VMADDH")
	ppu.Registers[0x18] = rf.NewRegister(nil, ppu.vmdatal, "VMDATAL")
	ppu.Registers[0x19] = rf.NewRegister(nil, ppu.vmdatah, "VMDATAH")
	ppu.Registers[0x1A] = rf.NewRegister(nil, ppu.m7sel, "M7SEL")
	ppu.Registers[0x1B] = rf.NewRegister(nil, ppu.m7a, "M7A")
	ppu.Registers[0x1C] = rf.NewRegister(nil, ppu.m7b, "M7B")
	ppu.Registers[0x1D] = rf.NewRegister(nil, ppu.m7c, "M7C")
	ppu.Registers[0x1E] = rf.NewRegister(nil, ppu.m7d, "M7D")
	ppu.Registers[0x1F] = rf.NewRegister(nil, ppu.m7x, "M7X")
	ppu.Registers[0x20] = rf.NewRegister(nil, ppu.m7y, "M7Y")
	ppu.Registers[0x21] = rf.NewRegister(nil, ppu.cgadd, "CGADD")
	ppu.Registers[0x22] = rf.NewRegister(nil, ppu.cgdata, "CGDATA")
	ppu.Registers[0x23] = rf.NewRegister(nil, ppu.w12sel, "W12SEL")
	ppu.Registers[0x24] = rf.NewRegister(nil, ppu.w34sel, "W34SEL")
	ppu.Registers[0x25] = rf.NewRegister(nil, ppu.wobjsel, "WOBJSEL")
	ppu.Registers[0x26] = rf.NewRegister(nil, ppu.wh0, "WH0")
	ppu.Registers[0x27] = rf.NewRegister(nil, ppu.wh1, "WH1")
	ppu.Registers[0x28] = rf.NewRegister(nil, ppu.wh2, "WH2")
	ppu.Registers[0x29] = rf.NewRegister(nil, ppu.wh3, "WH3")
	ppu.Registers[0x2A] = rf.NewRegister(nil, ppu.wbglog, "WBGLOG")
	ppu.Registers[0x2B] = rf.NewRegister(nil, ppu.wobjlog, "WOBJLOG")
	ppu.Registers[0x2C] = rf.NewRegister(nil, ppu.tm, "TM")
	ppu.Registers[0x2D] = rf.NewRegister(nil, ppu.ts, "TS")
	ppu.Registers[0x2E] = rf.NewRegister(nil, ppu.tmw, "TMW")
	ppu.Registers[0x2F] = rf.NewRegister(nil, ppu.tsw, "TSW")
	ppu.Registers[0x30] = rf.NewRegister(nil, ppu.cgwsel, "CGWSEL")
	ppu.Registers[0x31] = rf.NewRegister(nil, ppu.cgadsub, "CGADSUB")
	ppu.Registers[0x32] = rf.NewRegister(nil, ppu.coldata, "COLDATA")
	ppu.Registers[0x33] = rf.NewRegister(nil, ppu.setini, "SETINI")
	ppu.Registers[0x34] = rf.NewRegister(ppu.mpyl, nil, "MPYL")
	ppu.Registers[0x35] = rf.NewRegister(ppu.mpym, nil, "MPYM")
	ppu.Registers[0x36] = rf.NewRegister(ppu.mpyh, nil, "MPYH")
	ppu.Registers[0x37] = rf.NewRegister(ppu.slhv, nil, "SLHV")
	ppu.Registers[0x38] = rf.NewRegister(ppu.rdoam, nil, "RDOAM")
	ppu.Registers[0x39] = rf.NewRegister(ppu.rdvraml, nil, "RDVRAML")
	ppu.Registers[0x3A] = rf.NewRegister(ppu.rdvramh, nil, "RDVRAMH")
	ppu.Registers[0x3B] = rf.NewRegister(ppu.rdcgram, nil, "RDCGRAM")
	ppu.Registers[0x3C] = rf.NewRegister(ppu.ophct, nil, "OPHCT")
	ppu.Registers[0x3D] = rf.NewRegister(ppu.opvct, nil, "OPVCT")
	ppu.Registers[0x3E] = rf.NewRegister(ppu.stat77, nil, "STAT77")
	ppu.Registers[0x3F] = rf.NewRegister(ppu.stat78, nil, "STAT78")
	return ppu
}
