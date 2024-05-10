package encode

import (
	"fmt"
	"github.com/efexplose/bits/needed"
)

// Aligned to the left.
func IntL(bits *[]bool, size int, decludeCount int, n int) (int, error) {
	var isize int
	if (size <= 0) {
		isize = needed.Int(n)
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

// Aligned to the right.
func IntR(bits *[]bool, size int, decludeCount int, n int) (int, error) {
	var isize int
	if (size <= 0) {
		isize = needed.Int(n)
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