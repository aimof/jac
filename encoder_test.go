package jac

import "testing"

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
