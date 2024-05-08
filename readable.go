package github.com/efexplose/bits

func Readable(bits []bool) string {
	var result string
	for _, b := range bits {
		if b {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func ReadableGroup(bits []bool, groupSize int) string {
	var result string
	for i, b := range bits {
		if i != 0 && i%groupSize == 0 {
			result += " "
		}
		if b {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}