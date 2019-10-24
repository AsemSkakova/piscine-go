package piscine

func ToLower(s string) string {

	a := []rune(s)

	for index, letter := range s {
		if letter >= 65 && letter <= 90 {
			a[index] = letter + 32
		}
	}

	return string(a)

}
