package main

import (
	"fmt"
	_ "math"
	_ "sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	cur_a, cur_b := 0, 0
	var result []int
	for cur_a < len(nums1) && cur_b < len(nums2) {
		if nums1[cur_a] < nums2[cur_b] {
			result = append(result, nums1[cur_a])
			cur_a++
		} else {
			result = append(result, nums2[cur_b])
			cur_b++
		}
	}
	if cur_a == len(nums1) {
		result = append(result, nums2[cur_b:len(nums2)]...)

	} else if cur_b == len(nums2) {
		result = append(result, nums1[cur_a:len(nums1)]...)

	}
	fmt.Println(result)
	if len(result)%2 == 0 {

		return float64((result[len(result)/2-1] + result[len(result)/2]) / 2)
	} else {
		return float64(result[len(result)/2])
	}
	//	math.MaxInt32

}

//func main() {
//	a := []int{1}
//	b := []int{1, 4}
//	findMedianSortedArrays(a, b)
//	x := 1001
//	fmt.Println(x % 10)
//	i, j := 1, 2
//	compare(i, j)

//}

func compare(i, j int) int {
	switch {
	case i-j > 0:
		return 1
	case i-j < 0:
		return -1
	case i-j == 0:
		return 0
	default:
		return 0
	}
}
