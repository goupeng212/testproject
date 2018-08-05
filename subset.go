package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var result [][]int

	for i := 0; i <= len(nums); i++ {
		combine(nums, &result, i)
	}
	return result
}
func main() {
	a := []int{1, 2, 2}
	fmt.Println(subsets(a))
}
