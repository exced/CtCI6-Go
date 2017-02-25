package chapter1

func isUnique(str string) bool {
	if len(str) > 256 {
		return false
	}
	set := make(map[rune]bool)
	for _, r := range str {
		if _, done := set[r]; done {
			return false
		}
		set[r] = true
	}
	return true
}
