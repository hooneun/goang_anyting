package main

import "fmt"

// func hello(fin chan bool) {
// 	fmt.Println("Hello Goroutine")
// 	fin <- true
// }

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	// finished := make(chan bool)
	// go hello(finished)
	// <-finished // block!!!!
	// // time.Sleep(100 * time.Millisecond)
	// fmt.Println("Main")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)

	fmt.Println(s[:len(s)-1])
	fmt.Println(s[len(s)-1:])

	go sum(s[:len(s)-1], c) // 1 ~ 9
	go sum(s[len(s)-1:], c) // 10
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
