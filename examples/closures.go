package main

import "fmt"

func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func PowerGenerator(base int) func() int {
	curr := 0
	return func() int {
		if curr == 0 {
			curr++
			return curr
		}

		curr = curr * base

		return curr
	}
}

func UsingClosures() {
	nextPowerOfTwo := PowerGenerator(2)

	fmt.Println(nextPowerOfTwo()) // 0
	fmt.Println(nextPowerOfTwo()) // 1
	fmt.Println(nextPowerOfTwo()) // 2
	fmt.Println(nextPowerOfTwo()) // 3
	fmt.Println(nextPowerOfTwo()) // 4
	fmt.Println(nextPowerOfTwo()) // 5
	fmt.Println(nextPowerOfTwo()) // 6
	fmt.Println(nextPowerOfTwo()) // 7
	fmt.Println(nextPowerOfTwo()) // 8

}
