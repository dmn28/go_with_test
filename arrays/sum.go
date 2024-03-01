package arrays

func Sum(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sum_slices []int
	for _, array := range numbers {
		sum_slices = append(sum_slices, Sum(array))
	}

	return sum_slices
}

func SumAllTails(numbers ...[]int) []int {
	var sum_tails []int
	for _, array := range numbers {
		if len(array) == 0 {
			sum_tails = append(sum_tails, 0)
		} else {
			sum_tails = append(sum_tails, Sum(array[1:]))
		}
	}

	return sum_tails
}
