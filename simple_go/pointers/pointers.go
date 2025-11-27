package main

import "fmt"

func f() *int {
	v := 1

	return &v
}

func incr(p *int) int {
	*p++ // Doesn't increment the value of p (parameter) so that function returns an side-effect to the variable passed as
	//a parameter out of the box
	return *p
}

func main() {
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 4

	fmt.Println(x)

	k := 1
	var z, y *int
	z, y = &k, &k

	fmt.Println(z == y, &z == nil, *z == *y)
	fmt.Println(z, y, &k)

	p = f()
	fmt.Println(p == f(), p)
	fmt.Println()

	u := 1
	fmt.Println(incr(&u)) // "U" now is 2

	fmt.Println(u == 2)

	a := new(int16) // P, type *int16, that points to a int without name
	fmt.Println(*a)

	*a = 5
	fmt.Println(*a)

}
