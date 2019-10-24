package piscine

func AlphaCount(str string) int {

	a := []rune(str)
	counter := 0

	for _, letter := range a {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			counter++
		}
	}
	return counter

}
