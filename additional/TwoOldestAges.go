// https://www.codewars.com/kata/511f11d355fe575d2c000001/train/go

package main

import (
	"fmt"
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	var result [2]int
	copy(result[:], ages[len(ages)-2:])
	return result
}

func main() {
	fmt.Println(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}))
}
