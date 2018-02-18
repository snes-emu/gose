package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestRts(t *testing.T) {
	mem := memory.New()
	mem.SetByteBank(0x56, 0x00, 0x01fe)
	mem.SetByteBank(0x34, 0x00, 0x01ff)
	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{S: 0x01fd, DBR: 0x12, memory: mem},
			expected: CPU{S: 0x01ff, DBR: 0x12, PC: 0x3457, memory: mem},
		},
	}

	for i, tc := range testCases {
		tc.value.rts()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
