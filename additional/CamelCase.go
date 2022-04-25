// https://www.codewars.com/kata/587731fda577b3d1b0001196/go

package main

import (
	"fmt"
	"strings"
)

func CamelCase(s string) string {
	return strings.ReplaceAll(strings.Title(s), " ", "")
}

func main() {
	fmt.Println(CamelCase("hello world"))
	fmt.Println(CamelCase("hello world or universe"))
	fmt.Println(CamelCase(""))
	fmt.Println(CamelCase(" "))
	fmt.Println(CamelCase("d"))
	fmt.Println(CamelCase("     b   a       g"))
}
