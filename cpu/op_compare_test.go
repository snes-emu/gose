package cpu

import "testing"

func TestCmp(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
		operator       func(uint8, uint8)
	}{
		{
			value:    &CPU{C: 0x1234},
			expected: CPU{C: 0x1234, zFlag: true, cFlag: true},
			dataHi:   0x12, dataLo: 0x34,
		},
		{
			value:    &CPU{C: 0x1104, mFlag: true},
			expected: CPU{C: 0x1104, mFlag: true, zFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
		{
			value:    &CPU{C: 0x1103, mFlag: true},
			expected: CPU{C: 0x1103, mFlag: true, nFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
	}

	for i, tc := range testCases {
		tc.value.cmp(tc.dataHi, tc.dataLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
