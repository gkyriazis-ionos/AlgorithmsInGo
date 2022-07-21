package AlgorithmsWithGo

func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		a, b = b, a%b
		return GCD(a, b)
	}
}
