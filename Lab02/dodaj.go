//go run program.go -liczba1 5 -liczba2 8

package main

import (
	"flag"
	"fmt"
)

const ERR int = -100000

func main() {

	liczba1 := flag.Int("liczba1", ERR, "pierwsza liczba do sumy")
	liczba2 := flag.Int("liczba2", ERR, "druga liczba do sumy")
	flag.Parse()

	if *liczba1 == ERR {
		fmt.Print("Podaj pierwszą liczbę: ")
		fmt.Scanf("%d\n", *liczba1)
	}
	if *liczba2 == ERR {
		fmt.Print("Podaj drugą liczbę: ")
		fmt.Scanf("%d\n", *liczba2)
	}
	fmt.Println("Suma liczb", *liczba1, "+", *liczba2, "wynosi", *liczba1+*liczba2)
}
