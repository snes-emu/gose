package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestAsl(t *testing.T) {

	mem := memory.New()
	mem.SetByteBank(0x8f, 0x7E, 0xabcd)

	mem2 := memory.New()
	mem2.SetByteBank(0x1e, 0x7E, 0xabcd)

	testCases := []struct {
		value        *CPU
		expected     CPU
		haddr, laddr uint32
		isAcc        bool
	}{
		{
			value:    &CPU{C: 0x0c, DBR: 0x7E, mFlag: true, memory: mem},
			expected: CPU{C: 0x0c, DBR: 0x7E, cFlag: true, mFlag: true, memory: mem2},
			haddr:    0x0, laddr: 0x7eabcd,
		},
	}

	for i, tc := range testCases {
		tc.value.asl(tc.laddr, tc.haddr, tc.isAcc)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
