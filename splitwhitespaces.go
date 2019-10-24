package piscine

func SplitWhiteSpaces(str string) []string {
	//var word string
	strRune := []rune(str)

	count := 0

	space := 0
	var split []string
	for i := range strRune {
		count = i + 1
	}

	if count > 0 && IsSep(strRune[0]) == false {
		space++
		for i := 0; i < count; i++ {
			if IsSep(strRune[i]) == true {
				if i+1 < count {
					if IsSep(strRune[i+1]) == false {
						space++
					}
				}
			}
		}
	}

	split = make([]string, space)

	word := ""
	index := 0
	if count > 0 && IsSep(strRune[0]) == false {
		for i := 0; i < count; i++ {
			if IsSep(strRune[i]) == true {
				if i+1 < count {
					if IsSep(strRune[i+1]) == false {
						split[index] = word
						index++
						word = ""
					}
				}
			} else {
				word += string(strRune[i])
			}
		}
		split[space-1] = word
	}

	return split

}

func IsSep(r rune) bool {
	if r == ' ' || r == '\n' || r == '\t' {
		return true
	}
	return false
}
