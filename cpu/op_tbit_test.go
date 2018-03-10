package cpu

import (
	"testing"

	"github.com/snes-emu/gose/memory"
)

func TestTrb(t *testing.T) {

	mem := memory.New()
	mem.SetByteBank(0x9c, 0x7e, 0xabcd)

	mem2 := memory.New()
	mem2.SetByteBank(0x90, 0x7e, 0xabcd)

	testCases := []struct {
		value        *CPU
		expected     CPU
		haddr, laddr uint32
	}{
		{
			value:    &CPU{C: 0x0c, DBR: 0x12, mFlag: true, memory: mem},
			expected: CPU{C: 0x0c, DBR: 0x12, mFlag: true, memory: mem2},
			haddr:    0x0, laddr: 0x7eabcd,
		},
	}

	for i, tc := range testCases {
		tc.value.trb(tc.laddr, tc.haddr)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestTsb(t *testing.T) {

	mem := memory.New()
	mem.SetByteBank(0x9c, 0x12, 0xabcd)

	mem2 := memory.New()
	mem2.SetByteBank(0xdf, 0x12, 0xabcd)

	testCases := []struct {
		value        *CPU
		expected     CPU
		haddr, laddr uint32
	}{
		{
			value:    &CPU{C: 0x0043, DBR: 0x12, mFlag: true, memory: mem},
			expected: CPU{C: 0x0043, DBR: 0x12, mFlag: true, zFlag: true, memory: mem2},
			haddr:    0x0, laddr: 0x12abcd,
		},
	}

	for i, tc := range testCases {
		tc.value.tsb(tc.laddr, tc.haddr)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
