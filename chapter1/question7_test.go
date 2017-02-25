package chapter1

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		in1  [][]int
		in2  int
		want [][]int
	}{
		{[][]int{}, 0, [][]int{}},
		{[][]int{{1}}, 1, [][]int{{1}}},
		{[][]int{{1, 2}, {3, 4}}, 2, [][]int{{3, 1}, {4, 2}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 3, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 4,
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}},
	}
	for _, c := range cases {
		rotate(c.in1, c.in2)
		if !reflect.DeepEqual(c.in1, c.want) {
			t.Errorf("rotate(%q, %q) == %q, want %v", c.in1, c.in2, c.in1, c.want)
		}
	}
}
