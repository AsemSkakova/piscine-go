package piscine

func IsPrintable(str string) bool {

	a := []rune(str)
	for _, letter := range a {
		if checkPrint(letter) {
			return false
		}
	}
	return true

}

func checkPrint(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}
