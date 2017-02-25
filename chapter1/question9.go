package chapter1

import (
	"strings"
)

func isStringRotation(str1, str2 string) bool {
	return strings.Contains(str2+str2, str1)
}
