// https://www.codewars.com/kata/52774a314c2333f0a7000688/train/go

package main

import "fmt"

func ValidParentheses(parens string) bool {
	count := 0
	for _, c := range parens {
		if c == '(' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

func main() {
	fmt.Println(ValidParentheses("((()))()"))
	fmt.Println(ValidParentheses("((())))()"))
	fmt.Println(ValidParentheses("()"))
	fmt.Println(ValidParentheses(")"))
	fmt.Println(ValidParentheses(")()("))
}
