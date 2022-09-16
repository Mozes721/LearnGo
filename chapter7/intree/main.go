package main

import (
	"fmt"
)

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(3))
}
