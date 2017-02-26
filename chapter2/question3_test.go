package chapter2

import (
	"reflect"
	"testing"

	"github.com/exced/CtCI6-Go/libs/list"
)

func TestRemoveNode(t *testing.T) {
	l1 := list.New([]int{0, 1, 2, 3, 4, 5})
	cases := []struct {
		in1  *list.Node
		in2  int
		want error
	}{
		{&list.Node{}, 1, nil},
	}
	for _, c := range cases {
		got := kthToEnd(c.in1, c.in2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("kthToEnd(%v, %v) = %v, want %v", c.in1, c.in2, got.Head, c.want.Head)
		}
	}
}
