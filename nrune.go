package piscine

func NRune(s string, n int) rune {

	a := []rune(s)

	var nrune rune

	for index, letter := range a {
		if index == n-1 {
			nrune = letter
		}
	}

	return nrune

}
