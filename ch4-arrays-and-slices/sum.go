package arraysandslices

func Sum(numbers []int) int {
	var count int
	for _, value := range numbers {
		count += value
	}
	return count
}
