package piscine

func BasicAtoi2(s string) int {

	a := 0
	for _, letter := range s {
		ed := 0
		if letter < '0' || letter > '9' {
			return 0
		}
		for i := '1'; i <= letter; i++ {
			ed++
		}

		a = a*10 + ed
	}
	return a
}
