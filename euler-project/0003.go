package eulerproject

import (
	"math"
)

//Problem0003 solves problem 3
func Problem0003() int {
	number := 600851475143

	for i := 2; i < int(math.Sqrt(float64(number)))+1; i++ {
		if number%i == 0 {
			number = number / i
			i = 2
		}
	}
	return number
}
