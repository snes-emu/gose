package core

import "go.uber.org/zap"

func (ppu *PPU) renderLine() {
	ppu.lg.Debug("Render line", zap.Uint16("vCounter", ppu.vCounter))
	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter == ppu.VDisplay()+1 {
		ppu.lg.Debug("VBlank")
		ppu.cpu.enterVblank()
	}

	if ppu.vCounter == 0 {
		ppu.lg.Debug("End of VBlank")
		ppu.cpu.leavVblank()
	}
}
