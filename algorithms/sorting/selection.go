package sorting

func Selection(numbers []int) []int {
	arrayLength := len(numbers)

	for i := 0; i < arrayLength; i++ {
		min := i
		temp := numbers[i]
		for j := i + 1; j < arrayLength; j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		numbers[i] = numbers[min]
		numbers[min] = temp
	}

	return numbers
}
