package piscine

func Atoi(s string) int {

	a := 0
	sign := 0
	index := 0
	var signV rune
	for i, letter := range s {
		ed := 0
		if letter == '+' || letter == '-' {
			index = i
			sign++
			if letter == '-' {
				signV = '-'
			}
			if index > 0 {
				return 0
			}

		} else if letter < '0' || letter > '9' || sign > 1 {
			return 0
		}
		for i := '1'; i <= letter; i++ {
			ed++
		}

		a = a*10 + ed
	}
	if signV == '-' {
		return a * -1
	} else {
		return a
	}
}
