package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func UsingSha256() {
	var val *string
	var q [32]byte
	var k [48]byte
	var t [64]byte
	fmt.Println(os.Args)
	for idx, arg := range os.Args[1:] {
		if idx == 0 && arg != "" {
			val = &arg
		} else {
			panic("NOT ENOUGH ARGUMENTS MAN!")
		}

		if idx == 1 {
			switch arg {
			case "384":
				k = sha512.Sum384([]byte(*val))
			case "512":
				t = sha512.Sum512([]byte(*val))
			}
		} else {
			q = sha256.Sum256([]byte(*val))
		}

	}
	fmt.Printf("\n%x", q)
	fmt.Printf("\n%x", k)
	fmt.Printf("\n%x", t)
	// c1 := sha256.Sum256([]byte("xx"))
	// c2 := sha256.Sum256([]byte("xx"))

	// fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}
