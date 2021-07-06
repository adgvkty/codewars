package seven

func Divisors(n int) int {
	count := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			count += 1
		}
	}
	return count + 1
}
