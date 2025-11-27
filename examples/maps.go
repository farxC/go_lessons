package main

import (
	"fmt"
	"maps"
)

func MapsUsage() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]

	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println(len(m))

	delete(m, "k2") // Removes key/value pairs when called

	clear(m) // Delete all key/value pairs from a map

	// Check if an key is present in the map
	_, prs := m["k2"]
	val, prs2 := m["k4"]
	fmt.Println("Present?", prs)
	fmt.Println("Present?", prs2, val)

	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}

	// Utility: Compare both hashes.
	if maps.Equal(n, n2) {
		fmt.Println("Is equal")
	} else {
		fmt.Println("Not equal")
	}

}
