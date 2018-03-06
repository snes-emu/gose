package ppu

import (
	"testing"
)

func TestOam(t *testing.T) {
	// Set oam addr register to 0x104
	ppu := New()
	ppu.oamaddl(0x04)
	ppu.oamaddh(0x1)

	// Write 4 bytes
	ppu.oamdata(1)
	ppu.oamdata(2)
	ppu.oamdata(7)
	ppu.oamdata(12)

	// Write 1 to 0x2103
	ppu.oamaddh(0x1)

	if ppu.oam.lastWrittenAddr%512 != 2*4 {
		t.Errorf("Wrong value for internal oam address, expected 8, got: %v", ppu.oam.lastWrittenAddr)
	}
}
