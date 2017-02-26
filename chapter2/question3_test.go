package chapter2

import (
	"testing"

	"github.com/exced/CtCI6-Go/libs/list"
)

func TestRemoveNode(t *testing.T) {
	n0l1 := &list.Node{
		Value: 0,
	}
	n1l1 := &list.Node{
		Prev:  n0l1,
		Value: 1,
	}
	n2l1 := &list.Node{
		Prev:  n1l1,
		Value: 2,
	}
	n0l1.Next = n1l1
	n1l1.Next = n2l1
	if removeNode(n1l1) != nil || n0l1.Next.Value != 2 || n0l1.Next.Prev != n0l1 {
		t.Errorf("removeNode(%v) error", n1l1)
	}
}
