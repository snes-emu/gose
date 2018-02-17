package cpu

import "testing"

func TestInx(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
		operator       func(uint8, uint8)
	}{
		{
			value:    &CPU{X: 0x7FFF},
			expected: CPU{X: 0x8000, nFlag: true},
		},
	}

	for i, tc := range testCases {
		tc.value.opE8()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
