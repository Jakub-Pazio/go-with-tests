package arraysslices

func SumArray(numbers [5]int) int {
	var result int
	for _, val := range numbers {
		result += val
	}
	return result
}

func SumSlice(numbers []int) int {
	var result int
	for _, val := range numbers {
		result += val
	}
	return result
}

func SumAll(arrs ...[]int) []int {
	var result []int
	for _, arr := range arrs {
		var innerSum int

		for _, val := range arr {
			innerSum += val
		}
		result = append(result, innerSum)
	}
	return result
}

func SumAllTails(arrs ...[]int) []int {
	var result []int
	for _, arr := range arrs {
		var innerSum int
		// if we don't have at least 2 elements sum of tails will always be 0
		if len(arr) < 2 {
			result = append(result, innerSum)
		} else {
			for _, val := range arr[1:] {
				innerSum += val
			}
			result = append(result, innerSum)
		}
	}
	return result
}
