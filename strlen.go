package piscine

func StrLen(str string) int {

	var counter int
	for _, word := range []rune(str) {
		if word != ' ' {
			counter++
		}
	}
	return counter
}
