package main

import (
	"fmt"
	"time"
)

func f(param string) {
	for i := 0; i < 5; i++ {
		fmt.Println(param, ":", i)
		time.Sleep(time.Millisecond)
	}
}

func main() {

	f("direct")

	go f("Hello")
	go f("World!")

	for i := 0; i < 2; i++ {
		go func(msg string) {
			fmt.Println(msg)
		}(fmt.Sprintf("%s:%d", "going", i))
	}

	// time.Sleep(time.Second)
	fmt.Println("done")
}
