package piscine

func MakeRange(min, max int) []int {
	size := 0

	if min < max {

		size = max - min
	}
	answer := make([]int, size)

	if min >= max {
		answer = nil
	}
	for i := 0; i < size; i++ {
		answer[i] = min
		min++
	}
	return answer

}
