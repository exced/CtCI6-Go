package chapter2

import (
	"github.com/exced/CtCI6-Go/libs/list"
)

func removeDupsWithBuffer(l *list.List) *list.List {
	buffer := make(map[int]bool)
	n := l.Head
	for n != nil {
		if buffer[n.Value] {
			if n.Next != nil {
				n.Next.Prev = n.Prev
			}
			if n.Prev != nil {
				n.Prev.Next = n.Next
			}
			l.Len--
		} else {
			buffer[n.Value] = true
		}
		l.Tail = n
		n = n.Next
	}
	return l
}
