package chapter1

//aaa -> len = 3 and total 3

import (
	"unicode"
)

// O(n)
func isPalindromePermutation(input string) bool {
	m := make(map[rune]int)
	count := 0

	if len(input) == 1 || len(input) == 0 {
		return true
	}

	for _, char := range input {
		if string(char) != " " {
			lowerChar := unicode.ToLower(char)
			m[lowerChar] = (m[lowerChar] + 1) % 2
		}
	}

	for _, val := range m {
		count = count + val
		if count > 1 {
			return false
		}
	}
	return true
}
