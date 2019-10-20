package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	//var aString string = "0123456789"
	for i := 48; i <= 55; i++ {
		for j := 49; j <= 56; j++ {
			for k := 50; k <= 57; k++ {
				if i < j && j < k {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					z01.PrintRune(' ')
				}
			}

		}
	}
	z01.PrintRune(10)

}
