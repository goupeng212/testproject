package main

import (
	_ "fmt"
	"strings"
)

func lcs(x, y string) int {
	lx := len(x)
	ly := len(y)
	ax := strings.Split(x, "")
	ay := strings.Split(y, "")
	maxlen := 0
	var dp [][]int
	for k := 0; k < lx; k++ {
		slicey := make([]int, ly)
		dp = append(dp, slicey)
	}
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if ax[i] == ay[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1

				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if maxlen < dp[i][j] {
					maxlen = dp[i][j]
					//					maxindex = i - maxlen + 1
				}
			}
		}
	}

	return maxlen

}

//func main() {
//	a := " "
//	b := "JI"
//	fmt.Println(lcs(a, b))
//	strings.Compare(a)
//}
