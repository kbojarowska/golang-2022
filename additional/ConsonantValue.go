// https://www.codewars.com/kata/59c633e7dcc4053512000073/train/go

package main

import "fmt"

func solve(string_ string) (maximum int) {
	var sum int
	for _, r := range string_ {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			sum = 0
			continue
		}
		sum += int(r - 96)
		if sum > maximum {
			maximum = sum
		}
	}

	return maximum
}

func main() {
	fmt.Println(solve("zodiac"))
}
