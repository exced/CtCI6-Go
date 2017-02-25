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
	}
	for _, c := range cases {
		got := removeDupsWithBuffer(c.in1)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RemoveDupsWithBuffer(%v) == %v, want %v", c.in1, got, c.want)
		}
	}
}
