package core

import "fmt"

func (ppu *PPU) RenderLine() {
	fmt.Printf("Render line: %v\n", ppu.vCounter)
	ppu.vCounter = (ppu.vCounter + 1) % ppu.VDisplayEnd()

	if ppu.vCounter == ppu.VDisplay()+1 {
		// TODO Trigger cpu vblank here
		fmt.Println("VBlank !")
	}
}
