package arraysandslices

func Sum(numbers [5]int) int {
	var count int
	for i := 0; i < len(numbers); i++ {
		count += numbers[i]
	}
	return count
}
