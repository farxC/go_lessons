package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work done")
	case <-ctx.Done():
		fmt.Println("Context canceled:", ctx.Err())
	}
}

func UsingContext() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	go doWork(ctx) // Start the goroutine

	time.Sleep(3 * time.Second)
}
