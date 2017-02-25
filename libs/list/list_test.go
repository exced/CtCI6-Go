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
		head: n0l1,
		tail: n0l1,
		len:  1,
	}
	n0l2 := &Node{
		Value: 0,
	}
	n1l2 := &Node{
		prev:  n0l2,
		Value: 1,
	}
	n0l2.next = n1l2
	l2 := &List{
		head: n0l2,
		tail: n1l2,
		len:  2,
	}
	n0l3 := &Node{
		Value: 0,
	}
	n1l3 := &Node{
		prev:  n0l3,
		Value: 1,
	}
	n2l3 := &Node{
		prev:  n1l3,
		Value: 2,
	}
	n0l3.next = n1l3
	n1l3.next = n2l3
	l3 := &List{
		head: n0l3,
		tail: n2l3,
		len:  3,
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
