package main

import (
	"chapter10/backpressure"
	"chapter10/concurency"

	"fmt"
)

func main() {
	//concurency.Loop_process()
	for i := range concurency.CountTo(10) {
		fmt.Println(i)
	}
	backpressure.Pressure()
}

//ch1 := make(chan int)
//ch2 := make(chan int)
//go func() {
//	v := 1
//	ch1 <- v
//	v2 := <-ch2
//	fmt.Println(v, v2)
//}()
//v := 2
//var v2 int
//select {
//case ch2 <- v:
//case v2 = <-ch1:
//}
//fmt.Println(v, v2)
