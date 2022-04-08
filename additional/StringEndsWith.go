package main

// https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d/train/go

import (
	"fmt"
	"strings"
)

func stringEndsWith(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func main() {
	fmt.Println(stringEndsWith("Hello", "lo"))
	fmt.Println(stringEndsWith("Hello", "o"))
	fmt.Println(stringEndsWith("Hello", "a"))
	fmt.Println(stringEndsWith("Hello", "Hello"))
}
