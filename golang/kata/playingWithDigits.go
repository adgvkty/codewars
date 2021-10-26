package kata

import (
	"fmt"
	"math"
)

// DigPow ...
func DigPow(n, p int) int {
	var result int
	var digit int
	var numbersLeft int = n

	intString := fmt.Sprint(n)
	pow := len(intString) + p - 1
	for numbersLeft != 0 {
		digit = numbersLeft % 10
		numbersLeft = numbersLeft / 10

		result += int(math.Pow(float64(digit), float64(pow)))
		pow--

	}

	if result%n == 0 {
		return result / n
	}
	return -1
}
