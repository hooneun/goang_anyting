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

	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c1 <- "one"
	// }()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	c2 <- "two"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println("received", msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println("received", msg2)
	// 	}
	// }

	// Timeouts
	// https://gobyexample.com/timeouts

	// c1 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c1 <- "result 1"
	// }()

	// select {
	// case res := <-c1:
	// 	fmt.Println(res)
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("timeout 1")
	// }

	// c2 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c2 <- "result 2"
	// }()
	// select {
	// case res := <-c2:
	// 	fmt.Println(res)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("timeout 2")
	// }

	// Non-blocking channel Operations
	// https://gobyexample.com/non-blocking-channel-operations

	// messages := make(chan string)
	// signals := make(chan bool)

	// select {
	// case msg := <-messages:
	// 	fmt.Println("received message", msg)
	// default:
	// 	fmt.Println("no message received")
	// }

	// msg := "hi"
	// select {
	// case messages <- msg:
	// 	fmt.Println("sent message", msg)
	// default:
	// 	fmt.Println("no message sent")
	// }

	// select {
	// case msg := <-messages:
	// 	fmt.Println("received message", msg)
	// case sig := <-signals:
	// 	fmt.Println("received signal", sig)
	// default:
	// 	fmt.Println("no activity")
	// }

	// Closing Channels
	// https://gobyexample.com/closing-channels

	// jobs := make(chan int, 5)
	// done := make(chan bool)

	// go func() {
	// 	fmt.Println("func()")
	// 	for {
	// 		j, more := <-jobs
	// 		if more {
	// 			fmt.Println("recevied job", j)
	// 		} else {
	// 			fmt.Println("received all jobs")
	// 			done <- true
	// 			return
	// 		}
	// 	}
	// }()

	// for j := 1; j <= 3; j++ {
	// 	jobs <- j
	// 	fmt.Println("sent job", j)
	// }
	// close(jobs)
	// fmt.Println("sent all jobs")

	// <-done

	// Range over Channels
	// https://gobyexample.com/range-over-channels

	// queue := make(chan string, 2)
	// queue <- "one"
	// queue <- "two"
	// close(queue)

	// for elem := range queue {
	// 	fmt.Println(elem)
	// }

	// Timers
	// https://gobyexample.com/timers

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
