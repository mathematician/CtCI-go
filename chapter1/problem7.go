package chapter1

func rotateMatrix(input [][]int) {
	n := len(input)
	for layer := 0; layer < n/2; layer++ {
		last := n - 1 - layer
		for i := layer; i < last; i++ {
			offset := i - layer
			temp := input[layer][i]
			input[layer][i] = input[last-offset][layer]
			input[last-offset][layer] = input[last][last-offset]
			input[last][last-offset] = input[i][last]
			input[i][last] = temp
		}
	}
}
