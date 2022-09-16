package main

import (
	"fmt"
)

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	f1 := myAdder.AddTo
	fmt.Println(f1(10))

}