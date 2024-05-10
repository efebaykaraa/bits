package needed

import (
	"math"
)

func Int(n int) int {
	return int(math.Log2(float64(n))) + 1
}

func String(s string) int {
	return len(s) * 8
}