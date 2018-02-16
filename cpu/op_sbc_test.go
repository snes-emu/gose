package cpu

import "testing"

func TestSbcExample1(t *testing.T) {
	cpu := New()

	cpu.setCRegister(0x0001)
	cpu.mFlag = false
	cpu.dFlag = false
	cpu.cFlag = true

	cpu.sbc(0x20, 0x03)

	success := cpu.getCRegister() == 0xdffe &&
		cpu.nFlag &&
		!cpu.vFlag &&
		!cpu.zFlag &&
		!cpu.cFlag

	if !success {
		t.Error("Example 1 fails on the accumulator (decimal mode off)")
	}
}

func TestSbcExample2(t *testing.T) {
	cpu := New()

	cpu.setCRegister(0x0001)
	cpu.mFlag = false
	cpu.dFlag = true
	cpu.cFlag = true

	cpu.sbc(0x20, 0x03)

	success := cpu.getCRegister() == 0x7998 &&
		!cpu.nFlag &&
		!cpu.zFlag &&
		!cpu.cFlag

	if !success {
		t.Error("Example 2 fails on the accumulator (decimal mode on)")
	}
}
