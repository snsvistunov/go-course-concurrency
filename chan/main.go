package main

import "fmt"

func main() {
	// Default example how to use chanel with "make" function
	ch := make(chan string)
	// var ch chan string
	go func(ch chan string) { ch <- "ping" }(ch)
	msg := <-ch
	fmt.Println(msg)

	// Goroutine communication
	// ch := make(chan string)
	// go func(ch chan string) { ch <- "ping" }(ch)
	// go func(ch chan string) { fmt.Println(<-ch) }(ch)
	// time.Sleep(time.Second)

	// Read/Write nil chanel
	// var ch chan string
	// go func(ch chan string) { ch <- "ping" }(ch)
	// msg := <-ch
	// fmt.Println(msg)
}
