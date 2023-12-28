package array

func SumArray(numbers [6]int) int {
	sum := 0
	for i := 0; i < 6; i++ {
		sum += numbers[i]
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers [][]int) []int {
	sum := []int{}
	for _, inner := range numbers {
		sum = append(sum, SumSlice(inner))
	}
	return sum
}