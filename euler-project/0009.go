/*
a+b+c = 1000
a<b<c
c = math.Sqrt(a2+b2)
*/

package eulerproject

//Problem0009 solves problem 9
func Problem0009() int {
	for b := 1; b < 1000; b++ {
		for a := 1; a < b; a++ {
			c := 1000 - a - b
			if a+b+c == 1000 && a*a+b*b == c*c {
				return a * b * c
			}
		}
	}
	return -1
}
