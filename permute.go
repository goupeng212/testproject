package main

import (
	_ "fmt"
	_ "strconv"
)

//方法一：
//func permute(nums []int) [][]int {
//	if len(nums) == 0 {
//		return nil
//	}
//	var result [][]int
//	var left []int
//	m := make(map[string]bool)

//	permuteDFS(&left, &nums, &result, len(nums), m)
//	return result

//}

//func permuteDFS(left *[]int, right *[]int, result *[][]int, length int, m map[string]bool) {

//	if len(*left) == length {
//		key := ""
//		tmp := make([]int, length)
//		for i, value := range *left {
//			tmp[i] = value
//			key = key + strconv.Itoa(value)
//		}
//		fmt.Println("key:", key, len(key))
//		if _, ok := m[key]; !ok {
//			m[key] = true
//		}
//		//		strconv.Itoa()
//		*result = append(*result, tmp)
//		return
//	}
//	for i := 0; i < len(*right); i++ {
//		tmp := (*right)[i]
//		*left = append(*left, tmp)

//		*right = append((*right)[:i], (*right)[i+1:]...)

//		permuteDFS(left, right, result, length, m)

//		*left = (*left)[:(len(*left) - 1)]
//		tempright := make([]int, len(*right)-i)
//		for j := 0; j < len(*right)-i; j++ {
//			tempright[j] = (*right)[j+i]
//		}
//		*right = append((*right)[:i], tmp)
//		*right = append((*right)[:i+1], tempright[:]...)
//	}

//}

//方法二
func permute(nums []int) [][]int {
	var result [][]int
	permuteDFS(nums, 0, &result)
	return result

}

func permuteDFS(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		temp := make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			temp[i] = nums[i]
		}
		*result = append(*result, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permuteDFS(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

//func main() {
//	nums := []int{1, 2, 3, 4}
//	fmt.Println(permute(nums))
//}
