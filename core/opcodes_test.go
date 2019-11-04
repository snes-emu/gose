package core

import (
	"testing"
)

func newTestMemory() *Memory {
	mem := newMemory()
	mem.initMmap()

	return mem
}

func TestBit(t *testing.T) {

	testCases := []struct {
		expected       CPU
		value          *CPU
		dataHi, dataLo uint8
		immediate      bool
	}{
		{
			value:    &CPU{C: 0x0043, DBR: 0x12, mFlag: true},
			expected: CPU{C: 0x0043, DBR: 0x12, mFlag: true, nFlag: true, zFlag: true},
			dataHi:   0x00, dataLo: 0x9c,
		},
		{
			value:    &CPU{C: 0xabff, nFlag: true, vFlag: true, zFlag: true},
			expected: CPU{C: 0xabff, nFlag: true, vFlag: true, zFlag: false},
			dataHi:   0x00, dataLo: 0x06,
			immediate: true,
		},
	}

	for i, tc := range testCases {
		tc.value.bit(tc.dataLo, tc.dataHi, tc.immediate)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

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
		tc.value.adc(tc.dataLo, tc.dataHi)

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
		tc.value.sbc(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestCmp(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{C: 0x1234},
			expected: CPU{C: 0x1234, zFlag: true, cFlag: true},
			dataHi:   0x12, dataLo: 0x34,
		},
		{
			value:    &CPU{C: 0x1104, mFlag: true},
			expected: CPU{C: 0x1104, mFlag: true, zFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
		{
			value:    &CPU{C: 0x1103, mFlag: true},
			expected: CPU{C: 0x1103, mFlag: true, nFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
	}

	for i, tc := range testCases {
		tc.value.cmp(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestCpx(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{X: 0x1234},
			expected: CPU{X: 0x1234, zFlag: true, cFlag: true},
			dataHi:   0x12, dataLo: 0x34,
		},
		{
			value:    &CPU{X: 0x0004, xFlag: true},
			expected: CPU{X: 0x0004, xFlag: true, zFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
		{
			value:    &CPU{X: 0x0003, xFlag: true},
			expected: CPU{X: 0x0003, xFlag: true, nFlag: true},
			dataHi:   0x00, dataLo: 0x04,
		},
	}

	for i, tc := range testCases {
		tc.value.cpx(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestCpy(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{Y: 0x2567},
			expected: CPU{Y: 0x2567, zFlag: true, cFlag: true},
			dataHi:   0x25, dataLo: 0x67,
		},
		{
			value:    &CPU{Y: 0x0019, xFlag: true},
			expected: CPU{Y: 0x0019, xFlag: true, zFlag: true, cFlag: true},
			dataHi:   0x00, dataLo: 0x19,
		},
		{
			value:    &CPU{Y: 0x00da, xFlag: true},
			expected: CPU{Y: 0x00da, xFlag: true},
			dataHi:   0x00, dataLo: 0xd9,
		},
	}

	for i, tc := range testCases {
		tc.value.cpy(tc.dataLo, tc.dataHi)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestInx(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{X: 0x7FFF},
			expected: CPU{X: 0x8000, nFlag: true, PC: 1},
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

func TestXba(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
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

func TestXce(t *testing.T) {

	testCases := []struct {
		value          *CPU
		expected       CPU
		dataHi, dataLo uint8
	}{
		{
			value:    &CPU{eFlag: true},
			expected: CPU{cFlag: true, PC: 1},
		},
	}

	for i, tc := range testCases {
		tc.value.opFB()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestSep(t *testing.T) {
	memory := newTestMemory()
	memory.SetByteBank(0x21, 0x00, 0x0001)
	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{memory: memory},
			expected: CPU{mFlag: true, cFlag: true, memory: memory, PC: 2},
		},
	}

	for i, tc := range testCases {
		tc.value.sep()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestClc(t *testing.T) {
	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{cFlag: true},
			expected: CPU{PC: 0x01},
		},
	}

	for i, tc := range testCases {
		tc.value.clc()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestJsr(t *testing.T) {
	mem := newTestMemory()
	mem2 := newTestMemory()
	mem2.SetByteBank(0x34, 0x00, 0x01ff)
	mem2.SetByteBank(0x58, 0x00, 0x01fe)
	testCases := []struct {
		value    *CPU
		expected CPU
		addr     uint16
	}{
		{
			value:    &CPU{S: 0x01ff, DBR: 0x12, PC: 0x3456, memory: mem},
			expected: CPU{S: 0x01fd, DBR: 0x12, PC: 0xabcd, memory: mem2},
			addr:     0xabcd,
		},
	}

	for i, tc := range testCases {
		tc.value.jsr(tc.addr)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

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

func TestRts(t *testing.T) {
	mem := newTestMemory()
	mem.SetByteBank(0x56, 0x00, 0x01fe)
	mem.SetByteBank(0x34, 0x00, 0x01ff)
	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{S: 0x01fd, DBR: 0x12, memory: mem},
			expected: CPU{S: 0x01ff, DBR: 0x12, PC: 0x3457, memory: mem},
		},
	}

	for i, tc := range testCases {
		tc.value.rts()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestRti(t *testing.T) {
	mem := newTestMemory()
	mem.SetByteBank(0x08, 0x00, 0x01fc)
	mem.SetByteBank(0x12, 0x00, 0x01fd)
	mem.SetByteBank(0x34, 0x00, 0x01fe)
	mem.SetByteBank(0x56, 0x00, 0x01ff)
	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{S: 0x01fb, memory: mem},
			expected: CPU{S: 0x01ff, K: 0x56, PC: 0x3412, dFlag: true, memory: mem},
		},
	}

	for i, tc := range testCases {
		tc.value.rti()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestAsl(t *testing.T) {

	mem := newTestMemory()
	mem.SetByteBank(0x8f, 0x7E, 0xabcd)

	mem2 := newTestMemory()
	mem2.SetByteBank(0x1e, 0x7E, 0xabcd)

	testCases := []struct {
		value        *CPU
		expected     CPU
		haddr, laddr uint32
		isAcc        bool
	}{
		{
			value:    &CPU{C: 0x0c, DBR: 0x7E, mFlag: true, memory: mem},
			expected: CPU{C: 0x0c, DBR: 0x7E, cFlag: true, mFlag: true, memory: mem2},
			haddr:    0x0, laddr: 0x7eabcd,
		},
	}

	for i, tc := range testCases {
		tc.value.asl(tc.laddr, tc.haddr, tc.isAcc)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestTrb(t *testing.T) {

	mem := newTestMemory()
	mem.SetByteBank(0x9c, 0x7e, 0xabcd)

	mem2 := newTestMemory()
	mem2.SetByteBank(0x90, 0x7e, 0xabcd)

	testCases := []struct {
		value        *CPU
		expected     CPU
		haddr, laddr uint32
	}{
		{
			value:    &CPU{C: 0x0c, DBR: 0x12, mFlag: true, memory: mem},
			expected: CPU{C: 0x0c, DBR: 0x12, mFlag: true, memory: mem2},
			haddr:    0x0, laddr: 0x7eabcd,
		},
	}

	for i, tc := range testCases {
		tc.value.trb(tc.laddr, tc.haddr)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

func TestTsb(t *testing.T) {

	mem := newTestMemory()
	mem.SetByteBank(0x9c, 0x12, 0xabcd)

	mem2 := newTestMemory()
	mem2.SetByteBank(0xdf, 0x12, 0xabcd)

	testCases := []struct {
		value        *CPU
		expected     CPU
		haddr, laddr uint32
	}{
		{
			value:    &CPU{C: 0x0043, DBR: 0x12, mFlag: true, memory: mem},
			expected: CPU{C: 0x0043, DBR: 0x12, mFlag: true, zFlag: true, memory: mem2},
			haddr:    0x0, laddr: 0x12abcd,
		},
	}

	for i, tc := range testCases {
		tc.value.tsb(tc.laddr, tc.haddr)

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}

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

func TestBrk(t *testing.T) {
	mem := newTestMemory()
	mem.initMmap()
	mem.SetByteBank(0xab, 0x00, 0xffe6)
	mem.SetByteBank(0xcd, 0x00, 0xffe7)

	mem2 := newTestMemory()
	mem2.initMmap()
	mem2.SetByteBank(0xab, 0x00, 0xffe6)
	mem2.SetByteBank(0xcd, 0x00, 0xffe7)
	mem2.SetByteBank(0x12, 0x00, 0x01ff)
	mem2.SetByteBank(0x34, 0x00, 0x01fe)
	mem2.SetByteBank(0x58, 0x00, 0x01fd)
	mem2.SetByteBank(0x08, 0x00, 0x01fc)

	testCases := []struct {
		value    *CPU
		expected CPU
	}{
		{
			value:    &CPU{S: 0x01ff, PC: 0x3456, K: 0x12, dFlag: true, memory: mem},
			expected: CPU{S: 0x01fb, iFlag: true, PC: 0x0000, memory: mem2},
		},
	}

	for i, tc := range testCases {
		tc.value.brk()

		err := tc.value.compare(tc.expected)

		if err != nil {
			t.Errorf("Test %v failed: \n%v", i, err)
		}
	}
}
