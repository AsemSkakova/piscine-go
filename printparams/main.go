package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	a := []rune(arguments[i])

	for index, letter := range a {
		if index > 0 {
			z01.PrintRune[index]
		}
	}

}
