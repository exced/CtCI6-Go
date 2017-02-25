package chapter1

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want bool
	}{
		{"", "", true},
		{"abcd", "dbac", true},
		{"aa", "a", false},
		{"aba", "", false},
		{"c", "c", true},
	}
	for _, c := range cases {
		got := isPermutation(c.in1, c.in2)
		if got != c.want {
			t.Errorf("isPermutation(%q, %q) == %q, want %v", c.in1, c.in2, got, c.want)
		}
	}
}
