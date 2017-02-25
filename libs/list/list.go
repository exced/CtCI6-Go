package list

type Node struct {
	next, prev *Node
	Value      int
}

type List struct {
	head, tail *Node
	len        int
}

func (l *List) init(vals []int) *List {
	if len(vals) <= 0 {
		return l
	}
	curr := &Node{
		next:  nil,
		prev:  nil,
		Value: vals[0],
	}
	l.head = curr
	for i := 1; i < len(vals); i++ {
		curr.next = &Node{
			next:  nil,
			prev:  curr,
			Value: vals[i],
		}
		curr = curr.next
	}
	l.tail = curr
	l.len = len(vals)
	return l
}

func New(vals []int) *List {
	return new(List).init(vals)
}

func (l *List) InsertValueAt(val int, node *Node) *Node {
	node.next = &Node{
		next:  nil,
		prev:  node,
		Value: val,
	}
	l.len++
	return node.next
}

func (l *List) InsertValuesAt(vals []int, node *Node) *Node {
	for i := 0; i < len(vals); i++ {
		node.next = &Node{
			next:  nil,
			prev:  node,
			Value: vals[i],
		}
		node = node.next
	}
	l.len += len(vals)
	return node
}

func (l *List) InsertAt(node *Node, newNode *Node) *Node {
	node.next = newNode
	newNode.prev = node
	l.len++
	return newNode
}

func (l *List) RemoveAt(node *Node) *Node {
	if node.next != nil {
		node.next.next.prev = node
		node.next = node.next.next
	}
	l.len--
	return node.next
}
