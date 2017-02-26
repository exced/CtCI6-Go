package chapter2

import (
	"reflect"
	"testing"

	"github.com/exced/CtCI6-Go/libs/list"
)

func TestKthToEnd(t *testing.T) {
	cases := []struct {
		in1  *list.List
		in2  int
		want *list.List
	}{
		{list.New([]int{}), 1, nil},
		{list.New([]int{0, 1, 2}), 0, list.New([]int{0, 1, 2})},
		{list.New([]int{0, 1, 2}), 1, list.New([]int{1, 2})},
		{list.New([]int{0, 1, 2}), 2, list.New([]int{2})},
		{list.New([]int{0, 1, 2, 3, 4, 5}), 2, list.New([]int{2, 3, 4, 5})},
	}
	for _, c := range cases {
		got := kthToEnd(c.in1, c.in2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("kthToEnd(%v, %v) = %v, want %v", c.in1, c.in2, got, c.want)
		}
	}
}
