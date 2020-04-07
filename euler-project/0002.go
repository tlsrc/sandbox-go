package eulerproject

//Problem0002 solves problem 2
func Problem0002() int {
	result := 0
	prev, fib := 0, 1
	for fib < 4_000_000 {
		if fib%2 == 0 {
			result += fib
		}
		prev, fib = fib, fib+prev
	}
	return result
}
