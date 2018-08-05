package main

import (
	_ "fmt"
	"strings"
)

func strStr(haystack string, needle string) int {

	lenhay := len(haystack)
	lenneed := len(needle)
	if lenneed == 0 {
		return 0
	}
	if lenhay < lenneed {
		return -1
	}
	next := getPatternNext2(needle)
	i, j := 0, 0
	for i < lenhay && j < lenneed {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == lenneed {
		return i - j
	} else {
		return -1
	}

}
func getPatternNext2(pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}
	next := make([]int, len(pattern))
	arr := strings.Split(pattern, "")
	next[0] = -1
	k := -1
	j := 0
	for j < len(arr)-1 {
		if k == -1 || arr[k] == arr[j] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next

}

//func main() {
//	hay := "mississippi"
//	needle := "sip"
//	fmt.Println(strStr(hay, needle))
//}
