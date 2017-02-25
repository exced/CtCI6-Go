package chapter1

func urlify(str []rune, length int) {
	var nbSpaces int
	for i := 0; i < length; i++ {
		if str[i] == rune(' ') {
			nbSpaces++
		}
	}
	gLength := length + 2*nbSpaces
	for i := length - 1; i >= 0; i-- {
		if str[i] == rune(' ') {
			str[gLength-1] = rune('0')
			str[gLength-2] = rune('2')
			str[gLength-3] = rune('%')
			gLength -= 3
		} else {
			str[gLength-1] = str[i]
			gLength--
		}
	}
}
