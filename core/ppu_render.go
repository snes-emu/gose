package core

import "fmt"

func (ppu *PPU) renderLine() {
	fmt.Printf("Render line: %v\n", ppu.vCounter)
	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter == ppu.VDisplay()+1 {
		fmt.Println("VBlank !")
		ppu.cpu.enterVblank()
	}

	if ppu.vCounter == 0 {
		ppu.cpu.leavVblank()
	}
}
