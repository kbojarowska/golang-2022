// https://www.codewars.com/kata/566be96bb3174e155300001b/train/go

package main

import (
	"fmt"
	"math"
)

func MaxBall(v0 int) int {
	v := float64(v0)

	return int(math.Round((v * 100) / (9.81 * 36)))
}

func main() {
	fmt.Println(MaxBall(15))
	fmt.Println(MaxBall(25))
}
