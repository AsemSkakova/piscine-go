package piscine

func StrRev(s string) string {

	aChange := []rune(s)
	a := 0

	var counter int
	for range aChange {
		counter++
	}

	for i := counter - 1; i >= 0; i-- {
		aChange[i] = rune(s[a])
		a++
	}

	aFinalized := string(aChange)
	return aFinalized
}
