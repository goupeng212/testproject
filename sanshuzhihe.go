package main

import (
	_ "fmt"
	"math"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) <= 2 {
		re := make([]int, 3)
		result = append(result, re)
		return result
	}
	sort.Ints(nums)
	pivot := 0
	abstrV := math.MaxInt32
	for i := 0; i < len(nums); i++ {

		if math.Abs(float64(nums[i]-0)) < float64(abstrV) {
			abstrV = int(math.Abs(float64(nums[i] - 0)))
			pivot = i
		}
	}

	for left := 0; left < pivot; {
		for right := len(nums) - 1; right > pivot; {

			//		fmt.Println(nums[left]+nums[right], left, right)
			tmp := nums[left] + nums[right]
			if tmp > 0 {
				for i := pivot; i > left; i-- {
					if tmp+nums[i] > 0 {
						continue
					} else if tmp+nums[i] == 0 {
						re := []int{nums[left], nums[i], nums[right]}
						result = append(result, re)
						break
					} else {
						break
					}
				}
			}
			if tmp < 0 {
				for i := pivot; i < right; i++ {
					if tmp+nums[i] < 0 {
						continue
					} else if tmp+nums[i] == 0 {
						re := []int{nums[left], nums[i], nums[right]}
						result = append(result, re)
						break
					} else {
						break
					}
				}
			}
			if tmp == 0 {
				if nums[pivot] == 0 {
					re := []int{nums[left], 0, nums[right]}
					result = append(result, re)
					break
				}

			}
			j := right - 1
			for ; j > pivot; j-- {
				if nums[j] != nums[right] {
					break
				}
			}
			right = j
		}
		j := left + 1
		for ; j < pivot; j++ {
			if nums[j] != nums[left] {
				break
			}
		}
		left = j
	}
	return result
}

//func main() {
//	a := []int{-1, 0, 1, 2, -1, -4}

//	fmt.Println(threeSum(a))
//}
