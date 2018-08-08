package cjknum_test

import (
	"cjknum"
	"strings"
	"testing"
)

func TestDigit(t *testing.T) {
	cases := []struct {
		num int
		cjk string
	}{
		{1, "一"},
		{9, "九"},
	}

	for _, c := range cases {
		s, err := cjknum.Itoc(c.num)
		if (strings.Compare(s, c.cjk) != 0) || (err != nil) {
			t.Errorf("Itoc(%v) = %v want %v", c.num, s, c.cjk)
		}
	}

}
