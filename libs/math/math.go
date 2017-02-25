package math

// Abs computes the absolute value of integer
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
