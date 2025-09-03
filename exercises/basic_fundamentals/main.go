package main

import (
	"exercises/basic_fundamentals/exercises"
	"fmt"
)

func main() {

	// 1. Sum of two numbers
	sum := exercises.Sum(2)(4)
	fmt.Println(sum)

	// 2. Division of two numbers
	division := exercises.Divide(12)(10)
	fmt.Println(division)

	// 3. Successor and predecessor
	pred, suc := exercises.BetweenNumber(10)
	fmt.Println("Predecessor:", pred, "/Successor:", suc)

	// 4. Is even/odd and if is positive/negative
	isEven, isPositive := exercises.IsEvenAndPositive(-7)
	fmt.Println(isEven, isPositive)

	// 5. Is prime
	isPrime := exercises.IsPrime(8)
	fmt.Println(isPrime)

	// 6. Sort algorithm (quicksort)
	sliced_nums := []int{2, 1, 5, 8, 7, 10, 11, 21, 16, 9}
	sorted := exercises.Quicksort(sliced_nums, 1, 9)
	fmt.Println(sorted)

	// 7. Sorted string using slices.Sort() method
	my_str := "Hello world!!"
	sorted_str := exercises.CharactersSort(my_str)
	fmt.Println(sorted_str)

	// 8. Binary Decision Tree Algorithm

	// 9. Print the value and the memory address from a variable
	var p *int
	i := 42
	p = &i
	a, v := exercises.FindVariable(*p)
	fmt.Println(a, v)

	// 12. Equity between too variables
	v1 := 12
	v2 := "No way"
	var p1 *string
	var p2 *string
	myStr := "Hello world!"
	mySecondStr := "Hello world!"
	p1, p2 = &myStr, &mySecondStr
	isEqual := exercises.IsEqual(v1, v2)
	isEqualTwo := exercises.IsEqual(*p1, *p2)
	fmt.Println(isEqual)
	fmt.Println(isEqualTwo)

	// 13. Calculate vogue from a slice
	slVogue := []int{10, 20, 10, 10, 20, 20, 20}
	vogue, key := exercises.CalculateVogue(slVogue)
	fmt.Println("The vogue for the number", key, "is:", vogue)

	// 14. Is palindrome
	w := "No lemon, no melon"
	isPalindrome := exercises.IsPalindrome(w)
	fmt.Println(isPalindrome)

	// 15. Area from a retangle
	areaRetangle := exercises.AreaFromRetangle(10, 10)
	fmt.Println(areaRetangle)

	// 19. Count consoant and vowels in a text
	textExample := "AEIOU"
	conVow := exercises.CountConsonantAndVowels(textExample)
	fmt.Printf("%+v\n ", conVow)

	// 20. Count words present in a text
	secondTextExample := "Lumon lumon lumon"
	countWords := exercises.CountWords(secondTextExample)
	fmt.Println("Words count:", countWords)

	// 21. Calculate factorial from a number
	num := 8
	fat := exercises.Factorial(num)
	fmt.Println(fat)

	// 22. Print hello world.
	fmt.Println("Hello world!")

	//23. Calculate the IMC from a person and generate your classification
	height, weight := 1.78, 78
	imc := exercises.Imc(float32(weight), float32(height))
	fmt.Println(imc)
	classification := exercises.GetImcClassification(imc)
	fmt.Println(classification)

	// 24. Calculate MMC from two integers
	n1, n2 := 10, 20
	mmc := exercises.Mmc(n1, n2)
	fmt.Println(mmc)

	// 25. Calculate the average from a given list
	nums_sliced := []float32{10, 10, 10, 10, 10}
	average := exercises.Average(nums_sliced)
	fmt.Println(average)

}
