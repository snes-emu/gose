package cpu

import "testing"

func TestAnd(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
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
		tc.value.and(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestEor(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{C: 0x0f06},
			expected: CPU{C: 0xfe05, nFlag: true},
			dataHi:   0xf1, dataLo: 0x03,
		},
		{
			value:    &CPU{C: 0xffff, mFlag: true},
			expected: CPU{C: 0xff00, mFlag: true, zFlag: true},
			dataHi:   0x00, dataLo: 0xff,
		},
		{
			value:    &CPU{C: 0xaac4, mFlag: true},
			expected: CPU{C: 0xaa06, mFlag: true},
			dataHi:   0x00, dataLo: 0xc2,
		},
	}

	for i, tc := range testCases {
		tc.value.eor(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestOra(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{C: 0xf006},
			expected: CPU{C: 0xf107, nFlag: true},
			dataHi:   0xf1, dataLo: 0x03,
		},
		{
			value:    &CPU{C: 0x0000, mFlag: true},
			expected: CPU{C: 0x00ff, mFlag: true, nFlag: true},
			dataHi:   0x00, dataLo: 0xff,
		},
		{
			value:    &CPU{C: 0x0000, mFlag: true},
			expected: CPU{C: 0x0000, mFlag: true, zFlag: true},
			dataHi:   0x00, dataLo: 0x00,
		},
	}

	for i, tc := range testCases {
		tc.value.ora(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
