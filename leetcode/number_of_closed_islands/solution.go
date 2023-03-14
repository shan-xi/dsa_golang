package main

// 1254. Number of Closed Islands
// https://leetcode.com/problems/number-of-closed-islands/?envType=study-plan&id=graph-i
func closedIsland(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 { // start a new island
				if dfs(grid, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func dfs(grid [][]int, i int, j int) bool {
	row := len(grid)
	col := len(grid[0])
	if i < 0 || i >= row || j < 0 || j >= col { // out of boundary
		return false
	}
	if grid[i][j] == 1 { // hit a water
		return true
	}
	grid[i][j] = 1 // mark the land as water
	up := dfs(grid, i-1, j)
	down := dfs(grid, i+1, j)
	left := dfs(grid, i, j-1)
	right := dfs(grid, i, j+1)
	return up && down && left && right // check if this is a closed island
}
