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

const _10position = "十"

func Itoc(i int) (cjk string, err error) {
	err = nil
	cjk = ""

	if i == 0 {
		return digits[0], nil
	}

	digit := (i / 10) % 10
	if digit > 0 {
		if digit >= 2 {
			cjk += digits[digit]
		}
		cjk += _10position
	}

	digit = i % 10
	if digit > 0 {
		cjk += digits[digit]
	}

	return
}
