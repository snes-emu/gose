package utils

import (
	"fmt"
	"testing"
)

func TestJoinUint16(t *testing.T) {

	testCases := []struct {
		expected uint16
		HH, LL   uint8
	}{
		{
			expected: 0xabcd,
			HH:       0xab, LL: 0xcd,
		},
		{
			expected: 0x0034,
			HH:       0x00, LL: 0x34,
		},
		{
			expected: 0xdf00,
			HH:       0xdf, LL: 0x00,
		},
	}

	for i, tc := range testCases {
		value := JoinUint16(tc.LL, tc.HH)

		if value != tc.expected {
			t.Errorf("Test %v failed, got %v, expected %v", i, value, tc.expected)
		}
	}
}

func TestJoinUint32(t *testing.T) {

	testCases := []struct {
		expected   uint32
		HH, MM, LL uint8
	}{
		{
			expected: 0xabcdef,
			HH:       0xab, MM: 0xcd, LL: 0xef,
		},
		{
			expected: 0x000001,
			HH:       0x00, MM: 0x00, LL: 0x01,
		},
		{
			expected: 0x2567ea,
			HH:       0x25, MM: 0x67, LL: 0xea,
		},
	}

	for i, tc := range testCases {
		value := JoinUint32(tc.LL, tc.MM, tc.HH)
		fmt.Println(value, tc.HH, tc.MM, tc.LL)

		if value != tc.expected {
			t.Errorf("Test %v failed, got %v, expected %v", i, value, tc.expected)
		}
	}
}

func TestSplitUint16(t *testing.T) {

	testCases := []struct {
		data   uint16
		HH, LL uint8
	}{
		{
			data: 0xabcd,
			HH:   0xab, LL: 0xcd,
		},
		{
			data: 0x0034,
			HH:   0x00, LL: 0x34,
		},
		{
			data: 0xdf00,
			HH:   0xdf, LL: 0x00,
		},
	}

	for i, tc := range testCases {
		ll, hh := SplitUint16(tc.data)

		if hh != tc.HH || ll != tc.LL {
			t.Errorf("Test %v failed, got (%v, %v), expected (%v, %v)", i, hh, ll, tc.HH, tc.LL)
		}
	}
}

func TestSplitUint32(t *testing.T) {

	testCases := []struct {
		data       uint32
		HH, MM, LL uint8
	}{
		{
			data: 0xabcdef,
			HH:   0xab, MM: 0xcd, LL: 0xef,
		},
		{
			data: 0x000001,
			HH:   0x00, MM: 0x00, LL: 0x01,
		},
		{
			data: 0x2567ea,
			HH:   0x25, MM: 0x67, LL: 0xea,
		},
	}

	for i, tc := range testCases {
		ll, mm, hh := SplitUint32(tc.data)

		if hh != tc.HH || mm != tc.MM || ll != tc.LL {
			t.Errorf("Test %v failed, got (%v, %v, %v), expected (%v, %v, %v)", i, hh, mm, ll, tc.HH, tc.MM, tc.LL)
		}
	}
}

func TestLowByte(t *testing.T) {

	testCases := []struct {
		data   uint16
		result uint8
	}{
		{
			0xabcd, 0xcd,
		},
		{
			0x0001, 0x01,
		},
		{
			0x0000, 0x00,
		},
		{
			0xffaf, 0xaf,
		},
	}

	for i, tc := range testCases {
		res := LowByte(tc.data)

		if res != tc.result {
			t.Errorf("Test n°%v failed, got (%v), expected (%v)", i, res, tc.result)
		}
	}
}

func TestHighByte(t *testing.T) {

	testCases := []struct {
		data   uint16
		result uint8
	}{
		{
			0xabcd, 0xab,
		},
		{
			0x0001, 0x00,
		},
		{
			0x0000, 0x00,
		},
		{
			0xffaf, 0xff,
		},
	}

	for i, tc := range testCases {
		res := HighByte(tc.data)

		if res != tc.result {
			t.Errorf("Test n°%v failed, got (%v), expected (%v)", i, res, tc.result)
		}
	}
}

func TestSetHighByte(t *testing.T) {

	testCases := []struct {
		data   uint16
		hh     uint8
		result uint16
	}{
		{
			0xabcd, 0x00, 0x00cd,
		},
		{
			0x0001, 0xff, 0xff01,
		},
		{
			0x0000, 0x1a, 0x1a00,
		},
		{
			0xffaf, 0xff, 0xffaf,
		},
	}

	for i, tc := range testCases {
		res := SetHighByte(tc.data, tc.hh)

		if res != tc.result {
			t.Errorf("Test n°%v failed, got (%v), expected (%v)", i, res, tc.result)
		}
	}
}

func TestSetLowByte(t *testing.T) {

	testCases := []struct {
		data   uint16
		ll     uint8
		result uint16
	}{
		{
			0xabcd, 0x00, 0xab00,
		},
		{
			0x0001, 0xff, 0x00ff,
		},
		{
			0x0000, 0x1a, 0x001a,
		},
		{
			0xffaf, 0xaf, 0xffaf,
		},
	}

	for i, tc := range testCases {
		res := SetLowByte(tc.data, tc.ll)

		if res != tc.result {
			t.Errorf("Test n°%v failed, got (%v), expected (%v)", i, res, tc.result)
		}
	}
}
