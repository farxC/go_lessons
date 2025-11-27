package main

import "fmt"

func RangeOver() {

	kvs := map[string][]string{"a": {"apple", "avocado"}, "b": {"banana", "blackberry"}}

	// Range over keys/value pairs
	for k, v := range kvs {
		fmt.Println("Letter:", k)
		fmt.Println("Values:", v)
	}

	// Range over only in keys
	for k := range kvs {
		fmt.Println(k)
	}

}
