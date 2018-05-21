package chapter1

func zeroMatrix(input [][]int) {
	zeroRows := []int{}
	zeroCols := []int{}
	M := len(input[0])

	for i, row := range input {
		for j, val := range row {
			if val == 0 {
				zeroRows = append(zeroRows, i)
				zeroCols = append(zeroCols, j)
			}
		}
	}

	for _, row := range zeroRows {
		for j := 0; j < M; j++ {
			input[row][j] = 0
		}
	}

	for i, _ := range input {
		for _, col := range zeroCols {
			input[i][col] = 0
		}
	}
}
