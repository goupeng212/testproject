package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true

	}
	stack := new(Stack)

	for _, item := range s {
		if item != '{' && item != '(' && item != '[' && item != '}' && item != ')' && item != ']' {
			return false
		}
		if item == '{' || item == '(' || item == '[' {
			stack.push(string(item))
			//			fmt.Println(stack.top())
		} else if item == '}' || item == ')' || item == ']' {
			if stack.Len() == 0 {
				return false
			}
			top := stack.top()
			switch item {
			case '}':
				if top != "{" {
					return false
				}
			case ']':
				if top != "[" {
					return false
				}

			case ')':
				if top != "(" {
					return false
				}
			}
			stack.pop()
		} else {
			return false
		}
	}
	if stack.Len() != 0 {
		return false
	} else {
		return true
	}

}

type Stack []string

func (s Stack) top() string {
	length := s.Len()
	fmt.Println("stack top:", s)

	return s[length-1]

}
func (s Stack) Len() int {
	fmt.Println("stack length:", s)

	return len(s)

}
func (s *Stack) pop() string {

	length := (*s).Len()
	tmp_v := (*s)[length-1]
	*s = (*s)[:length-1]
	fmt.Println("stack pop:", s)

	return tmp_v

}
func (s *Stack) push(v string) {
	*s = append(*s, v)
	fmt.Println("stack push:", *s)
}

func main() {
	s := "()"
	fmt.Println(isValid(s))

}
