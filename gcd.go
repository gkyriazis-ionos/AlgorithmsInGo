package AlgorithmsWithGo

// Calculates recursively the Greater Common Divisor of a number according to the Euclidean Algorithm

func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		a, b = b, a%b
		return GCD(a, b)
	}
}
