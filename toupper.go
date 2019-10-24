package piscine

func ToUpper(s string) string {

	a := []rune(s)

	for index, letter := range s {
		if letter >= 97 && letter <= 122 {
			a[index] = letter - 32
		}
	}

	return string(a)
}
