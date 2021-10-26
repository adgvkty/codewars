package kata

// MaximumSubarraySum ...
func MaximumSubarraySum(numbers []int) int {
	var result int
	var sums []int

	if len(numbers) == 0 {
		return 0
	}

	for numberID, number := range numbers {
		numberID++
		for ; numberID < len(numbers); numberID++ {
			number += numbers[numberID]
			sums = append(sums, number)
		}
	}

	for _, number := range sums {
		if number > result {
			result = number
		}
	}
	return result
}
