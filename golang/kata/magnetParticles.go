package kata

import "math"

// Doubles ...
func Doubles(maxk, maxn int) float64 {

	var result float64

	var k float64
	var n float64

	for k = 1; k <= float64(maxk); k++ {
		for n = 1; n <= float64(maxn); n++ {
			result += 1 / (k * math.Pow(n+1, 2*k))
		}
	}

	return result
}
