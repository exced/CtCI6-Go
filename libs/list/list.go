// Package list gives us the tools to work with doubly linked lists.
// Each fields of the struct are exported to allow programmers to work with the
// interview exercises and show he can creates this package himself.
package list

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
