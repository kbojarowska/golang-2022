package main

// https://www.codewars.com/kata/52b757663a95b11b3d00062d/train/go

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	result := ""
	i := 0
	for _, char := range str {
		if string(char) == " " {
			i = 0
			result += " "
		} else {
			if i%2 == 0 {
				result += strings.ToUpper(string(char))
			} else {
				result += strings.ToLower(string(char))
			}
			i++
		}
	}
	return result
}

func main() {
	fmt.Println(toWeirdCase("Hello World!"))
}
