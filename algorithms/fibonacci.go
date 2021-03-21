package algorithms

// 0, 1, 1, 2, 3, 5, 8
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
