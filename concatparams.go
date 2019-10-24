package piscine

func ConcatParams(args []string) string {
	var res string
	for i, word := range args {
		if i == 0 {
			res = word
		} else {
			res = res + "\n" + word
		}
	}
	return res
}
