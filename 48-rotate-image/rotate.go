package rotate

func rotate(matrix [][]int) {
	r := len(matrix) / 2
	l := len(matrix[0])

	for i := 0; i < r; i++ {
		for j := i; j < l-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[l-1-j][i]
			matrix[l-1-j][i] = matrix[l-1-i][l-1-j]
			matrix[l-1-i][l-1-j] = matrix[j][l-1-i]
			matrix[j][l-1-i] = temp
		}
	}
}
