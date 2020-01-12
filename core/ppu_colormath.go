package core

import "github.com/snes-emu/gose/render"

type colorMath struct {
	mainScreenBlack uint8 // Force main screen black (possible values: (3=Always, 2=MathWindow, 1=NotMathWin, 0=Never))
	enable          uint8 // Global color math enable (possible values: (0=Always, 1=MathWindow, 2=NotMathWin, 3=Never))
	enableSubscreen bool  // Sub Screen BG/OBJ Enable    (0=No/Backdrop only, 1=Yes/Backdrop+BG+OBJ)
	directColor     bool  // Direct Color (for 256-color BGs)  (0=Use Palette, 1=Direct Color)
	red             uint8 // Red color
	blue            uint8 // Blue color
	green           uint8 // Green color
	opSign          int8  // Sign used for the math operation (possible values are -1 and 1)
	div2            bool  // Whether or not the colors should be divided by 2 (only in certain cases though)
	backdrop        bool  // If color math should be used when the main screen = Backdrop
	obj             bool  // If color math should be used when main screen = OBJ/Palette4..7
	windowMask1     uint8 // mask for window 1 (0..1=Disable, 2=Inside, 3=Outside)
	windowMask2     uint8 // mask for window 2 (0..1=Disable, 2=Inside, 3=Outside)
	windowMaskLogic uint8 // 0=OR, 1=AND, 2=XOR, 3=XNOR)
}

// 2130 - CGWSEL - Color Math Control Register A (W)
func (ppu *PPU) cgwsel(data uint8) {
	ppu.colorMath.mainScreenBlack = (data & 0xc0) >> 6
	ppu.colorMath.enable = (data & 0x30) >> 4
	ppu.colorMath.enableSubscreen = (data & 0x2) != 0
	ppu.colorMath.directColor = (data & 0x1) != 0
}

// 2131 - CGADSUB - Color Math Control Register B (W)
func (ppu *PPU) cgadsub(data uint8) {
	if (data & 0x80) != 0 {
		ppu.colorMath.opSign = -1
	} else {
		ppu.colorMath.opSign = 1
	}

	ppu.colorMath.div2 = (data & 0x40) != 0

	ppu.backgroundData.bg[3].colorMath = (data & 0x8) != 0
	ppu.backgroundData.bg[2].colorMath = (data & 0x4) != 0
	ppu.backgroundData.bg[1].colorMath = (data & 0x2) != 0
	ppu.backgroundData.bg[0].colorMath = (data & 0x1) != 0

	ppu.colorMath.backdrop = (data & 0x20) != 0
	ppu.colorMath.obj = (data & 0x10) != 0
}

// 2132 - COLDATA - Color Math Sub Screen Backdrop Color (W)
func (ppu *PPU) coldata(data uint8) {
	intensity := data & 0x1f

	if (data & 0x80) != 0 {
		ppu.colorMath.blue = intensity
	}
	if (data & 0x40) != 0 {
		ppu.colorMath.green = intensity
	}
	if (data & 0x20) != 0 {
		ppu.colorMath.red = intensity
	}
}

func (ppu *PPU) applyColorMath(base []render.Pixel) {
	for i := range base {
		if ppu.colorMath.opSign == 1 {
			base[i].Color = base[i].Color.Add(ppu.subScreen[i].Color)
		} else {
			base[i].Color = base[i].Color.Sub(ppu.subScreen[i].Color)
		}
		if ppu.colorMath.div2 {
			base[i].Color = base[i].Color.Halve()
		}
	}
}
