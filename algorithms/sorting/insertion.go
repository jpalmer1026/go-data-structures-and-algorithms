package sorting

func Insertion(numbers []int) []int {
	arrayLength := len(numbers)

	for i := 1; i < arrayLength; i++ {
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[i] {
				numbers[j], numbers[i] = numbers[i], numbers[j]
			}
		}
	}

	return numbers
}
