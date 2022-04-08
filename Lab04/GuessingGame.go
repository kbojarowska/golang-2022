package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	random := rand.Intn(max-min) + min

	fmt.Println("Welcome to the guessing game!")
	fmt.Println("Now you will be guessing the number I choosed")
	fmt.Println("Type number: ")
	fmt.Scan(&n)

	attempts := 0
	for {
		attempts++

		if n > random {
			fmt.Println("Your guess is bigger than the number. Try again")
			fmt.Scan(&n)
		} else if n < random {
			fmt.Println("Your guess is smaller than the number. Try again")
			fmt.Scan(&n)
		} else {
			fmt.Println("Congratulations you win after", attempts, "attempts")
			break
		}
	}
}
