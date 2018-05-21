package chapter1

import "strings"

//function from question
func isSubstring(str1 string, str2 string) bool {
	return strings.Contains(str1, str2)
}

func isRotatedString(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	return isSubstring(str2+str2, str1)
}
