package main

import "fmt"

func longestPalindrome(s string) string {
	length := len(s)
	maxLen := 1
	start_index := 0
	if length <= 1 {
		return s
	}
	dp := make([][]int, length)
	for k := 0; k < length; k++ {
		dp[k] = make([]int, length)
		dp[k][k] = 1

		if k > 0 && s[k] == s[k-1] {
			dp[k-1][k] = 1
			maxLen = 2
			start_index = k - 1
		}
	}
	fmt.Println(dp)
	for i := 2; i < length; i++ {
		for start := 0; start < length-i; start++ {
			end := start + i
			if s[start] == s[end] && dp[start+1][end-1] == 1 {
				dp[start][end] = 1
				if end-start+1 > maxLen {
					maxLen = end - start + 1
					start_index = start
				}
			}

		}
	}
	fmt.Println(start_index, maxLen)
	fmt.Println(string(s[start_index : start_index+maxLen]))
	return string(s[start_index : start_index+maxLen])

}

//func main() {
//	//s := "cbbd"

//	//longestPalindrome(s)
//	arr := []int{1, 2, 3, 4, 5}
//	arr1 := arr[0:0]
//	arr2 := arr[5:]
//	fmt.Print(arr1, arr2)
//}
