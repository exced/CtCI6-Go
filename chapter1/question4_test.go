package chapter1

import (
	"testing"
)

func TestIsPalindromePermutation(t *testing.T) {
	cases := []struct {
		in1  string
		want bool
	}{
		{"", true},
		{"aab", true},
		{"aa", true},
		{"tactcoa", true},
		{"c", true},
		{"kyak", false},
	}
	for _, c := range cases {
		got := isPalindromePermutation(c.in1)
		if got != c.want {
			t.Errorf("isPalindromePermutation(%q) == %q, want %v", c.in1, got, c.want)
		}
	}
}
