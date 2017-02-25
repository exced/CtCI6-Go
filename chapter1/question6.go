package chapter1

import (
	"bytes"
	"strconv"
)

func compress(str string) string {
	count := 1
	var buffer bytes.Buffer
	lenStr := len(str)
	for i := 1; i < lenStr; i++ {
		if str[i-1] == str[i] {
			count++
		} else {
			cstr := []byte(strconv.Itoa(count))
			buffer.WriteByte(str[i-1])
			buffer.Write(cstr)
			count = 1
		}
	}
	cstr := []byte(strconv.Itoa(count))
	buffer.WriteByte(str[lenStr-1])
	buffer.Write(cstr)
	if buffer.Len() > lenStr {
		return str
	}
	return buffer.String()
}
