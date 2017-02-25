package chapter1

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	cases := []struct {
		in1  [][]int
		want [][]int
	}{
		{[][]int{}, [][]int{}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{0, 2}, {3, 4}}, [][]int{{0, 0}, {0, 4}}},
		{[][]int{{0, 2}, {3, 4}, {5, 6}}, [][]int{{0, 0}, {0, 4}, {0, 6}}},
	}
	for _, c := range cases {
		zeroMatrix(c.in1)
		if !reflect.DeepEqual(c.in1, c.want) {
			t.Errorf("zeroMatrix(%v) == %v, want %v", c.in1, c.in1, c.want)
		}
	}
}
