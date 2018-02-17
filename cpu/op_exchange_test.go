package cpu

import "testing"

func TestXba(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
		operator       func(uint8, uint8)
	}{
		{
			value:    &CPU{C: 0x6789},
			expected: CPU{C: 0x8967, PC: 1},
		},
	}

	for i, tc := range testCases {
		tc.value.opEB()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
