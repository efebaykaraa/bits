package needed

import (
	"math"
)

func Int(n int) int {
	return int(math.Log2(float64(n))) + 1
}

func Octets(n int) int {
	return n * 8
}

func Decimals(n int) int {
	return int(math.Log10(float64(n))) + 1
}