package cpu

import "testing"

func TestCmp(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
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

func TestCpx(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{X: 0x1234},
			expected: CPU{X: 0x1234, zFlag: true, cFlag: true},
			dataHi:   0x12, dataLo: 0x34,
		},
		{
			value:    &CPU{X: 0x0004, xFlag: true},
			expected: CPU{X: 0x0004, xFlag: true, zFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
		{
			value:    &CPU{X: 0x0003, xFlag: true},
			expected: CPU{X: 0x0003, xFlag: true, nFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
	}

	for i, tc := range testCases {
		tc.value.cpx(tc.dataHi, tc.dataLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestCpy(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{Y: 0x2567},
			expected: CPU{Y: 0x2567, zFlag: true, cFlag: true},
			dataHi:   0x25, dataLo: 0x67,
		},
		{
			value:    &CPU{Y: 0x0019, xFlag: true},
			expected: CPU{Y: 0x0019, xFlag: true, zFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x19,
		},
		{
			value:    &CPU{Y: 0x00da, xFlag: true},
			expected: CPU{Y: 0x00da, xFlag: true},
			dataHi:   0x00, dataLo: 0xd9,
		},
	}

	for i, tc := range testCases {
		tc.value.cpy(tc.dataHi, tc.dataLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
