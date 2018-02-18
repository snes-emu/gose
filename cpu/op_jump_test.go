package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestJsr(t *testing.T) {
	mem := memory.New()
	mem2 := memory.New()
	mem2.SetByteBank(0x34, 0x00, 0x01ff)
	mem2.SetByteBank(0x58, 0x00, 0x01fe)
	testCases := []struct {
		value    *CPU
		expected CPU
		addr     uint16
	}{
		{
			value:    &CPU{S: 0x01ff, DBR: 0x12, PC: 0x3456, memory: mem},
			expected: CPU{S: 0x01fd, DBR: 0x12, PC: 0xabcd, memory: mem2},
			addr:     0xabcd,
		},
	}

	for i, tc := range testCases {
		tc.value.jsr(tc.addr)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
