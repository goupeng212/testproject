package main

import (
	"bytes"
	"strconv"
	//	"strings"
)

func combine(nums []int, result *[][]int, k int) {
	if k == 0 {
		tmp := []int{}
		*result = append(*result, tmp)
		return
	}
	var temp []int
	combineDFS(nums, result, &temp, 0, k)

}
func combineDFS(nums []int, result *[][]int, temp *[]int, start, k int) {
	if len(*temp) == k {
		var t []int
		for _, v := range *temp {
			t = append(t, v)
		}
		*result = append(*result, t)
		return
	}
	m := make(map[string]bool)
	for i := start; i < len(nums); i++ {
		buff := bytes.NewBufferString("")
		for _, v := range *temp {
			buff.WriteString(strconv.Itoa(v))
		}
		buff.WriteString(strconv.Itoa(nums[i]))
		skey := buff.String()
		if _, ok := m[skey]; ok {
			continue
		} else {
			m[skey] = true
		}
		*temp = append(*temp, nums[i])
		combineDFS(nums, result, temp, i+1, k)
		*temp = (*temp)[:(len(*temp) - 1)]
	}
}
