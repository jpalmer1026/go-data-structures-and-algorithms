package sorting

func Bubble(numbers []int) []int {
	arrayLength := len(numbers)

	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}
