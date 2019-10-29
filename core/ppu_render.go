package core

import (
	"github.com/snes-emu/gose/log"
	"go.uber.org/zap"
)

func (ppu *PPU) renderLine() {
	log.Debug("Render line", zap.Uint16("vCounter", ppu.vCounter))
	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter == ppu.VDisplay()+1 {
		log.Debug("VBlank")
		ppu.cpu.enterVblank()
	}

	if ppu.vCounter == 0 {
		log.Debug("End of VBlank")
		ppu.cpu.leavVblank()
	}
}
