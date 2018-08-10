package cjknum

var digits = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}

var digitsBases = []struct {
	n uint64
	c string
}{
	{10, "十"},
	{100, "百"},
	{1000, "千"},
}

var unitBases = []string{"", "万"}

// Itoc returns the CJK numeric representation of i
func Itoc(i uint64) (cjk string, err error) {
	for _, baseChar := range unitBases {
		unit, err := convertUnit(i)
		if err != nil {
			return "", err
		}
		cjk = unit + baseChar + cjk

		i = i / 10000
		if i == 0 {
			break
		}
	}

	return
}

func convertUnit(i uint64) (cjk string, err error) {
	if i == 0 {
		return digits[0], nil
	}

	// 1 ... 9
	digit := i % 10
	if digit > 0 {
		cjk = digits[digit]
	}

	// 10 ... n
	for _, base := range digitsBases {
		if (i / base.n) == 0 {
			break
		}

		// get the number of base position
		n := (i / base.n) % 10

		s := ""
		if n >= 2 {
			s = digits[n]
		}
		if n > 0 {
			s += base.c
		}

		cjk = s + cjk
	}

	return
}
