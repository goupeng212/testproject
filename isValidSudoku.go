package main

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < len(board); i++ {
		colmap := make(map[byte]bool)
		rowmap := make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {

				if _, ok := rowmap[board[i][j]]; ok {
					return false
				} else {
					rowmap[board[i][j]] = true
				}
			}
			if board[j][i] != '.' {

				if _, ok := colmap[board[j][i]]; ok {
					return false
				} else {
					colmap[board[j][i]] = true
				}
			}
			if i%3 == 0 && j%3 == 0 {
				result := checkblock(i, j, board)
				if !result {
					return false
				}
			}
		}

	}
	return true
}
func checkblock(row int, column int, board [][]byte) bool {

	blockmap := make(map[byte]bool)
	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := blockmap[board[i][j]]; ok {
				return false
			} else {
				blockmap[board[i][j]] = true
			}
		}
	}
	return true

}
