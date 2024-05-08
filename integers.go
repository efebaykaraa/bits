package bits

import (
	"fmt"
)

// Encode an integer to a slice of bits.
// Aligned to the left.
func EncodeIntL(bits *[]bool, size int, decludeCount int, n int) (int, error) {
	var isize int
	if (size <= 0) {
		isize = BitsNeededInt(n)
	} else {
		isize = size
	}
	decludedLen := len(*bits) - decludeCount
	if decludedLen < isize {
		return 0, fmt.Errorf("Integer %d does not fit in %d bits", n, decludedLen)
	}

	for i := 0; i < decludedLen-1; i++ {
		if n&(1<<i) > 0 {
			(*bits)[decludeCount+i] = true
		} else {
			(*bits)[decludeCount+i] = false
		}
	}

	return isize, nil
}

// Encode an integer to a slice of bits.
// Aligned to the right.
func EncodeIntR(bits *[]bool, size int, decludeCount int, n int) (int, error) {
	var isize int
	if (size <= 0) {
		isize = BitsNeededInt(n)
	} else {
		isize = size
	}
	decludedLen := len(*bits) - decludeCount
	if decludedLen < isize {
		return 0, fmt.Errorf("Integer %d does not fit in %d bits", n, decludedLen)
	}

	end := decludedLen
	for i := end-1; i >= end-isize; i-- {
		if n&(1<<end-i) > 0 {
			(*bits)[i] = true
		} else {
			(*bits)[i] = false
		}
	}

	return isize, nil
}

// Decode an integer from a slice of bits.
// Range must be specified.
func DecodeInt(bits *[]bool, start int, end int) (int, error) {
	n := 0
	for i := start; i < end; i++ {
        if (*bits)[i] {
            n |= 1 << (i - start)
        }
    }

	return n, nil
}