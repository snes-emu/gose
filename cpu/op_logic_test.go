package cpu

import "testing"

func TestAnd(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
		operator       func(uint8, uint8)
	}{
		{
			value:    &CPU{C: 0xf231},
			expected: CPU{C: 0x8230, nFlag: true},
			dataHi:   0x82, dataLo: 0x34,
		},
		{
			value:    &CPU{C: 0xffff, mFlag: true},
			expected: CPU{C: 0xff00, mFlag: true, zFlag: true},
			dataHi:   0x00, dataLo: 0x00,
		},
		{
			value:    &CPU{C: 0xaa03, mFlag: true},
			expected: CPU{C: 0xaa02, mFlag: true},
			dataHi:   0x00, dataLo: 0x02,
		},
	}

	for i, tc := range testCases {
		tc.value.and(tc.dataHi, tc.dataLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
