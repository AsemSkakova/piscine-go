package piscine

func IterativeFactorial(nb int) int {
	result := 0

	if nb < 0 || nb > 50 {
		return 0
	} else {
		for i := 0; i < nb+1; i++ {
			result = result + i
		}
		return result
	}
}
