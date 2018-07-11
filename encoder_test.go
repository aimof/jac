package jac

import "testing"

func TestIdentify(t *testing.T) {
	testcases := []struct {
		input    rune
		expected uint64
	}{
		{[]rune("あ")[0], 0x342},
		{[]rune("「")[0], 0x30C},
		{[]rune("￥")[0], 0x4E5},
		{[]rune("一")[0], 0x701},
		{[]rune("A")[0], 0x141},
		{[]rune("ੳ")[0], 0xEE0A73},
		{[]rune("吅")[0], 0xEE5405},
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
