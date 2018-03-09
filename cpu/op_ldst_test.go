package cpu

import (
	"testing"
)

func TestLda(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{},
			expected: CPU{C: 0xcdab, nFlag: true},
			dataHi:   0xcd, dataLo: 0xab,
		},
	}

	for i, tc := range testCases {
		tc.value.lda(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
