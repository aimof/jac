package jac

import (
	"errors"
	"fmt"
	"strconv"
)

func Encode(runes []rune) ([]uint8, error) {
	var tmp uint64 = 0x1000000
	u := make([]uint8, 0, 3*len(runes))
	for _, r := range runes {
		n, err := identify(r)
		if err != nil {
			return nil, err
		}
		if n < 0x1000 {
			if tmp < 0x1000 {
				u = append(u, combine(tmp, n)...)
				tmp = 0x1000000
				continue
			} else if tmp > 0xFFFFFF {
				tmp = n
				continue
			}
		} else if n < 0x1000000 {
			u = append(u, uint8((n&0xFF0000)>>16))
			u = append(u, uint8((n&0x00FF00)>>8))
			u = append(u, uint8(n&0x0000FF))
			continue
		} else if n < 0x1000000000000 {
			u = append(u, uint8((n&0xFF0000000000)>>40))
			u = append(u, uint8((n&0x00FF00000000)>>32))
			u = append(u, uint8((n&0x0000FF000000)>>24))
			u = append(u, uint8((n&0x000000FF0000)>>16))
			u = append(u, uint8((n&0x00000000FF00)>>8))
			u = append(u, uint8(n&0x0000000000FF))
			continue
		} else {
			return nil, errors.New("Too big rune")
		}
	}
	if tmp != 0x1000000 {
		u = append(u, combine(tmp, 0x555)...)
	}
	return u, nil
}

func combine(a, b uint64) []uint8 {
	u := make([]uint8, 0, 3)
	a = (a << 12) | b
	u = append(u, uint8((a&0xFF0000)>>16))
	u = append(u, uint8((a&0x00FF00)>>8))
	u = append(u, uint8(a&0x0000FF))
	return u
}

/*
	0xFFF以下のuint64もしくは0x6...........を返却
*/
func identify(r rune) (uint64, error) {
	u, err := strconv.ParseUint(fmt.Sprintf("%U", r)[2:6], 16, 64)
	if err != nil {
		return 0x900, err
	}
	if v, ok := mapUToJoc[u]; ok {
		return v, nil
	}
	if u >= 0x3000 && u < 0x3100 {
		return (u & 0x00FF) | 0xA00, nil
	}
	if u >= 0xFF00 && u <= 0xFFFF {
		return (u & 0x00FF) | 0xB00, nil
	}
	switch {
	case u < 0x100:
		return u | 0x900, nil
	case u < 0x100000:
		return u | 0xC00000, nil
	default:
		return u | 0xD00000000000, nil
	}
}
