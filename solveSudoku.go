package main

func solveSudoku(board [][]byte) {
	if len(board) == 0 || len(board) != 9 {
		return
	}
	for j := 0; j < len(board); j++ {
		if len(board[j]) != 9 {
			return
		}
	}
	solveSudokuDFS(board, 0, 0)
}
func solveSudokuDFS(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	if j >= 9 {
		return solveSudokuDFS(board, i+1, 0)
	}
	if board[i][j] == '.' {
		for k := 1; k <= 9; k++ {
			board[i][j] = byte(k + '0')
			//			fmt.Println("added number", string(byte(k+'0')))
			if isValidSudoku2(board, i, j) {
				if solveSudokuDFS(board, i, j+1) {
					return true
				}
			}
			board[i][j] = '.'
		}

	} else {

		return solveSudokuDFS(board, i, j+1)
	}
	return false
}

func isValidSudoku2(board [][]byte, row, col int) bool {
	for k := 0; k < 9; k++ {
		if k != col && board[row][k] == board[row][col] {
			return false
		}
	}
	for k := 0; k < 9; k++ {
		if k != row && board[k][col] == board[row][col] {
			return false
		}
	}
	for i := row / 3 * 3; i < row/3*3+3; i++ {
		for j := col / 3 * 3; j < col/3*3+3; j++ {
			if i != row && j != col && board[i][j] == board[row][col] {
				return false
			}
		}
	}
	return true
}
