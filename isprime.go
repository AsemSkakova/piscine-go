package piscine

func IsPrime(nb int) bool {

	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}

	//if nb%nb == 0 && nb%1 == 0 {
	//return true
	//}
	return true
}
