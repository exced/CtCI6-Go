package chapter1

func isPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	set := make(map[rune]int)
	for _, r := range str1 {
		set[r]++
	}
	for _, r := range str2 {
		set[r]--
		if set[r] < 0 {
			return false
		}
	}
	return true
}
