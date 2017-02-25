package math

import (
	"testing"
)

func TestAbs(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{-1, 1},
		{0, 0},
		{2, 2},
	}
	for _, c := range cases {
		got := Abs(c.in)
		if got != c.want {
			t.Errorf("Abs(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
