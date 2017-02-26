// Package list gives us the tools to work with doubly linked lists.
// Each fields of the struct are exported to allow programmers to work with the
// interview exercises and show he can creates this package himself.
package list

import (
	"errors"
)

// Node stores the data and the links between previous and next.
type Node struct {
	Next, Prev *Node
	Value      int
}

// List stores the head, which is the first node, and the tail, which is the last.
type List struct {
	Head, Tail *Node
	Len        int
}

func (l *List) init(vals []int) *List {
	if len(vals) < 1 {
		return l
	}
	curr := &Node{
		Next:  nil,
		Prev:  nil,
		Value: vals[0],
	}
	l.Head = curr
	for i := 1; i < len(vals); i++ {
		curr.Next = &Node{
			Next:  nil,
			Prev:  curr,
			Value: vals[i],
		}
		curr = curr.Next
	}
	l.Tail = curr
	l.Len = len(vals)
	return l
}

// New initializes the list by pushing the input array
func New(vals []int) *List {
	return new(List).init(vals)
}

// NodeAt returns the node at index position : [0:l.Len)
func (l *List) NodeAt(index int) (*Node, error) {
	if index < 0 || index > l.Len-1 {
		return nil, errors.New("index not accessible in this list")
	}
	n := l.Head
	if index <= (l.Len - index) { // Head to Tail
		for i := 0; i < index; i++ {
			n = n.Next
		}
	} else { // Tail to Head
		n = l.Tail
		for i := 0; i < l.Len-index-1; i++ {
			n = n.Prev
		}
	}
	return n, nil
}

// ToArray turns the list into an int array of values between head and tail
func (l *List) ToArray() []int {
	vals := make([]int, l.Len)
	n := l.Head
	i := 0
	for n != nil {
		vals[i] = n.Value
		i++
		n = n.Next
	}
	return vals
}
