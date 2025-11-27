package main

import (
	"fmt"
	"time"
)

/*
Example of Unbuffered Channels
Key elements to understand:

	For unbuffered channels its important both receiver and sender ready.
	So, if you send a message the receiver needs to receive. The thread will be blocked until both ready.


	Gobyexample about channels
	By default sends and receives block until both the sender and receiver are ready.
	This property allowed us to wait at the end of our program for the "ping" message without having to use any
	other synchronization
*/
func UnbufferedChannels() {

	messages := make(chan string)

	// Why the need of a goroutine?
	// Because this allows both receiver and sender ready to accept messages
	go func() {
		messages <- "ping" // Send a _value_ to the messages channel
	}()

	msg := <-messages // Receives a _value_

	fmt.Println(msg)
}

// This turns out in a deadlock.
// The sender is waiting the receiver that will never come
// and the receiver can't receive because NEVER executes
// A Sequential Execution Problem
func UnbufferedChannelWithoutGoroutines() {
	messages := make(chan string)

	messages <- "ping"

	msg := <-messages

	fmt.Println(msg)

}

/*
By default, channels are unbuffered, meaning that they will only accepts sends if there is a corresponing
receiver
In constrast of unbuffered, buffered channels can send values into the channel without
a corresponding concurrent receiver.
*/
func BufferedChannel() {

	messages := make(chan string, 2)

	messages <- "I'm buffering a channel"
	messages <- "Hi channel"

	fmt.Println(<-messages) // Incremental
	fmt.Println(<-messages)
}

// FIFO QUEUE
func BufferedChannelOverflowing() {

	messages := make(chan string, 2)

	messages <- "I'm buffering a channel"
	messages <- "Hi channel"

	messages <- "I'm overflowing the channel"
	fmt.Println(<-messages) // Here we are withdrawing the value from the channel, so the space becomes available
	fmt.Println(<-messages)

}

func worker(channel chan bool) {
	fmt.Println("Some work to do..")
	time.Sleep(time.Millisecond)
	fmt.Println("Finishing..")

	channel <- true
}

func ChannelSynchronization() {

	done := make(chan bool, 1)
	another_done := make(chan bool, 1)

	go worker(done)
	go worker(another_done)

	<-done
	<-another_done
}

/* If we remove the receiver (<- done) this program could exit before the worker finished its work */
func ChannelWithoutSynchronization() {

	done := make(chan bool, 1)

	go worker(done)

}

// Channel Directions
func ping(pings chan<- string, msg string) {
	// Here the channel is only to send messages
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// Here the channel is only to receive messages
	msg := <-pings
	pongs <- msg
}

/*
This pattern isn't strictly enforced by the compiler but we can specify if a channel is meant
only to send or receive values. Of course this increases the type-safety of the program and prevent unexpected
behaviors..
*/
func TargetedChannels() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Hello there")
	pong(pings, pongs)

	fmt.Println(<-pongs)

}

func SelectChannels() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "from One"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "from two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}

	}
}
