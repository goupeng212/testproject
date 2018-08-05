package main

import (
	_ "fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	var dp [][]int
	for i := 0; i < m; i++ {
		t := make([]int, n)
		dp = append(dp, t)
	}
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			} else if i == 0 && j > 0 {
				dp[0][j] = dp[0][j-1]
			} else if j == 0 && i > 0 {
				dp[i][0] = dp[i-1][0]
			} else if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]

}

//func main() {
//	a := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
//	fmt.Println(uniquePathsWithObstacles(a))
//}
