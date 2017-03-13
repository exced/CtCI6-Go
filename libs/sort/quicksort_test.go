package sort

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{5, 3, 1, 7, 88, 9}, []int{1, 3, 5, 7, 9, 88}},
		{[]int{-1, 4, 0, 4, 5}, []int{-1, 0, 4, 4, 5}},
	}
	for _, c := range cases {
		quickSort(c.in, 0, len(c.in)-1)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("quickSort(%v) == %v, want %v", c.in, c.in, c.want)
		}
	}
}
