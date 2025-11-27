package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
} // It is a factory?

func UsingStructs() {

	fmt.Println(person{"bob", 20}) // Creates a new struct

	// Naming fields when initializing the struct
	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s

	fmt.Println(s.age) // Value itself before

	//Idiomatic construction of a struct
	a := newPerson("Jon")
	fmt.Println(a)
	a.age = 55

	sp.age = 51
	fmt.Println(sp.age) // Pointer reference
	fmt.Println(s.age)  // Value itself after (changed)
}
