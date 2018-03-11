package cpu

import (
	"testing"
)

func TestBit(t *testing.T) {

	testCases := []struct {
		expected       CPU
		value          *CPU
		dataHi, dataLo uint8
		immediate      bool
	}{
		{
			value:    &CPU{C: 0x0043, DBR: 0x12, mFlag: true},
			expected: CPU{C: 0x0043, DBR: 0x12, mFlag: true, nFlag: true, zFlag: true},
			dataHi:   0x00, dataLo: 0x9c,
		},
		{
			value:    &CPU{C: 0xabff, nFlag: true, vFlag: true, zFlag: true},
			expected: CPU{C: 0xabff, nFlag: true, vFlag: true, zFlag: false},
			dataHi:   0x00, dataLo: 0x06,
			immediate: true,
		},
	}

	for i, tc := range testCases {
		tc.value.bit(tc.dataLo, tc.dataHi, tc.immediate)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
