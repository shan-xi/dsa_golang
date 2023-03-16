package main

// 74. Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/?envType=study-plan&id=data-structure-i
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}
