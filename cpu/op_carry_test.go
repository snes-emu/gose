package cpu

import "testing"

func TestAdcExample1(t *testing.T) {
	cpu := &CPU{
		C:     0x0001,
		mFlag: false,
		dFlag: false,
		cFlag: true,
	}

	cpu.adc(0x20, 0x03)

	cpu2 := CPU{
		C:     0x2005,
		nFlag: false,
		zFlag: false,
		vFlag: false,
		cFlag: false,
	}

	err := cpu.compare(cpu2)

	if err != nil {
		t.Error(err)
	}
}

func TestSbcExample1(t *testing.T) {
	cpu := &CPU{
		C:     0x0001,
		mFlag: false,
		dFlag: false,
		cFlag: true,
	}

	cpu.sbc(0x20, 0x03)

	cpu2 := CPU{
		C:     0xdffe,
		nFlag: true,
		zFlag: false,
		vFlag: false,
		cFlag: false,
	}

	err := cpu.compare(cpu2)

	if err != nil {
		t.Error(err)
	}
}

// func TestSbcExample2(t *testing.T) {
// 	cpu := &CPU{
// 		C:     0x0001,
// 		mFlag: false,
// 		dFlag: true,
// 		cFlag: true,
// 	}

// 	cpu.sbc(0x20, 0x03)

// 	cpu2 := CPU{
// 		C:     0x7998,
// 		nFlag: false,
// 		zFlag: false,
// 		cFlag: false,
// 	}

// 	err := cpu.compare(cpu2)

// 	if err != nil {
// 		t.Error(err)
// 	}
// }
