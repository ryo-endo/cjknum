package cjknum

var digits = map[int]string{
	0: "零",
	1: "一",
	2: "二",
	3: "三",
	4: "四",
	5: "五",
	6: "六",
	7: "七",
	8: "八",
	9: "九",
}

var digitsBase = map[int]string{
	10:  "十",
	100: "百",
}

func Itoc(i int) (cjk string, err error) {
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
	for base, c := range digitsBase {
		if (i / base) == 0 {
			break
		}

		// get the number of base position
		n := (i / base) % 10

		s := ""
		if n >= 2 {
			s = digits[n]
		}
		if n > 0 {
			s += c
		}

		cjk = s + cjk
	}

	return
}
