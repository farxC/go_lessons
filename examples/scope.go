package main

import (
	"log"
	"os"
)

var cwd string

func init() {
	cwd, err := os.Getwd() // Declaring cwd function scoped (is n't the same as the external cwd )

	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
