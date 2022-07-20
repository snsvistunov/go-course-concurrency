package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	for i, v := range n {

		wg.Add(1)
		go func(i int, v []int) {
			defer wg.Done()
			fmt.Printf("slice %v: %v\n", i, sum(v))
		}(i, v)
	}
	wg.Wait()
}

func sum(sl []int) int {
	sum := 0
	for _, v := range sl {
		sum += v
	}
	return sum
}