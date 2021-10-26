package kata

import (
	"fmt"
	"math"
	"sort"
)

// SumOfDivided ...
func SumOfDivided(lst []int) string {

	var resultString string
	var mapKeys []int
	result := make(map[int]int)

	for _, integer := range lst {
		divisors := findDivisors(integer)
		for _, divisor := range divisors {
			result[divisor] += integer
		}
	}

	for key := range result {
		mapKeys = append(mapKeys, key)
	}

	sort.Ints(mapKeys)

	for _, key := range mapKeys {
		resultString += fmt.Sprintf("(%v %v)", key, result[key])
	}
	return resultString
}

func findDivisors(n int) []int {
	result := []int{}
	if n < 0 {
		n = -n
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if isPrime(i) {
				result = append(result, i)
			}

		}
	}
	return result
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
