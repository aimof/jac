package jac

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		input    []rune
		expected []uint8
	}{
		{[]rune("aa"), []uint8{0x96, 0x96, 0x96}},
	}
	for _, tt := range testCases {
		r, err := Encode(tt.input)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(r, tt.expected) {
			t.Errorf("input: %X, expected: %X, actual: %X", tt.input, tt.expected, r)
		}
	}
}

func TestCombine(t *testing.T) {
	testCases := []struct {
		input    []uint64
		expected []uint8
	}{
		{[]uint64{0x000, 0x000}, []uint8{0x00, 0x00, 0x00}},
		{[]uint64{0xABC, 0x000}, []uint8{0xAB, 0xC0, 0x00}},
		{[]uint64{0xFFF, 0xFFF}, []uint8{0xFF, 0xFF, 0xFF}},
		{[]uint64{0xAB0, 0x0F9}, []uint8{0xAB, 0x00, 0xF9}},
	}

	for _, tt := range testCases {
		b := combine(tt.input[0], tt.input[1])
		if !reflect.DeepEqual(b, tt.expected) {
			t.Errorf("input: %X, expected: %X, actual: %X", tt.input, tt.expected, b)
		}
	}
}

func TestIdentify(t *testing.T) {
	testcases := []struct {
		input    rune
		expected uint64
	}{
		{[]rune("あ")[0], 0xA42},
		{[]rune("「")[0], 0xA0C},
		{[]rune("￥")[0], 0xBE5},
		{[]rune("一")[0], 0x003},
		{[]rune("A")[0], 0x941},
		{[]rune("ੳ")[0], 0xC00A73},
		{[]rune("吅")[0], 0xC05405},
	}

	for _, tt := range testcases {
		n, err := identify(tt.input)
		if err != nil {
			t.Error()
			continue
		}
		if n != tt.expected {
			t.Errorf("input: %U, expected: %X, actual: %X", tt.input, tt.expected, n)
		}
	}
}
