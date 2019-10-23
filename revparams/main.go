package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	len := 0

	for index := range arguments {
		len = index + 1
	}

	for i := len; i >= 0; i-- {
		strAsRune := []rune(arguments[i])
		for _, words := range strAsRune {
			z01.PrintRune(words)
		}
		z01.PrintRune('\n')
	}
}
