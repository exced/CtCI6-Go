package chapter2

import (
	"reflect"
	"testing"

	"github.com/exced/CtCI6-Go/libs/list"
)

func TestRemoveDupsWithBuffer(t *testing.T) {
	cases := []struct {
		in1  *list.List
		want *list.List
	}{
		{list.New([]int{}), list.New([]int{})},
		{list.New([]int{0}), list.New([]int{0})},
		{list.New([]int{0, 1, 2}), list.New([]int{0, 1, 2})},
		{list.New([]int{1, 1}), list.New([]int{1})},
		{list.New([]int{1, 1, 1}), list.New([]int{1})},
		{list.New([]int{1, 1, 1, 2}), list.New([]int{1, 2})},
		{list.New([]int{1, 2, 1, 2}), list.New([]int{1, 2})},
	}
	for _, c := range cases {
		removeDupsWithBuffer(c.in1)
		if !reflect.DeepEqual(c.in1, c.want) {
			t.Errorf("RemoveDupsWithBuffer(%v), want %v", c.in1.Tail, c.want.Tail)
		}
	}
}
