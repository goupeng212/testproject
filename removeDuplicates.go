package main

import (
	"fmt"
	"math"
	//	"bytes"
	//	"sort"
	"strconv"
	"strings"
)

func containsDuplicate(nums []int) bool {
	//nums = MergeSorts(nums)
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false

}

func MergeSorts(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}
	length := len(nums)
	middle := length / 2
	left := MergeSorts(nums[:middle])
	right := MergeSorts(nums[middle:])
	nums = merges(left, right)
	return nums

}
func merges(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}

	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
func intersect(nums1 []int, nums2 []int) []int {
	var result []int

	l1, l2 := len(nums1), len(nums2)
	if l1 <= l2 {
		result = subset(nums1, nums2)
	} else {
		result = subset(nums2, nums1)
	}
	return result

}
func subset(short, large []int) (result []int) {
	for i := 0; i < len(short); i++ {
		for j := 0; j < len(large); j++ {
			if short[i] == large[j] {
				result = append(result, short[i])
				large = append(large[:j], large[j+1:]...)
				break
			}
		}
	}
	return result
}

func reverse(x int) int {
	max := strconv.Itoa(2 << 30)
	fmt.Println("max:", max)
	maxLen := len(max)
	neg := 0
	str := strconv.Itoa(x)
	str0 := strings.Split(str, "")
	if str0[0] == "-" {
		neg = 1
		str0 = str0[1:]
		str = strings.Join(str0, "")
	}
	//	strings.Replace()

	revertstr := convertStr(str)
	fmt.Println("revert:", revertstr)
	revertLen := len(revertstr)
	if revertLen == maxLen {
		result := strings.Compare(revertstr, max)
		if result == 1 {
			//revertstr大于最大； 越界
			return 0
		}
	}
	if neg == 1 {
		revertstr = "-" + revertstr
	}
	i, _ := strconv.Atoi(revertstr)
	fmt.Println(i)
	return i
}
func convertStr(str string) string {
	arr := strings.Split(str, "")
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return strings.Join(arr, "")
}
func isPalindrome(s string) bool {

	lowS := strings.ToLower(s)
	lowS = AlphaBeticClean(lowS)
	i, j := 0, len(lowS)-1
	for i <= j {
		fmt.Println("current i:", string(lowS[i]), "currentj:", string(lowS[j]))

		if lowS[i] == lowS[j] {
			if i == j || i == j-1 {
				return true

			} else {
				i++
				j--
				continue
			}

		} else {
			return false
		}

	}
	return false
}
func AlphaBeticClean(s string) string {
	var result []rune
	//	temp := strings.Split(s, "")
	for _, v := range s {
		if checkAlphaBetic(v) {
			result = append(result, v)
		}
	}

	return string(result)
}
func checkAlphaBetic(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
func myAtoi(str string) int {

	str = strings.TrimLeft(str, " ")
	if len(str) == 0 || (len(str) == 1 && (str[0] == '-' || str[0] == '+')) {
		return 0
	}
	if (str[0] == '-' || str[0] == '+') && (str[1] < '0' || str[1] > '9') {
		return 0
	}

	fmt.Println(str)
	if (str[0] != '-' && str[0] != '+') && (str[0] < '0' || str[0] > '9') {
		return 0
	}
	p := 1
	if str[0] == '-' {
		p = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	fmt.Println(str)
	str = strings.TrimLeft(str, "0")
	if len(str) == 0 || str[0] < '0' || str[0] > '9' {
		return 0
	}
	i := 0
	for i < len(str) {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		i++
	}
	str = str[:i]
	maxInt := math.MaxInt32
	minInt := math.MinInt32
	s_maxInt := strconv.Itoa(maxInt)
	if len(str) > len(s_maxInt) {
		if p == 1 {
			return maxInt
		} else {
			return minInt

		}

	} else if len(str) == len(s_maxInt) {
		if (p == -1 && strings.Compare(str, s_maxInt) == 1) || (p == 1 && strings.Compare(str, s_maxInt) == -1) {
			return minInt
		} else {
			i, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			return i * p
		}

	} else {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return i * p
	}

}

//func main() {
//	a := "race a car"
//	fmt.Println(isPalindrome(a))
//	fmt.Println(math.MaxInt32, math.MinInt32)
//	fmt.Println(myAtoi("+0000000000000123"))
//	fmt.Println(strings.Index("alibaba", ""))
//	arr := []int{10, 4, 2, 4, 1, 5, 7, 9, 7, 6, 1, 1}
//	QuickSort(arr, 0, len(arr)-1)
//	fmt.Println(arr)
//}
