package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var a rune = 48
	var b rune = 49
	var c rune = 50
	for i := a; i <= 54; i++ {
		for j := b; j <= 55; j++ {
			for k := c; k <= 56; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}

		}
	}
	z01.PrintRune(55)
	z01.PrintRune(56)
	z01.PrintRune(57)
	z01.PrintRune(10)

}
