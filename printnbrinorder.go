package piscine

import "github.com/01-edu/z01"

func toDigits(n int) (digits []int, counter int) {
	counter = 0
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n = n / 10
		counter++
	}
	return
}

func Sorting(nn []int, cc int) []int {
	for sort := true; sort; {
		sort = false
		for i := 1; i < cc; i++ {
			if nn[i] < nn[i-1] {
				nn[i-1], nn[i] = nn[i], nn[i-1]
				sort = true
			}
		}
	}
	return nn
}
func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	} else if n == 0 {
		z01.PrintRune('0')
		return
	} else {
		digits := Sorting(toDigits(n))
		for _, letter := range digits {

			z01.PrintRune(rune(letter) + '0')
		}
	}

}
