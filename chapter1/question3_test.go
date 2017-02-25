package chapter1

import (
	"reflect"
	"testing"
)

func TestURLify(t *testing.T) {
	cases := []struct {
		in1  []rune
		in2  int
		want []rune
	}{
		{[]rune("Mr John Smith    "), 13, []rune("Mr%20John%20Smith")},
	}
	for _, c := range cases {
		urlify(c.in1, c.in2)
		if !reflect.DeepEqual(c.in1, c.want) {
			t.Errorf("urlify(%q, %q). Got %q, want %q", c.in1, c.in2, c.in1, c.want)
		}
	}
}
