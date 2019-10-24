package piscine

func IsLower(str string) bool {
	a := []rune(str)
	for _, letter := range a {
		if checkLowerCase(letter) == false {
			return false
		}
	}
	return true

}

func checkLowerCase(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}
