package main

import (
	_ "fmt"
	"strings"
)

func simplifyPath(path string) string {
	var result string
	if len(path) == 0 {
		return ""
	}
	var stack []string
	a := strings.Split(path, "/")
	for _, v := range a {
		switch v {
		case "", " ", ".":
			continue
		case "..":
			if v == ".." && len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, v)

		}

	}
	if len(stack) == 0 {
		return "/"
	}
	for i := 0; i < len(stack); i++ {
		result = result + "/" + stack[i]
	}
	return result
}

//func main() {
//	a := "/home/../var//./dir/../home//"
//	//	aa := strings.Split(a, "/")
//	fmt.Println(simplifyPath(a))
//}
