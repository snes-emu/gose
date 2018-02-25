package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestBrk(t *testing.T) {
	mem := memory.New()
	mem.SetByteBank(0xab, 0x00, 0xffe6)
	mem.SetByteBank(0xcd, 0x00, 0xffe7)

	mem2 := memory.New()
	mem2.SetByteBank(0xab, 0x00, 0xffe6)
	mem2.SetByteBank(0xcd, 0x00, 0xffe7)
	mem2.SetByteBank(0x12, 0x00, 0x01ff)
	mem2.SetByteBank(0x34, 0x00, 0x01fe)
	mem2.SetByteBank(0x58, 0x00, 0x01fd)
	mem2.SetByteBank(0x08, 0x00, 0x01fc)

	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{S: 0x01ff, PC: 0x3456, K: 0x12, dFlag: true, memory: mem},
			expected: CPU{S: 0x01fb, iFlag: true, PC: 0x0000, memory: mem2},
		},
	}

	for i, tc := range testCases {
		tc.value.brk()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
