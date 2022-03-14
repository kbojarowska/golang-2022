package main

import "fmt"

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

func fibonacciPrint(n int) {
	for i := 0; i <= n-1; i += 1 {
		fmt.Printf("%d ", fibonacci(i))
	}
}

func main() {
	var n int
	fmt.Print("Type number: ")
	fmt.Scan(&n)
	fibonacciPrint(n)
}
