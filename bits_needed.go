package bits

import (
	"math"
)

func BitsNeededInt(n int) int {
	return int(math.Log2(float64(n))) + 1
}

func BitsNeededString(s string) int {
	println(len(s) * 8)
	return len(s) * 8
}