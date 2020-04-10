package eulerproject

import "testing"

type PalindromeTestcase struct {
	input    int
	expected bool
}

func TestPalindrome(t *testing.T) {
	testcases := []PalindromeTestcase{
		PalindromeTestcase{input: 1, expected: true},
		PalindromeTestcase{input: 11, expected: true},
		PalindromeTestcase{input: 12, expected: false},
		PalindromeTestcase{input: 121, expected: true},
		PalindromeTestcase{input: 1221, expected: true},
		PalindromeTestcase{input: 1231, expected: false},
		PalindromeTestcase{input: 12321, expected: true},
		PalindromeTestcase{input: 12345, expected: false},
		PalindromeTestcase{input: 123321, expected: true},
	}

	for _, tc := range testcases {
		actual := isPalindrome(tc.input)
		if actual != tc.expected {
			t.Errorf("%d : expected '%t' but got '%t'",
				tc.input, tc.expected, actual)
		}
	}

}
