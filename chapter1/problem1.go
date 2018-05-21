package chapter1

// O(n^2)
func isUniqBase(str string) bool {
	for _, baseChar := range str {
		count := 0
		for _, char := range str {
			if baseChar == char {
				count = count + 1
			}
		}
		if count > 1 {
			return false
		}
	}

	return true
}

// O(n) but maybe less
func isUniq(str string) bool {
	m := make(map[rune]struct{})
	for _, char := range str {
		if _, ok := m[char]; ok {
			return false
		}
		m[char] = struct{}{}
	}

	return true
}
