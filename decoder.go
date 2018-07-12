package jac

import (
	"log"
)

func Decode(n []uint8) []uint64 {
	if len(n)%3 != 0 {
		log.Fatalln("Decode(): n must be 3*n")
	}

	unicodes := make([]uint64, 0, len(n)/3*2)
	tmpBigUnicodes := make([]uint8, 0, 3)
	for i := 0; i < len(n)/3; i++ {
		target := n[3*i : 3*(i+1)]
		if len(tmpBigUnicodes) == 3 {
			unicodes = append(unicodes, ((uint64(tmpBigUnicodes[0])&0x0F)<<20)|(uint64(tmpBigUnicodes[1])<<16)|(uint64(tmpBigUnicodes[2])<<12)|(uint64(target[0])<<8)|(uint64(target[1])<<4)|uint64(target[2]))
			tmpBigUnicodes = nil
		} else {
			var u []uint64
			u, tmpBigUnicodes = toUnicodes(target)
			for _, n := range u {
				if n < 0x10000000 {
					unicodes = append(unicodes, n)
				}
			}
		}
	}
	return unicodes
}

func toUnicodes(u []uint8) (unicodes []uint64, tmp []uint8) {
	unicodes = make([]uint64, 0, len(u)/3*2)
	if len(u) < 3 {
		log.Fatalln("toUnicodes(): u length must be 3")
	}

	switch (uint8(u[0]) & 0xF0) >> 4 {
	case 0xD:
		return nil, []uint8(u)
	case 0xC:
		return []uint64{(uint64(u[0])&0x0F)<<16 | (uint64(u[1]) << 8) | uint64(u[2])}, nil
	}

	chara0 := ((uint64(u[0])) << 4) | ((uint64(u[1]) & 0xF0) >> 4)
	chara1 := ((uint64(u[1]) & 0x0F) << 8) | uint64(u[2])

	unicodes = []uint64{toUnicodeFrom12bits(chara0), toUnicodeFrom12bits(chara1)}

	return unicodes, nil
}

func toUnicodeFrom12bits(c uint64) uint64 {
	if u, ok := mapJocToU[c]; ok {
		return u
	}
	switch (c & 0xF00) >> 8 {
	case 0x5:
		return 0x10000000
	case 0x6:
		return 0x10000000
	case 0x9:
		return c & 0xFF
	case 0xA:
		return (c & 0x00FF) | 0x3000
	case 0xB:
		return (c & 0x00FF) | 0xFF00
	case 0xC:
		return c & 0x0FFFFF
	case 0xD:
		return 0x10000000
	default:
		return 0x10000000
	}
	return 0x1000000
}
