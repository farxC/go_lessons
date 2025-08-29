package main

import (
	"fmt"
	"strconv"
)

type my_custom_interface interface{}

func main() {
	var rn rune
	// It is an alias for i32 conventional used for distinguish character values from integer values
	fmt.Println(rn)

	// := It's a shortcut assign operator that infer what's the type of the variable based in the value present in the right.
	n := 0
	fmt.Println(n)

	const arabic_encoded = "ࢹࢸࢶࢨ"
	const chinese_hello = "你好世界"
	for i := 0; i < len(arabic_encoded); i++ {
		fmt.Printf("%#U", arabic_encoded[i])
	}

	// for i, v := range chinese_hello {
	// 	fmt.Printf("%#U starts at byte position %d", v, i)
	// }

	fmt.Println(chinese_hello)

	sv := "\nOne hour ago"   // String variable using ""
	sl := `\n This happens:` // Raw string
	fmt.Println(sl + sv)     // String concatenation

	str := "Hello world!"
	// byte is an alias for uint8 and is equivalent to uint8 in all ways.
	// It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.

	fmt.Println(len(str))

	newWlrld := string(str[0:5]) + " new world"
	// Simple concatenation

	fmt.Println(newWlrld)

	num := 10
	snum := "10"
	fmt.Println(float64(num) + 0.2)
	// Converting a number into float

	fmt.Println(strconv.Atoi(snum))
	fmt.Println(strconv.Itoa(num))

	// Declaring a interface that can receives a value of any type
	var i interface{} = "hello"

	my_str := "Golang is fun!"
	my_byte_slice := []byte(my_str)

	my_converted_slice := string(my_byte_slice)

	fmt.Println(my_byte_slice)
	fmt.Println("That is my converted slice: " + my_converted_slice)

	// That is a operator for type assertion,
	// as a result the operator returns the result of the boolean expression and the value itself
	str_value, ok := i.(string)
	if ok {
		fmt.Println("'i' holds a string:", str_value)
	} else {
		fmt.Println("'i' does not hold a string")
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	for i, v := range nums {
		fmt.Println("Index ", i, "Value:", v)
	}

	newNums := nums[3:11]
	newNums = append(newNums, 10, 20, 30)

	sliceMake := make([]int, 10)
	fmt.Println(newNums)

	fmt.Println("Using slice make:", len(sliceMake))

	indexToRemove := 9

	sliceMake = append(sliceMake[:indexToRemove], sliceMake[indexToRemove+1:]...)

	fmt.Println(len(sliceMake))
}
