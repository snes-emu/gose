package cpu

import "testing"

func TestAdc(t *testing.T) {

	testCases := []struct {
		expected       CPU
		value          *CPU
		dataHi, dataLo uint8
	}{
		{
			expected: CPU{C: 0x2005},
			value:    &CPU{C: 0x0001, cFlag: true},
			dataHi:   0x20, dataLo: 0x03,
		},
		{
			expected: CPU{C: 0x0006, mFlag: true},
			value:    &CPU{C: 0x00ff, mFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x06,
		},
	}

	for i, tc := range testCases {
		tc.value.adc(tc.dataHi, tc.dataLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestSbc(t *testing.T) {

	testCases := []struct {
		expected       CPU
		value          *CPU
		dataHi, dataLo uint8
	}{
		{
			expected: CPU{C: 0x00ff, mFlag: true, nFlag: true},
			value:    &CPU{C: 0x0002, mFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x03,
		},
		{
			expected: CPU{C: 0xdffe, nFlag: true},
			value:    &CPU{C: 0x0001, cFlag: true},
			dataHi:   0x20, dataLo: 0x03,
		},
	}

	for i, tc := range testCases {
		tc.value.sbc(tc.dataHi, tc.dataLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
