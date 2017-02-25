package chapter1

import (
	"testing"
)

func TestOneEditAway(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bae", false},
	}
	for _, c := range cases {
		got := oneEditAway(c.in1, c.in2)
		if got != c.want {
			t.Errorf("oneEditAway(%q, %q) == %q, want %v", c.in1, c.in2, got, c.want)
		}
	}
}
