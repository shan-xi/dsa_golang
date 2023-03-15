package main

// 118. Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/?envType=study-plan&id=data-structure-i
func generate(numRows int) [][]int {
	triangle := [][]int{}
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j < len(row); j++ {
			if j == 0 || j == len(row)-1 {
				row[j] = 1
			} else {
				row[j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
		triangle = append(triangle, row)
	}
	return triangle
}
