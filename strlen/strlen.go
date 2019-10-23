package piscine

func StrLen(str string) int {

	var counter int
	for range []rune(str) {
		counter++
	}
	return counter
}
