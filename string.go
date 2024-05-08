package bits

import (
	"fmt"
)

func Println(bits []bool) {
	for _, b := range bits {
		if b {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
}

func EncodeStringL(bits *[]bool, decludeCount int, s string) (int, error) {
	space := len(s) * 8
	if len(*bits) - decludeCount < space {
		return 0, fmt.Errorf("String %s does not fit in %d bits", s, len(*bits) - decludeCount)
	}

	for i, b := range s {
		for j := 0; j < 8; j++ {
			if b&(1<<uint(7-j)) > 0 {
				(*bits)[decludeCount+i*8+j] = true
			} else {
				(*bits)[decludeCount+i*8+j] = false
			}
		}
	}

	return space, nil
}

func EncodeStringR(bits *[]bool, decludeCount int, s string) (int, error) {
	space := BitsNeededString(s)

	end := len(*bits) - decludeCount
	if end < space {
		return 0, fmt.Errorf("String %s does not fit in %d bits", s, end)
	}

	sLen := len(s)
	for i, b := range s {
		for j := 0; j < 8; j++ {
			if b&(1<<uint(7-j)) > 0 {
				(*bits)[end-sLen*8+i*8+j] = true
			} else {
				(*bits)[end-sLen*8+i*8+j] = false
			}
		}
	}
	

	return space, nil
}

func DecodeString(bits []bool) string {
	var result string
	for i := 0; i < len(bits); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			if bits[i+j] {
				b |= 1 << uint(7-j)
			}
		}
		result += string(b)
	}
	return result
}