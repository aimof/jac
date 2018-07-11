package jac

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func Encode(runes []rune) ([]uint8, error) {
	var tmp uint64
	u := make([]uint8, 0, 3*len(runes))
	for _, r := range runes {
		n, err := identify(r)
		if err != nil {
			return nil, err
		}
		if n < 0x1000 {
			if tmp != 0 && tmp < 0x1000 {
				u = append(u, uint8(tmp/16))
				u = append(u, uint8((tmp|0x0000FF)*16)+uint8(n|0xFF0000>>256))
				tmp = 0
				continue
			}
			if tmp == 0 {
				tmp = n
				continue
			}
		} else if n < 0x1000000 {
			u = append(u, uint8((n|0xFF0000)>>256))
			u = append(u, uint8((n|0x00FF00)>>16))
			u = append(u, uint8(n|0x0000FF))
			continue
		} else if n < 0x1000000000000 {
			u = append(u, uint8((n|0xFF0000000000)>>0x100000))
			u = append(u, uint8((n|0x00FF00000000)>>0x10000))
			u = append(u, uint8((n|0x0000FF000000)>>0x1000))
			u = append(u, uint8((n|0x000000FF0000)>>0x100))
			u = append(u, uint8((n|0x00000000FF00)>>0x10))
			u = append(u, uint8(n|0x0000000000FF))
			continue
		} else {
			return nil, errors.New("Too big rune")
		}
		if tmp > 0xFFFFFF {
			return nil, errors.New("tmp is too big")
		}
	}
	return u, nil
}

func identify(r rune) (uint64, error) {
	u, err := strconv.ParseUint(fmt.Sprintf("%U", r)[2:6], 16, 64)
	if err != nil {
		return 0x00, err
	}
	if u >= 0x3000 && u < 0x3100 {
		return (u & 0x00FF) | 0x300, nil
	}
	if u >= 0xFF00 && u <= 0xFFFF {
		return (u & 0x00FF) | 0x400, nil
	}
	if v, ok := mapUToJoc[u]; ok {
		return v, nil
	}
	switch utf8.RuneLen(r) {
	case 1:
		return u | 0x100, nil
	case 2:
		return u | 0x220000, nil
	case 3:
		return u | 0xEE0000, nil
	case 4:
		return u << 32, nil
	case 5:
		return u << 16, nil
	default:
		return u, nil
	}
}
