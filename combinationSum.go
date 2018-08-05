package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)

	var result [][]int
	var temp []int
	combinationSumDFS(candidates, target, 0, &temp, 0, &result)
	return result

}
func combinationSumDFS(candidates []int, target int, index int, tempA *[]int, tempResult int, result *[][]int) bool {
	if tempResult == target {
		//*tempA = append(*tempA, candidates[index])
		var copyA []int

		for _, value := range *tempA {
			copyA = append(copyA, value)
		}
		*result = append(*result, copyA)
		return true
	} else if tempResult > target {
		return false
	} else {
		for i := index; i < len(candidates); i++ {
			tempResult += candidates[i]
			fmt.Println("index:", i, " tempResult:", tempResult)
			*tempA = append(*tempA, candidates[i])
			combinationSumDFS(candidates, target, i, tempA, tempResult, result)
			tempResult -= candidates[i]
			*tempA = (*tempA)[0:(len(*tempA) - 1)]

		}
	}
	return true
}

//func main() {
//	a := []int{1, 2, 3, 4}
//	target := 10
//	fmt.Println(combinationSum(a, target))
//}
