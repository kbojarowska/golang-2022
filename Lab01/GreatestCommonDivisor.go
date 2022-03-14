package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	fmt.Print(GCD(3, 8))
}
