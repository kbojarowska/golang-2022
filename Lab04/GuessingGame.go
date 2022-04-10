package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
end:
	for {
		var n string
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 100
		random := rand.Intn(max-min) + min

		fmt.Println("Welcome to the guessing game!")
		fmt.Println("Now you will be guessing the number I choosed")
		fmt.Println("If you want to quit game - type \"Exit\"")
		fmt.Println("Type number: ")
		fmt.Scan(&n)

		attempts := 0
		for {
			attempts++
			number, error := strconv.Atoi(n)
			_ = error
			if n == "Exit" {
				fmt.Println("Goodbye :(")
				break end
			} else if number > random {
				fmt.Println("Your guess is bigger than the number. Try again")
				fmt.Scan(&n)
			} else if number < random {
				fmt.Println("Your guess is smaller than the number. Try again")
				fmt.Scan(&n)
			} else {
				fmt.Println("Congratulations you win after", attempts, "attempts")
				fmt.Println("Do you want to play again? [T/N]")
				var again string
				fmt.Scan(&again)
				if again != "T" && again != "N" {
					fmt.Println("Bad input")
					fmt.Scan(&again)
				} else if again == "N" {
					break end
				} else if again == "T" {
					break
				}
			}
		}
	}
}
