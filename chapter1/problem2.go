package chapter1

func isPermutation(str1, str2 string) bool {
	counts := make(map[rune]int)

	if len(str1) != len(str2) {
		return false
	}

	for _, char := range str1 {
		counts[char]++
	}

	for _, char := range str2 {
		counts[char]--
	}

	for _, val := range counts {
		if val != 0 {
			return false
		}
	}

	return true
}
