package chapter1

func isOneChangeAway(biggerInput string, smallerInput string) bool {

	for i, j := 0, 0; j <= len(smallerInput); i, j = i+1, j+1 {
		slicedString := biggerInput[:i] + biggerInput[i+1:]
		if smallerInput == slicedString {
			return true
		}
	}
	return false
}

func isOneAway(input1 string, input2 string) bool {
	len1 := len(input1)
	len2 := len(input2)
	countWrongPosition := 0

	if len1 != len2 && (len1-len2 > 1 || len2-len1 > 1) {
		return false
	}

	if len1 == len2 {
		for i, _ := range input1 {
			if input1[i] != input2[i] {
				countWrongPosition++
				if countWrongPosition > 1 {
					return false
				}
			}
		}

		return true
	} else if len1 == len2+1 {
		return isOneChangeAway(input1, input2)
	} else {
		return isOneChangeAway(input2, input1)
	}
}
