package main

import (
	"fmt"
)

type 


func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func main() {
	fmt.Println(div(20, 4))
}
