package kata

// Solution ...
func Solution(ar []int) int {

	if len(ar) == 1 {
		return ar[0]
	}
	for {

		// fmt.Println(ar)

		f, t, index := findLargest(ar)

		if index == -1 {
			return sum(ar)
		}

		ar[index] = f - t

	}
}

func findLargest(ar []int) (int, int, int) {

	var f, t, index int

	for i := 0; i < len(ar); i++ {
		if f < ar[i] {
			f = ar[i]
			index = i
		}
	}

	for i := 0; i < len(ar); i++ {
		if t < ar[i] && ar[i] != f {
			t = ar[i]
		}
	}

	if f == t || t == 0 {
		index = -1
	}

	return f, t, index
}

func sum(ar []int) int {
	var result int

	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}

	return result
}
