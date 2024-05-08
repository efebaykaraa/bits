package github.com/efexplose/bits

func Encode(res *[]bool, byteSlice []byte) {
	for _, b := range byteSlice {
		for j := 0; j < 8; j++ {
			if b&(1<<uint(7-j)) > 0 {
				*res = append(*res, true)
			} else {
				*res = append(*res, false)
			}
		}
	}
}

func Decode(res *[]byte, bits []bool) {
	for i := 0; i < len(bits); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			if bits[i+j] {
				b |= 1 << uint(7-j)
			}
		}
		*res = append(*res, b)
	}
}