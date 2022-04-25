// https://www.codewars.com/kata/55e7280b40e1c4a06d0000aa/train/go

package main

import "fmt"

func ChooseBestSum(t, k int, ls []int) int {
	arr := findAllSums(k, ls)
	result := -1
	for _, v := range arr {
		if result < v && v <= t {
			result = v
		}
	}
	return result
}

func findAllSums(k int, ls []int) []int {
	if k == 1 {
		return ls
	}

	result := []int{}
	for i := 0; i < len(ls)-k+1; i++ {
		arr := findAllSums(k-1, ls[i+1:])
		for _, v := range arr {
			result = append(result, ls[i]+v)
		}
	}
	return result
}

func main() {
	fmt.Println(ChooseBestSum(174, 3, []int{50, 55, 57, 58, 60}))
}
