package jac

import (
	"log"
)

/*func Decode(bytes []byte) []uint64 {
	if len(bytes)%3 != 0 {
		log.Fatalln(err)
	}

	unicodes := make([]uint64, len(bytes)/3*2)
	bigUnicodes := make([]uint8, 0, 3)
	for i := 0; i < len(bytes)/3; i++ {
		tmp := bytes[i : i+3]
		target := []uint8(tmp)
		if bigUnicodes != 0x0 {
			unicodes = unicodes.append()

		}
		if target[0]&0xF0 == 0xD {

		}
	}
}
*/

func toUnicodes(bytes []byte) (unicodes []uint64, tmp []uint8) {
	unicodes = make([]uint64, 0, len(bytes)/3*2)
	if len(bytes) < 3 {
		log.Fatalln("toUnicodes(): bytes length must be 3")
	}

	switch (uint8(bytes[0]) & 0xF0) >> 4 {
	case 0xD:
		return nil, []uint8(bytes)
	case 0xC:
		return []uint64{((uint64(bytes[0]) & 0x0F) << 16) | (uint64(bytes[1]) << 8) | uint64(bytes[2])}, nil
	}

	chara0 := ((uint64(bytes[0])) << 4) | ((uint64(bytes[1]) & 0xF0) >> 4)
	chara1 := ((uint64(bytes[1]) & 0x0F) << 8) | uint64(bytes[2])

	unicodes = []uint64{toUnicodeFrom12bits(chara0), toUnicodeFrom12bits(chara1)}

	return unicodes, nil
}

func toUnicodeFrom12bits(c uint64) uint64 {
	switch (c & 0xF00) >> 8 {
	case 0x5:
		return 0x0
	case 0x6:
		return 0x0
	case 0x9:
		return c & 0xFF
	case 0xA:
		return (c & 0x00FF) | 0x3000
	case 0xB:
		return (c & 0x00FF) | 0xFF00
	case 0xC:
		return 0x00
	case 0xD:
		return 0x0
	default:
		if u, ok := mapJocToU[c]; ok {
			return u
		} else {
			return 0x00
		}
	}
	return 0x00
}
