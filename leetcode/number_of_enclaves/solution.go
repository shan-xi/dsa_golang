package main

// 1020. Number of Enclaves
// https://leetcode.com/problems/number-of-enclaves/description/

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func numEnclaves(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		if grid[i][0] == 1 {
			dfs(grid, i, 0)
		}
		if grid[i][cols-1] == 1 {
			dfs(grid, i, cols-1)
		}
	}
	for i := 0; i < cols; i++ {
		if grid[0][i] == 1 {
			dfs(grid, 0, i)
		}
		if grid[rows-1][i] == 1 {
			dfs(grid, rows-1, i)
		}
	}
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count

}

func dfs(grid [][]int, i, j int) {
	grid[i][j] = 0
	for _, direction := range directions {
		x := i + direction[0]
		y := j + direction[1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
			dfs(grid, x, y)
		}
	}
}
