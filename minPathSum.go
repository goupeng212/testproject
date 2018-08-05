package main

import _ "fmt"

func minPathSum(grid [][]int) int {
	var dp [][]int
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		t := make([]int, n)
		dp = append(dp, t)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i > 0 && j > 0 {
				min := 0
				if dp[i-1][j] > dp[i][j-1] {
					min = dp[i][j-1]
				} else {
					min = dp[i-1][j]
				}
				dp[i][j] = min + grid[i][j]
			}

		}
	}
	return dp[m-1][n-1]

}

//func main() {
//	a := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
//	fmt.Println(minPathSum(a))
//}
