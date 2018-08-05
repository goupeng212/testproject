package main

import (
	"math"
)

func solveNQueens(n int) [][]string {

	var result [][]string
	current := make([]int, n)
	for i := 0; i < n; i++ {
		current[i] = -1
	}

	nQueenHelper(&result, &current, 0, n)
	return result

}
func nQueenHelper(result *[][]string, current *[]int, row int, n int) {

	if row == n {

		temp := make([]string, n)
		for i, v := range *current {
			s := ""
			for j := 0; j < n; j++ {
				if j != v {
					s = s + "."
				} else {
					s = s + "Q"
				}
			}
			temp[i] = s

		}
		*result = append(*result, temp)
		return
	}
	for col := 0; col < n; col++ {
		if isValidNQueen(current, row, col) {
			(*current)[row] = col
			nQueenHelper(result, current, row+1, n)
			(*current)[row] = -1
		}
	}

}

func isValidNQueen(current *[]int, row, col int) bool {

	for i := 0; i < row; i++ {
		if (*current)[i] == col || math.Abs(float64(row-i)) == math.Abs(float64((*current)[i]-col)) {
			return false
		}
	}
	return true

}
