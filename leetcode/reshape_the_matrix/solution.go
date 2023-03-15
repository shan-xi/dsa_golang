package main

// 566. Reshape the Matrix
// https://leetcode.com/problems/reshape-the-matrix/description/?envType=study-plan&id=data-structure-i
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	row := 0
	col := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[row][col] = mat[i][j]
			col++
			if col == c {
				row++
				col = 0
			}
		}
	}
	return res
}
