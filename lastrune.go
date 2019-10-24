package piscine

func LastRune(s string) rune {
	strAsRune := []rune(s)
	//var nrune rune
	counter := 0
	for range strAsRune {
		//nrune = letter
		counter++
	}
	return strAsRune[counter-1]
}
