package main

import _ "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	dp := make([]int, len(nums))
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < dp[i-1]+nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max

}

//func main() {
//	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
//	fmt.Println(maxSubArray(a))
//}
