package sprint

func TransposeMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)

	for i := range result {
		result[i] = make([]int, rows)
	}

	for i, row := range matrix {
		for j, val := range row {
			result[j][i] = val
		}
	}

	return result
}
