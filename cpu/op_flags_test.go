package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestSep(t *testing.T) {
	memory := memory.New()
	memory.SetByteBank(0x21, 0x00, 0x0001)
	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
		operator       func(uint8, uint8)
	}{
		{
			value:    &CPU{memory: memory},
			expected: CPU{mFlag: true, cFlag: true, memory: memory, PC: 2},
		},
	}

	for i, tc := range testCases {
		tc.value.opE2()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
