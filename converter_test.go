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
		{0, "零"},
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

func Test10Position(t *testing.T) {
	cases := []struct {
		num int
		cjk string
	}{
		{10, "十"},
		{20, "二十"},
		{90, "九十"},
		{11, "十一"},
		{21, "二十一"},
		{99, "九十九"},
	}

	for _, c := range cases {
		s, err := cjknum.Itoc(c.num)
		if (strings.Compare(s, c.cjk) != 0) || (err != nil) {
			t.Errorf("Itoc(%v) = %v want %v", c.num, s, c.cjk)
		}
	}

}
