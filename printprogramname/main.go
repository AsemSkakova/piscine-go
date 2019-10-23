package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	nameAsRune := []rune(os.Args[0])

	for _, letter := range nameAsRune {
		z01.PrintRune(letter)
	}

	z01.PrintRune('\n')
}
