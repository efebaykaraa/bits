package decode

func String(bits *[]bool, size int, decludeCount int) string {
	var result string
	for i := decludeCount; i < len(*bits); i += 8 {
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