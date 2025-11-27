package main

import "fmt"

const testingLiteral = `Hi


				That's meee!!
												Testing
																				Raw
							Literal
			Strings
																															
`

func UsingRawLiteralString() {
	fmt.Println(testingLiteral)

	const hello = "hello, 世界!"

	for i, r := range hello {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s := "abcccc"
	b := []byte(s)
	fmt.Println(b)
}
