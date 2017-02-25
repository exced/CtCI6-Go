package chapter1

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"abcd", true},
		{"aa", false},
		{"aba", false},
	}
	for _, c := range cases {
		got := isUnique(c.in)
		if got != c.want {
			t.Errorf("isUnique(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
