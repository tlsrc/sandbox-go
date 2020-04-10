package eulerproject

import (
	"strconv"
)

//Problem0004 solves problem 4
func Problem0004() int {
	largest := 0
	for a := 999; a >= 100; a-- {
		for b := 999; b >= a; b-- {
			num := a * b
			if num > largest && isPalindrome(num) {
				largest = num
			}
		}
	}
	return largest
}

func isPalindrome(i int) bool {
	s := strconv.Itoa(i)
	if len(s) < 2 {
		return true
	}

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
