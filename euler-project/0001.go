package eulerproject

//Problem0001 solves first problem
func Problem0001() int {
	var res int = 0
	for i := 1; i < 1000; i++ {
		if i%3 == 00 || i%5 == 0 {
			res += i
		}
	}
	return res
}
