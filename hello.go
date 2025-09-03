package main

import (
	"fmt"
)

type person struct {
	fname           string
	lname           string
	age             int8
	cliche_property []string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Hi, there!"`)
}

func (sa secretAgent) speak() {
	if sa.licenseToKill {
		fmt.Println("Kill!")
	} else {
		fmt.Println("Please don't!")
	}

}

type human interface {
	speak()
} // Struct Person and Secret Agent implicity implements that interface, therefore, polymorphism.

func saySomething(h human) {
	h.speak()
}

func main() {
	xi := []int{2, 4, 7, 8, 9, 10}
	fmt.Println(xi)

	m := map[string]int{
		"Rafael": 21,
	}

	fmt.Println(m)

	p1 := person{
		fname: "Rafael",
		lname: "Nunes",
		age:   20,
	}

	sa1 := secretAgent{
		person{
			"Rafael",
			"Nunes",
			21,
			[]string{"What?"}, // Creation of slice literal of type []string containing "What?"
		},
		true,
	}

	fmt.Println(p1)
	p1.speak()
	fmt.Println(sa1)
	sa1.speak()

	my_car := car{
		color: "blue",
	}

	my_car.changeColor("red")

	saySomething(sa1)
	saySomething(p1)

}
