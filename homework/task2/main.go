package main

import "fmt"

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”

func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	// Ваша реалізація
	var res int
	ch := make(chan int)
	res = 0
	for _, v := range n {
		go func(ch chan int, v []int) {
			ch <- sum(v)
		}(ch, v)
		
	}
	for range n {
		buf := <-ch
		res += int(buf)
	}
	fmt.Printf("result: %v\n", res)
}

func sum(sl []int) int {
	var sum int
	sum = 0
	for _, v := range sl {
		sum += v
	}
	return sum
}
