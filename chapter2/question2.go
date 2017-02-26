package chapter2

import (
	"github.com/exced/CtCI6-Go/libs/list"
)

func kthToEnd(l *list.List, k int) *list.List {
	if k < 0 || k > l.Len-1 {
		return nil
	}
	n := l.Head
	if k <= (l.Len - k) { // Head to Tail
		for i := 0; i < k; i++ {
			n = n.Next
		}
	} else { // Tail to Head
		n = l.Tail
		for i := 0; i < l.Len-k-1; i++ {
			n = n.Prev
		}
	}
	newTail := l.Tail
	if k == l.Len-1 {
		newTail.Prev = nil
	}
	return &list.List{
		Head: &list.Node{
			Next:  n.Next,
			Prev:  nil,
			Value: n.Value,
		},
		Tail: newTail,
		Len:  l.Len - k,
	}
}
