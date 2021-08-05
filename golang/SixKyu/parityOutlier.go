package six

// FindOutlier ...
func FindOutlier(integers []int) int {
	var odd []int
	var even []int

	for _, number := range integers {
		if number%2 == 0 {
			odd = append(odd, number)
		} else {
			even = append(even, number)
		}
	}

	if len(odd) == 1 {
		return odd[0]
	}

	return even[0]
}
