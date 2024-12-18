package arraysandslices

func Sum(numbers []int) int {
	var count int
	for _, value := range numbers {
		count += value
	}
	return count
}

func SumAll(numbersToSum ...[]int) []int {
	res := []int{}
	for _, numbers := range numbersToSum {
		res = append(res, Sum(numbers))
	}
	return res
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
