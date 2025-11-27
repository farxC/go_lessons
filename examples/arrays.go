package main

import "fmt"

func UsingArrays() {

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		BRL
	)

	symbol := [...]string{USD: "$", EUR: "E", GBP: "&", BRL: "R$"}

	for i, v := range symbol {
		fmt.Println(i, v)
	}

	r := [...]int{28: 1, 99: -1, 1: -4}
	fmt.Println(r)

	r[29] = 12120
	fmt.Println(r)
}
