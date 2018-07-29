package main

import (
	_ "fmt"
	_ "strconv"
	_ "strings"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var result []string
	helper(digits, "", &result, 0)
	return result

}
func helper(digits string, tmp_result string, result *[]string, index int) {

	if index == len(digits) {
		*result = append(*result, tmp_result)
		return
	}
	string_map := []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	for i := 0; i < len(string_map[digits[index]-'0']); i++ {
		m := string_map[digits[index]-'0']

		tmp_result = tmp_result + string(m[i])
		//		fmt.Println("tem_result", tmp_result)
		helper(digits, tmp_result, result, index+1)
		tmp_result = tmp_result[:(len(tmp_result) - 1)]
	}

}

func generateParenthesis(n int) []string {

	if n == 0 {
		return []string{}
	}
	left, right := n, n
	var result []string
	parenthese(&result, "", left, right)
	return result

}
func parenthese(result *[]string, tmp_str string, left, right int) {
	if left == 0 && right == 0 {
		*result = append(*result, tmp_str)
	}
	if left > 0 {
		tmp_str += "("
		parenthese(result, tmp_str, left-1, right)
		tmp_str = tmp_str[:(len(tmp_str) - 1)]
	}
	if left < right {
		tmp_str += ")"
		parenthese(result, tmp_str, left, right-1)
		//tmp_str = tmp_str[:(len(tmp_str) - 1)]

	}

}

//func main() {
//	//digit := "2359"
//	//fmt.Println(letterCombinations(digit))
//	n := 3
//	fmt.Println(generateParenthesis(n))
//}
