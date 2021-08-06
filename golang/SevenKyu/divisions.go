package seven

// Divisors ...
func Divisors(n int) int {
	count := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count + 1
}
