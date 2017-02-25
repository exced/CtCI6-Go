package chapter1

import (
	"testing"
)

func TestCompress(t *testing.T) {
	cases := []struct {
		in1  string
		want string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"a", "a"},
		{"aaa", "a3"},
	}
	for _, c := range cases {
		got := compress(c.in1)
		if got != c.want {
			t.Errorf("compress(%q) == %q, want %v", c.in1, got, c.want)
		}
	}
}
