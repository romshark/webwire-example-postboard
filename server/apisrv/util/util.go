package util

// Max returns the biggest of the given integers
func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
