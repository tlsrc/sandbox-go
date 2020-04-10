package eulerproject

//Problem0005 solves problem5
func Problem0005() int {
	i := 2502
	for {
		if isEvenlyDivisible(i) {
			return i
		}
		i++
	}
}

func isEvenlyDivisible(n int) bool {
	for i := 20; i > 0; i-- {
		if n%i != 0 {
			return false
		}
	}
	return true
}
