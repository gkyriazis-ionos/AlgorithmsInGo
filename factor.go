package AlgorithmsWithGo

// Factor breaks a number into different prime factors
func Factor(primes []int, number int) []int {
	var res []int
	for _, prime := range primes {
		for number%prime == 0 {
			res = append(res, prime)
			number = number / prime
		}
	}
	return res
}
