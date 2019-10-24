package piscine

func IsAlpha(str string) bool {

	a := []rune(str)
	for _, letter := range a {
		if checkAlpNum(letter) == false {
			return false
		}
	}
	return true

}

func checkAlpNum(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}
