package main

import "fmt"

func Sum(nums ...int) int {

	fmt.Print(nums, " ")

	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
