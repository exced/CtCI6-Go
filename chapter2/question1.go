package chapter2

import (
	"github.com/exced/CtCI6-Go/libs/list"
)

func removeDupsWithBuffer(l *list.List) {
	buffer := make(map[int]bool)
	n := l.Head
	for n != nil {
		if buffer[n.Value] { // remove node
			n.Prev.Next = n.Next
			if n.Next != nil {
				n.Next.Prev = n.Prev
			} else {
				l.Tail = n.Prev
			}
			l.Len--
		} else {
			buffer[n.Value] = true
		}
		n = n.Next
	}
}
