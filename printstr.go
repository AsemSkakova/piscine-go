package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {

	//strToRune := []rune(str)

	for _, word := range []rune(str) {
		z01.PrintRune(word)
	}
}
