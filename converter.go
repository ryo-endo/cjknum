package cjknum

var characters = map[int]string{
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

func Itoc(i int) (cjk string, err error) {
	err = nil
	cjk = characters[i]

	return
}
