package algorithms

func Factorial(v int) int {
	if v < 3 {
		return v
	}

	return v * Factorial(v-1)
}
