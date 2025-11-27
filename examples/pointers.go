package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func ManipulatingUsingPointer() {
	i := 1

	fmt.Println("Initial i", i)

	zeroval(i)                 // Passing as a value, so zeroval will get a copy of ival(in the case the 'i' variable)
	fmt.Println("zeroval:", i) // Nothing has changed here.

	zeroptr(&i)                // Passing as a reference. zeroptr will access the variable address
	fmt.Println("zeroptr:", i) // Here the value becomes 0, it's was changed.
	// In this case we are applying a side effect in the variable I

	fmt.Println("pointer:", &i) // Variable address
}
