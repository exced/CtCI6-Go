package chapter1

func isPalindromePermutation(str string) bool {
	var nbOdd int
	charFrequency := make(map[rune]int, len(str))
	for _, r := range str {
		charFrequency[r]++
		if charFrequency[r]&1 == 1 {
			nbOdd++
		} else {
			nbOdd--
		}
	}
	return nbOdd <= 1
}
