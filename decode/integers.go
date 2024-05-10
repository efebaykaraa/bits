package decode

// Range must be specified.
func Int(bits *[]bool, start int, end int) (int, error) {
	n := 0
	for i := start; i < end; i++ {
        if (*bits)[i] {
            n |= 1 << (i - start)
        }
    }

	return n, nil
}