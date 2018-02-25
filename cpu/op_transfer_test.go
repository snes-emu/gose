package cpu

import (
	"testing"
)

func TestTdc(t *testing.T) {
	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{D: 0x1234},
			expected: CPU{C: 0x1234, D: 0x1234},
		},
	}

	for i, tc := range testCases {
		tc.value.tdc()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
