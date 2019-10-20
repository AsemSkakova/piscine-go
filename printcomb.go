package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var aString string = "0123456789"
	for i := 0; i <= 7; i++ {
		for j := 1; j <= 8; j++ {
			for k := 2; k <= 9; k++ {
				if i < j && j < k {
					z01.PrintRune(rune(aString[i]))
					z01.PrintRune(rune(aString[j]))
					z01.PrintRune(rune(aString[k]))
					z01.PrintRune(' ')
				}
			}

		}
	}
	z01.PrintRune(10)

}
