package AlgorithmsWithGo

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//

// Use a helper function so that we can pass in an initialized map.
func Fibonacci(n int) int {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	return helper(n, cache)
}

// This helper function does all the heavy lifting. If our desired fib number
// isn't in the cache, then calculate it first and store it in the cache before
// returning the answer.
func helper(n int, cache map[int]int) int {
	i, ok := cache[n]
	if ok {
		return i
	}
	cache[n] = helper(n-1, cache) + helper(n-2, cache)
	return cache[n]
}
