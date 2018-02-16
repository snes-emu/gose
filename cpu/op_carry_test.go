package cpu

import "testing"

func TestAdc(t *testing.T) {

	testCases := []struct {
		expected       CPU
		value          *CPU
		dataHi, dataLo uint8
		operator       func(uint8, uint8)
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

	for _, tc := range testCases {
		tc.value.adc(tc.dataHi, tc.dataLo)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Error(err)
		}
	}
}

func TestSbcExample1(t *testing.T) {
	cpu := &CPU{
		C:     0x0001,
		cFlag: true,
	}

	cpu.sbc(0x20, 0x03)

	cpu2 := CPU{
		C:     0xdffe,
		nFlag: true,
	}

	err := cpu.compare(cpu2)

	if err != nil {
		t.Error(err)
	}
}

func TestSbcExample2(t *testing.T) {
	cpu := &CPU{
		C:     0x0002,
		mFlag: true,
		cFlag: true,
	}

	cpu.sbc(0x00, 0x03)

	cpu2 := CPU{
		C:     0x00ff,
		mFlag: true,
		nFlag: true,
	}

	err := cpu.compare(cpu2)

	if err != nil {
		t.Error(err)
	}
}
