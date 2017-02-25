package chapter1

func zeroMatrix(matrix [][]int) {
	hasRowZero := false
	hasColZero := false
	// m x n matrix
	m := len(matrix)
	n := 0
	if m >= 1 {
		n = len(matrix[0])
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			hasRowZero = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			hasColZero = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if hasRowZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if hasColZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}
