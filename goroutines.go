package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// goroutine !

	// f("direct")

	// go f("goroutine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	// channel !
	// messages := make(chan string)

	// go func() { messages <- "ping" }()

	// msg := <-messages
	// fmt.Println(msg)

	// Channel Buffering
	// messages := make(chan string, 2)

	// messages <- "buffered"
	// messages <- "channel"

	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// Channel Synchronization
	// done := make(chan bool, 1)
	// go worker(done)

	// fmt.Println("------Done!")
	// <-done
	// fmt.Println("------End!")

	// Channel Directions
	// https://gobyexample.com/channel-directions

	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)
	// ping(pings, "passed message")
	// pong(pings, pongs)
	// fmt.Println(<-pongs)

	// Select!!!
	// https://gobyexample.com/select
}
