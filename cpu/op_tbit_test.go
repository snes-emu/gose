package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestTrb(t *testing.T) {

	mem := memory.New()
	mem.SetByteBank(0x9c, 0x12, 0xabcd)

	mem2 := memory.New()
	mem2.SetByteBank(0x90, 0x12, 0xabcd)

	testCases := []struct {
		value          *CPU
		expected       CPU
		addrHi, addrLo uint32
	}{
		{
			value:    &CPU{C: 0x0c, DBR: 0x12, mFlag: true, memory: mem},
			expected: CPU{C: 0x0c, DBR: 0x12, mFlag: true, memory: mem2},
			addrHi:   0x0, addrLo: 0x12abcd,
		},
	}

	for i, tc := range testCases {
		tc.value.trb(tc.addrHi, tc.addrLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
