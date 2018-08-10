package cjknum_test

import (
	"cjknum"
	"math"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		num uint64
		cjk string
	}{
		{0, "零"},
		{1, "一"},
		{9, "九"},
		{10, "十"},
		{20, "二十"},
		{90, "九十"},
		{11, "十一"},
		{21, "二十一"},
		{99, "九十九"},
		{100, "百"},
		{909, "九百九"},
		{999, "九百九十九"},
		{1000, "千"},
		{1001, "千一"},
		{1999, "千九百九十九"},
		{10000, "一万"},
		{12345678, "千二百三十四万五千六百七十八"},
		{99999999, "九千九百九十九万九千九百九十九"},
		{math.MaxUint64, "千八百四十四京六千七百四十四兆七百三十七億九百五十五万千六百十五"},
	}

	for _, c := range cases {
		s, err := cjknum.Itoc(c.num)
		if (strings.Compare(s, c.cjk) != 0) || (err != nil) {
			t.Errorf("Itoc(%v) = %v want %v", c.num, s, c.cjk)
		}
	}

}
