package cjknum

var digits = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}

var digitsBases = []struct {
	n int
	c string
}{
	{10, "十"},
	{100, "百"},
	{1000, "千"},
}

var bases = []string{"万"}

func Itoc(i int) (cjk string, err error) {
	// 0 ... 9999
	cjk, err = thousandsUnits(i)
	if err != nil {
		return "", err
	}

	// 10000 ... n
	for _, baseChar := range bases {
		i = i / 10000
		if i == 0 {
			break
		}

		unit, err := thousandsUnits(i)
		if err != nil {
			return "", err
		}

		cjk = unit + baseChar + cjk
	}

	return
}

func thousandsUnits(i int) (cjk string, err error) {
	err = nil
	cjk = ""

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
