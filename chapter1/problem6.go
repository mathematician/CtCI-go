package chapter1

import "strconv"

func compressString(input string) string {
	output := string(input[0])
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i-1] != input[i] {
			output = output + strconv.Itoa(count) + string(input[i])
			count = 0
		}

		count++
	}

	return output + strconv.Itoa(count)
}
