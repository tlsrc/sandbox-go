package eulerproject

//Problem0006 solves problem 6
func Problem0006() int {
	sumOfSquares := 0
	squareOfSum := 0
	for i := 0; i <= 100; i++ {
		squareOfSum += i
		sumOfSquares += i * i
	}
	squareOfSum *= squareOfSum
	return squareOfSum - sumOfSquares
}
