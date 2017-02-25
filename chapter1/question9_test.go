package chapter1

import (
	"testing"
)

func TestIsStringRotation(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want bool
	}{
		{"", "", true},
		{"waterbottle", "erbottlewat", true},
	}
	for _, c := range cases {
		got := isStringRotation(c.in1, c.in2)
		if got != c.want {
			t.Errorf("isStringRotation(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}
