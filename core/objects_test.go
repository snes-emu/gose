package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpriteSizeTable(t *testing.T) {
	spriteSizeTable := [16][2]uint8{
		// TODO: this should be 8, 8
		{16, 16},
		{8, 8},
		{8, 8},
		{16, 16},
		{16, 16},
		{32, 32},
		{16, 32},
		{16, 32},
		{16, 16},
		{32, 32},
		{64, 64},
		{32, 32},
		{64, 64},
		{64, 64},
		{32, 64},
		{32, 32},
	}

	for sizeFlag := 0; sizeFlag <= 1; sizeFlag++ {
		for objectSize := 0; objectSize < 8; objectSize++ {
			old := spriteSizeTable[objectSize|sizeFlag<<3]
			h, v := spriteSize(sizeFlag == 1, uint8(objectSize))
			assert.EqualValuesf(t, old, [2]uint8{uint8(h), uint8(v)}, "for object size: %d, sizeFlag: %t", objectSize, sizeFlag == 1)
		}
	}
}
