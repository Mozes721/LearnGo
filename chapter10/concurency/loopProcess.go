package concurency

import "fmt"

func Loop_process() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ch := make(chan int, len(a))

	for _, v := range a {
		go func(val int) {
			ch <- v * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
