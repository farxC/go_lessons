package main

import (
	"fmt"
	"time"
)

func say(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func RunningGoroutines() {
	say("direct")

	go say("goroutine")

	// Passing an anonymous function to the goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
