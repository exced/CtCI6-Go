package chapter1

import (
	"github.com/exced/CtCI6-Go/libs/math"
)

func oneEditAway(str1, str2 string) bool {
	if math.Abs(len(str1)-len(str2)) > 1 {
		return false
	}
	longest := str1
	shortest := str2
	sameLength := len(str1) == len(str2)
	if len(str1) < len(str2) {
		longest = str2
		shortest = str1
	}
	if len(shortest) == 0 {
		if len(longest) != 1 {
			return false
		}
	}
	curs := 0
	diff := false
	for i := range shortest {
		if shortest[i] != longest[curs] {
			if diff {
				return false
			}
			diff = true
			if !sameLength { // insert
				curs++
			}
		}
		curs++
	}
	return true
}
