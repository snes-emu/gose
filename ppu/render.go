package ppu

type pixel struct {
	rgb     uint32
	visible bool
}

func (ppu PPU) renderSpriteLine() [HMax]pixel {
	var pixels [HMax]pixel
	for i := uint8(0); i < 128; i++ {
		y := ppu.oam.bytes[4*i+1]
		x := uint16(ppu.oam.bytes[4*i+1]) | (((uint16(ppu.oam.bytes[512+uint16(i)/4]) & (2 * (uint16(i) % 4))) >> (2 * (uint16(i) % 4))) << 8)
		// Fetch the sprite size and put the bit at index 4
		size := ((ppu.oam.bytes[512+uint16(i)/4] & (2*(i%4) + 1)) >> (2*(i%4) + 1)) << 4
		xSize := uint16(spriteSizeTable[size|ppu.oam.objectSize][0])
		ySize := spriteSizeTable[size|ppu.oam.objectSize][1]
		if uint8(ppu.vCounter) >= y && uint8(ppu.vCounter) < y+ySize {
			// characterIndex := uint16(ppu.oam.bytes[4*i+2])
			// paletteIndex := (ppu.oam.bytes[4*i+3] >> 1) & 0x07
			// tileAddress := ((ppu.oam.objectTileBaseAddress << 14) + (characterIndex << 5) + (uint16(ppu.oam.bytes[4*i+3])&0x1)*(ppu.oam.objectTileGapAddress+1)<<13) & 0xFFFE
			for pixel := x; pixel < x+xSize; pixel++ {
			}
		}
	}
	return pixels
}
