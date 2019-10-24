package piscine

func IsNumeric(str string) bool {

	a := []rune(str)
	for _, letter := range a {
		if isNumeric(letter) == false {
			return false
		}
	}
	return true

}

func isNumeric(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false

}
