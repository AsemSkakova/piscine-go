package piscine

func BasicAtoi(s string) int {

	a := 0
	ed := 0
	aChange := []rune(s)

	for _, letter := range aChange {
		if letter >= 49 && letter <= 57 {
			ed++
		}
		a = a*10 + ed
	}
	return a
}
