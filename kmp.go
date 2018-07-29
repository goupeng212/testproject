package main

import (
	_ "fmt"
	"strings"
)

func kmpSubstr(source, pattern string) int {
	sLen := len(source)
	pLen := len(pattern)
	if sLen < pLen {
		return -1
	}

	next := getPatternNext(pattern)
	i, j := 0, 0
	for i < sLen && j < pLen {

		if j == -1 || source[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}

	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}

}

func getPatternNext(pattern string) []int {
	if len(pattern) == 0 {
		return nil
	}

	pA := strings.Split(pattern, "")
	pAlen := len(pA)
	next := make([]int, pAlen)
	next[0] = -1
	k := -1
	j := 0
	for j < pAlen-1 {
		if k == -1 || pA[k] == pA[j] {
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
//	s := "ababceabcdeabc"
//	p := "cea"
//	index := kmpSubstr(s, p)
//	fmt.Println(index)
//}
