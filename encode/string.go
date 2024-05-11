package encode

import (
	"fmt"
	"github.com/efexplose/bits/needed"
)

// Aligned to the left.
func StringL(bits *[]bool, decludeCount int, s string) (int, error) {
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

// Aligned to the right.
func StringR(bits *[]bool, decludeCount int, s string) (int, error) {
	space := needed.String(s)

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