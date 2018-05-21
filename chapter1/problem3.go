package chapter1

func URLify(str string, length int) string {
	cutStr := str[0:length]
	returnStr := ""

	for _, char := range cutStr {
		if string(char) == " " {
			returnStr = returnStr + "%20"
		} else {
			returnStr = returnStr + string(char)
		}
	}

	return returnStr
}
