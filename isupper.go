package piscine

func IsUpper(str string) bool {
	a := []rune(str)
	for _, letter := range a {
		if checkUpperCase(letter) == false {
			return false
		}
	}
	return true

}

func checkUpperCase(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
