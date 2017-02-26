package list

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	l0 := new(List)
	n0l1 := &Node{
		Value: 0,
	}
	l1 := &List{
		Head: n0l1,
		Tail: n0l1,
		Len:  1,
	}
	n0l2 := &Node{
		Value: 0,
	}
	n1l2 := &Node{
		Prev:  n0l2,
		Value: 1,
	}
	n0l2.Next = n1l2
	l2 := &List{
		Head: n0l2,
		Tail: n1l2,
		Len:  2,
	}
	n0l3 := &Node{
		Value: 0,
	}
	n1l3 := &Node{
		Prev:  n0l3,
		Value: 1,
	}
	n2l3 := &Node{
		Prev:  n1l3,
		Value: 2,
	}
	n0l3.Next = n1l3
	n1l3.Next = n2l3
	l3 := &List{
		Head: n0l3,
		Tail: n2l3,
		Len:  3,
	}
	cases := []struct {
		in1  []int
		want *List
	}{
		{[]int{}, l0},
		{[]int{0}, l1},
		{[]int{0, 1}, l2},
		{[]int{0, 1, 2}, l3},
	}
	for _, c := range cases {
		got := New(c.in1)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("New(%v) == %v, want %v", c.in1, got, c.want)
		}
	}
}

func TestToArray(t *testing.T) {
	cases := []struct {
		in1  *List
		want []int
	}{
		{New([]int{}), []int{}},
		{New([]int{0}), []int{0}},
		{New([]int{0, 1, 2, 3, 4}), []int{0, 1, 2, 3, 4}},
	}
	for _, c := range cases {
		got := c.in1.ToArray()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ToArray(%v) = %v, want %v", c.in1, got, c.want)
		}
	}
}

func TestNodeAt(t *testing.T) {
	n0l1 := &Node{
		Value: 0,
	}
	n1l1 := &Node{
		Prev:  n0l1,
		Value: 1,
	}
	n2l1 := &Node{
		Prev:  n1l1,
		Value: 2,
	}
	n0l1.Next = n1l1
	n1l1.Next = n2l1
	l1 := &List{
		Head: n0l1,
		Tail: n2l1,
		Len:  3,
	}
	cases := []struct {
		in1  int
		want *Node
		err  error
	}{
		{0, n0l1, nil},
		{1, n1l1, nil},
		{2, n2l1, nil},
	}
	for _, c := range cases {
		got, gotErr := l1.NodeAt(c.in1)
		if gotErr != nil || !reflect.DeepEqual(got, c.want) {
			t.Errorf("(%v).NodeAt(%v) == %v, want %v", l1, c.in1, got, c.want)
		}
	}
}
