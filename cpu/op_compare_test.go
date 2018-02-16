package cpu

import "testing"

func TestCmpExample1(t *testing.T) {
	cpu := &CPU{
		C:     0x1234,
		mFlag: false,
		dFlag: false,
		cFlag: false,
	}

	cpu.cmp(0x12, 0x34)

	cpu2 := CPU{
		C:     0x1234,
		nFlag: false,
		zFlag: true,
		vFlag: false,
		cFlag: true,
	}

	err := cpu.compare(cpu2)

	if err != nil {
		t.Error(err)
	}
}
