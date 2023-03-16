package main

// 36. Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/?envType=study-plan&id=data-structure-i
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j]
				boxIndex := (i/3)*3 + j/3
				if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
					return false
				}
				rows[i][num] = true
				cols[j][num] = true
				boxes[boxIndex][num] = true
			}
		}
	}

	return true
}
