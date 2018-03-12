package ppu

type pixel struct {
	rgb     uint32
	visible bool
}

func (ppu PPU) renderSpriteLine() [HMax]pixel {
	var pixels [HMax]pixel
	for i := uint16(0); i < 128; i++ {
		sprite := ppu.decodeSprite(i)
		if ppu.vCounter >= sprite.y && ppu.vCounter < sprite.y+uint16(sprite.vSize) {
			for pixel := sprite.x; pixel < sprite.x+uint16(sprite.hSize); pixel++ {
			}
		}
	}
	return pixels
}
