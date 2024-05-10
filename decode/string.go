package decode

// Range must be specified.
func String(bits *[]bool, start int, end int) string {
	var result string
	for i := start; i < end; i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			if (*bits)[i+j] {
				b |= 1 << uint(7-j)
			}
		}
		result += string(b)
	}
	return result
}