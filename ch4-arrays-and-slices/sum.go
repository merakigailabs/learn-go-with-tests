package arraysandslices

func Sum(numbers [5]int) int {
	var count int
	for _, value := range numbers {
		count += value
	}
	return count
}
