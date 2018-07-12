package jac

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToUnicodes(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected []uint64
		next     []uint8
	}{
		{[]byte{0x00, 0x00, 0x00}, []uint64{0x3044, 0x3044}, nil},
		{[]byte{0x80, 0x08, 0x00}, []uint64{0x3066, 0x3066}, nil},
		{[]byte{0x93, 0x59, 0x53}, []uint64{0x35, 0x53}, nil},
		{[]byte{0xC2, 0x35, 0x46}, []uint64{0x23546}, nil},
		{[]byte{0xD0, 0x00, 0x04}, nil, []uint8{0xD0, 0x00, 0x04}},
	}
	for _, tt := range testCases {
		actual, next := toUnicodes(tt.input)
		if !reflect.DeepEqual(next, tt.next) {
			t.Errorf("input: %v\nexpected: %v\nactual: %v", tt.input, tt.next, next)
		}
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("input: %v\nexpected: %v\nactual: %v", tt.input, tt.expected, actual)
		}
	}
}

func TestBoth(t *testing.T) {
	testCases := []struct {
		chara []rune
	}{
		{[]rune("あ")},
		{[]rune("こころ")},
		{[]rune("Hello, world!")},
		{[]rune("亜")},
	}

	for _, tt := range testCases {
		b, err := Encode(tt.chara)
		if err != nil {
			t.Error(err)
		}
		unicodes := Decode([]byte(b))

		var r = make([]rune, 0, len(unicodes))
		for _, u := range unicodes {
			tmp := fmt.Sprintf("%c", u)
			r = append(r, []rune(tmp)[0])
		}
		if !reflect.DeepEqual(r, tt.chara) {
			t.Errorf("expected: %v\nactual: %v", tt.chara, r)
		}
	}
}
