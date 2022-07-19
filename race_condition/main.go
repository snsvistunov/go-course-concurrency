package main

import (
	"fmt"
	"sync"
)

func main() {
	const num = 1500
	var wg sync.WaitGroup
	counter := 0

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
